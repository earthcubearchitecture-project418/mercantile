BINARY := server
DOCKERVER :=`/bin/cat VERSION`
.DEFAULT_GOAL := linux

linux:
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)

docker:
		docker build  --tag="fils/mercantile:$(DOCKERVER)"  --file=./build/Dockerfile .

dockerlatest:
		docker build  --tag="fils/mercantile:latest"  --file=./build/Dockerfile .

publish:  
		docker push fils/mercantile:$(DOCKERVER)
		docker push fils/mercantile:latest

buildpub: linux docker dockerlatest publish
