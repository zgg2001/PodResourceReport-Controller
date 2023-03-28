# PodResourceReport-Controller
A k8s pod resource report controller

## 快速启动(手动部署controller)
```shell
cd crd
kubectl apply -f namespaceresourcereports_crd.yaml
cd ..
go build cmd/main.go
./main -kubeconfig=$HOME/.kube/config
```

## 镜像部署controller
```shell
bash ./build.sh
kubectl apply -f example/reporter-deployment.yaml 
```

## 查看资源消耗
```shell
kubectl create ns zhj
cd example/
kubectl apply -f namespaceresourcereport_test.yaml
kubectl describe NamespaceResourceReport -n zhj
```