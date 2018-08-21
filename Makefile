C_ID := $(shell docker ps -aq)

build:
	@docker build -t newsapp-$(name):latest ./$(name)

run:
	@docker run -d -p 5000:5000 -e PORT=5000 newsapp-$(name):latest

stop:
	@docker stop $(C_ID)

remove:
	@docker rm $(C_ID)

clean: stop remove

test:
	@curl http://localhost:5000/
	@curl http://localhost:5000/api/v1/time
	@curl http://localhost:5000/api/v1/source

loadtest:
	@for i in `seq 1 $(n)`; do curl http://localhost:5000/; done;
	@for i in `seq 1 $(n)`; do curl http://localhost:5000/api/v1/time; done;
	@for i in `seq 1 $(n)`; do curl http://localhost:5000/api/v1/source; done;

tag:
	@docker tag newsapp-$(name):latest balanus/newsapp-$(name):latest

push:
	@docker push balanus/newsapp-$(name):latest

k8s-create-namespace:
	@kubectl apply -f kube/namespace/app-namespace.yml

k8s-expose-as-common-service:
	@kubectl apply -f kube/services/app-service.yml

k8s-expose-nodejs-service:
	@kubectl apply -f kube/nodejs/nodejs-app-service.yml

k8s-expose-go-service:
	@kubectl apply -f kube/go/go-app-service.yml

k8s-expose-python-service:
	@kubectl apply -f kube/python/python-app-service.yml


k8s-deploy-nodejs:
	@kubectl apply -f kube/nodejs/nodejs-app-deployment.yml

k8s-deploy-go:
	@kubectl apply -f kube/go/go-app-deployment.yml

k8s-deploy-python:
	@kubectl apply -f kube/python/python-app-deployment.yml

k8s-delete-deployment-python:
	@kubectl delete deployment news-app-python --namespace=news-app-ns

k8s-delete-deployment-go:
	@kubectl delete deployment news-app-go --namespace=news-app-ns

k8s-delete-deployment-nodejs:
	@kubectl delete deployment news-app-nodejs --namespace=news-app-ns




k8s-get-ns:
	@kubectl get ns

k8s-get-svc:
	@kubectl get svc --namespace=news-app-ns

k8s-get-pods:
	@kubectl get pods --namespace=news-app-ns

k8s-get-deployment:
	@kubectl get deployment --namespace=news-app-ns