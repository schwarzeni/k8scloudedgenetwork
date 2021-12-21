rm -rf ./crd/generated
go mod vendor
bash vendor/k8s.io/code-generator/generate-groups.sh all \
cloudedgenetwork/crd/generated  cloudedgenetwork/crd/api \
cloudedgeservice:v1alpha1 \
--go-header-file hack/boilerplate.go.txt \
--output-base ../
rm -rf vendor
