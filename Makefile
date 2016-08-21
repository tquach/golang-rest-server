REPO=tquach
APP_NAME=golang-rest-server
TAG=latest

deps:
	@hash godep || { echo "Install godep first: go get github.com/tools/godep"; exit 1; }

all: $(APP_NAME)

$(APP_NAME): test
	@echo "Running tests..."
	@godep go build .

build:
	@echo "Building ${APP_NAME}/${APP_NAME}:${TAG} ..."
	docker build -t $(REPO)/$(APP_NAME):$(TAG) .

start: build
	docker run -it --rm --name $(APP_NAME) -p 9000:9000 $(REPO)/$(APP_NAME) golang-rest-server --hostname localhost:9000

test: deps
	@godep go test ./... 

clean:
	@rm -f $(APP_NAME)

deploy: build
	docker push $(REPO)/$(APP_NAME):$(TAG) 

.PHONY:
	start clean test build deploy