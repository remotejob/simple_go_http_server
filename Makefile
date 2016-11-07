all: push

# 0.0 shouldn't clobber any released builds
TAG =0.3
PREFIX = remotejob/simple_go_http_server

binary: server.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o server

container: binary
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

set: 
	 kubectl  --kubeconfig  ~/admin.conf set image deployment/simple-go-http-server simple-go-http-server=$(PREFIX):$(TAG) -n test

clean:
	docker rmi -f $(PREFIX):$(TAG) || true
