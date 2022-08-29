# コマンド一覧
.PHONY: help
help:
	@cat Makefile
# dockerのビルド
.PHONY: dc/build
dc/build:
	docker-compose build

# コンテナ立ち上げ
.PHONY: dc/up
dc/up:
	docker-compose up

# コンテナ, image, volumesの削除
.PHONY: dc/down-remove
dc/down-remove:
	docker-compose down --rmi all --volumes --remove-orphans
