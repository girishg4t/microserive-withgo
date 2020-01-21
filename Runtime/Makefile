build:
	docker save shippy-service-vessel:v2 | (eval $(minikube docker-env) && docker load)
	docker save shippy-service-consignment:v2 | (eval $(minikube docker-env) && docker load)
	docker save shippy-server:v2 | (eval $(minikube docker-env) && docker load)
	docker save shippy-ui:v2 | (eval $(minikube docker-env) && docker load)