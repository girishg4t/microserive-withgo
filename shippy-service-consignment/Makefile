build:
	protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/consignment/consignment.proto
	go build
	docker build -t shippy-service-consignment:v2 .

run:
	docker run -d --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		shippy-service-consignment

deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml