// TODO： 暂时不考虑不同的 namespace，目前默认都是 default
package main

import (
	"cloudedgenetwork/crd/api/cloudedgeservice/v1alpha1"
	crdclientset "cloudedgenetwork/crd/generated/clientset/versioned"
	crdinformers "cloudedgenetwork/crd/generated/informers/externalversions"
	dbhttpclient "cloudedgenetwork/store/httpclient"
	"cloudedgenetwork/store/model"
	"flag"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	podclient        typedcorev1.PodInterface
	k8sclient        *kubernetes.Clientset
	kubeConfigPathF  = flag.String("k8sconf", "/Users/nizhenyang/Desktop/论文 workspace/code/cloudedgenetwork/k8sconfig.yaml", "k8s config path")
	serviceStoreAddr = flag.String("svcstoreaddr", "http://10.211.55.2:8881", "service db store address")
)

func main() {
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfigPathF)
	if err != nil {
		panic(err)
	}
	go crdHandler(config)

	if k8sclient, err = kubernetes.NewForConfig(config); err != nil {
		panic(err)
	}
	// TODO: 这里默认全部 Pod 都位于 default 命名空间中
	podclient = k8sclient.CoreV1().Pods(corev1.NamespaceDefault)

	select {}
}

func crdHandler(config *rest.Config) {
	crdclient, err := crdclientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	sharedInformers := crdinformers.NewSharedInformerFactory(crdclient, time.Second*2)
	informer := sharedInformers.Cloudedgeservice().V1alpha1().CloudEdgeServices().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			svc := obj.(*v1alpha1.CloudEdgeService)
			log.Printf("New svc add: %s\n", ServiceInfoString(svc))
			serviceInfo := &model.ServiceInfo{
				Name:     svc.Name,
				Port:     strconv.Itoa(int(*svc.Spec.Port)),
				Selector: svc.Spec.Selector,
			}
			if err := dbhttpclient.PutServiceInfo(*serviceStoreAddr, serviceInfo); err != nil {
				log.Println("[err] add svc to db", err)
				return
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldSVC := oldObj.(*v1alpha1.CloudEdgeService)
			newSVC := newObj.(*v1alpha1.CloudEdgeService)
			if reflect.DeepEqual(*oldSVC, *newSVC) {
				return
			}
			log.Printf("svc update: \n%s\n%s\n", ServiceInfoString(oldSVC), ServiceInfoString(newSVC))
			serviceInfo := &model.ServiceInfo{
				Name:     newSVC.Name,
				Port:     strconv.Itoa(int(*newSVC.Spec.Port)),
				Selector: newSVC.Spec.Selector,
			}
			if err := dbhttpclient.PutServiceInfo(*serviceStoreAddr, serviceInfo); err != nil {
				log.Println("[err] update svc in db", err)
				return
			}
		},
		DeleteFunc: func(obj interface{}) {
			svc := obj.(*v1alpha1.CloudEdgeService)
			log.Printf("svc delete: %s\n", ServiceInfoString(svc))
			if err := dbhttpclient.DeleteServiceInfo(*serviceStoreAddr, svc.Name); err != nil {
				log.Println("[err] delete svc in db", err)
				return
			}
		},
	})
	stopCh := make(chan struct{})
	defer close(stopCh)
	informer.Run(stopCh)
}

func ServiceInfoString(ces *v1alpha1.CloudEdgeService) string {
	return fmt.Sprintf("Name: %s\tPort: %d\tLabels: %+v", ces.Name, *ces.Spec.Port, ces.Spec.Selector)
}
