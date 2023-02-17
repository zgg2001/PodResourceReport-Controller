# PodResourceReport-Controller
A k8s pod resource report controller


## 快速启动
```shell
cd crd
kubectl apply -f namespaceresourcereports_crd.yaml
cd ..
go build main.go
./main -kubeconfig=$HOME/.kube/config
```

## 查看资源消耗
```shell
kubectl create ns zhj
cd example/
kubectl apply -f namespaceresourcereport_test.yaml
kubectl describe NamespaceResourceReport -n zhj
```