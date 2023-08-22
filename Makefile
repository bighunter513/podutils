
all:docker

APP = getpodip

build:
	go build -o ${APP} src/main.go

docker:
	eval $(minikube docker-env) && docker buildx build . -t ${APP}

role-bind: # only execute once if local-test failed
	kubectl apply -f role_bind.yaml

local-test: ## run as a pod, use kubectl delete pod ${APP} to delete
	eval $(minikube docker-env) && kubectl run ${APP} --image=${APP}:latest --image-pull-policy=Never
	sleep 5 && kubectl logs ${APP}
	kubectl delete pod ${APP}

## create job yaml and run it
dry-run:
	kubectl create job ${APP} --image=${APP}:latest --namespace=default --image-pull-policy=Never --dry-run=client -o=yaml > ${APP}.yaml

create-job:
	kubectl create -f ./${APP}.yaml

job-logs:
	kubectl logs job/${APP}

delete-job:
	kubectl delete job ${APP}

clean:
	@rm ${APP}

.PHONY: build docker dry-run create-job job-logs delete-job clean
