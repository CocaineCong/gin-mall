.PHONY: build		# 构建项目
build:
	@echo "build project to ./main"
	@cd ./cmd && go build -o ../main

.PHONY: test		# 构建同时运行
test:
	@echo "build && run project at ./main"
	@cd ./cmd && go build -o ../main
	@./main

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
	-t gin_mall:3.0 \
	-f ./Dockerfile \
	./
	docker run \
	-it \
	--name mall_server \
	--network host \
	-d gin_mall:3.0
	@echo "container run success at localhost:5001"

.PHONY: docker-down # 结束docker部署,同时删除容器和镜像
docker-down:
	docker stop mall_server
	docker rm mall_server
	docker rmi gin_mall:3.0
	@echo "container stop && rm success"

default: test