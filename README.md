# podutils

this util gives you a choice to use k8s api to get a pod's IP

if err occurred, it will use interface loop to get a non-lo ipv4 ip 

## prerequistes
you need a [docker](https://www.docker.com/) envirenment 
and a [minikube](https://minikube.sigs.k8s.io/docs/start/) k8s cluster started on local 

windows推荐 wsl 2 + docker desktop + minikube
如果比较擅长linux开发，minikube 推荐安装在wsl里面，从wsl里面启动
当然安装在windows里面的minikube，也可以在wsl里面使用，不过需要加上 .exe 后缀

Linux 推荐 docker + minikube

## usage
```
  make docker
  make local-test  ## if error occurred , make role-bind first

  ## you may also  create a job task if you want
  make create-job
  make job-logs  ## to see the app's output
  make delete-job
```
