# run:
# 	docker run -p 8080:8080 -e MICRO_REGISTRY=mdns microhq/micro api --enable_rpc
# docker run -p 8000:80 shippy-ui
build:
	docker build -t shippy-server:v2 .

run:
	docker run -d --net="host" --expose=8083 shippy-server 