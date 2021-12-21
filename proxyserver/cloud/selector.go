package main

import (
	dbhttpclient "cloudedgenetwork/store/httpclient"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SelectTarget(serviceName string) (targetNode, targetAddr string, err error) {
	serviceInfo, err := dbhttpclient.GetServiceInfo(*serviceDBAddr, serviceName)
	if err != nil {
		return "", "", fmt.Errorf("query service db: %v", err)
	}
	// TODO: just randomly select one pod from list
	backend := serviceInfo.Backends[rand.Intn(len(serviceInfo.Backends))]
	return backend.NodeName, backend.ContainerIP + ":" + serviceInfo.Port, nil
}
