deploy:
		go mod vendor
		gcloud config set project arpknowledge
		gcloud auth configure-docker
		docker build -t gcr.io/arpknowledge/todos:v1 .
		docker push gcr.io/arpknowledge/todos:v1
		kubectl apply -f todos-deployment.yaml
		kubectl apply -f todos-service.yaml