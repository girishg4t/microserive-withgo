
build:
	protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/vessel/vessel.proto  
	docker build -t shippy-service-vessel:v2 .

run:
	docker run -d --net="host" \
		-p 50053 \
		-e MICRO_SERVER_ADDRESS=:50053 \
		-e MICRO_REGISTRY=mdns \
		shippy-service-vessel