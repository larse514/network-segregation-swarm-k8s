default: build

build: 
	docker build . --tag ${TAG}

dependencies: 
	@go get net/http
	@go get encoding/json

test: test-all

test-all:
	@go test -v ./...

test-min:
	@go test ./...

run:
	docker run -p 80:80 -d ${TAG} 

release:
	docker login -u ${DOCKER_USER} -p ${DOCKER_PASSWORD}
	docker push ${TAG}