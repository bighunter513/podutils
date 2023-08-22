# podutils

this util gives you a choice to use k8s api to get a pod's IP

if err occurred, it will use interface loop to get a non-lo ipv4 ip 

## prerequistes
you need a [docker](https://www.docker.com/) envirenment 
and a [minikube](https://minikube.sigs.k8s.io/docs/start/) k8s cluster started on local 

## usage
```
  make docker
  make local-test  ## if error occurred , make role-bind first

  ## you may also  create a job task if you want
  make create-job
  make job-logs  ## to see the app's output
  make delete-job
```
