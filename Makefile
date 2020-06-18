BINARY := server
DOCKERVER :=`/bin/cat VERSION`
.DEFAULT_GOAL := linux

linux:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)

docker:
		docker build  --tag="fils/ecgql:$(DOCKERVER)"  --file=./build/Dockerfile .

dockerlatest:
		docker build  --tag="fils/ecgql:latest"  --file=./build/Dockerfile .

publish:  
		docker push fils/ecgql:$(DOCKERVER)
		docker push fils/ecgql:latest

