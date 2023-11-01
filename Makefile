TAG ?= latest

build:
	docker build -t dekuyo/migrate-go:${TAG} --target production --no-cache .

push:
	docker push dekuyo/migrate-go:${TAG}