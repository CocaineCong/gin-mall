# For unix-based system, such as macOS/Linux, ARCH can be replace with $(shell uname -m)
# For windows, can use 'systeminfo | findstr /C:"System Type"' to get ARCH
# Or easily, just use 'amd64' for windows

# amd64 arm64
ARCH = arm64

# linux darwin windows
OS = darwin

DIR := $(shell pwd)
OUTPUT = main

BINARY = skywalking-go-agent
TOOLS_PATH = $(DIR)/tools
AGENT_SOURCE_PATH = $(DIR)/skywalking-go/tools/go-agent
AGENT_PATH = $(TOOLS_PATH)/$(BINARY)-$(VERSION)-$(OS)-$(ARCH)
AGENT_CONFIG = $(DIR)/config/locales/agent.yaml

CONTAINER_NAME = gin_mall_server
IMAGE_NAME = gin_mall:3.0

GO = go
GO_BUILD = $(GO) build
GO_BUILD_FLAGS = -v
GO_BUILD_LDFLAGS = -X main.version=$(VERSION)

.PHONY: tools
tools:
	cd $(AGENT_SOURCE_PATH) && make deps
	cd $(AGENT_SOURCE_PATH) && \
	GOOS=$(OS) GOARCH=$(ARCH) $(GO_BUILD) $(GO_BUILD_FLAGS) -ldflags "$(GO_BUILD_LDFLAGS)" -o $(TOOLS_PATH)/$(BINARY)-$(VERSION)-$(OS)-$(ARCH) ./cmd

.PHONY: run			# 构建同时运行
test:
	@make build
	@./$(OUTPUT)

.PHONY: build		# 构建项目
build:
	@echo "build project to ./$(OUTPUT)"
	$(GO_BUILD) \
	-toolexec="$(AGENT_PATH) -config $(AGENT_CONFIG)" \
	-a -o ./$(OUTPUT) ./cmd

.PHONY: env-up		# 启动环境
env-up:
	docker-compose up -d
	@echo "env start success"

.PHONY: env-down	# 关闭环境
env-down:
	docker-compose down
	@echo "env stop success"

.PHONY: docker-up	# 以容器形式部署项目
docker-up:
	docker build \
	-t $(IMAGE_NAME) \
	-f ./Dockerfile \
	./
	docker run \
	-it \
	--name $(CONTAINER_NAME) \
	--network host \
	-d $(IMAGE_NAME)
	@echo "container run success at localhost:5001"

.PHONY: docker-down # 结束docker部署,同时删除容器和镜像
docker-down:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)
	@echo "container stop && rm success"

default: run