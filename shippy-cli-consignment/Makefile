build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-cli-consignment .

run:
	docker run -p 8080:8080 \ 
    -e MICRO_REGISTRY=mdns \
    microhq/micro api \
    --handler=rpc \
    --address=:8080 \
    --namespace=shippy
	 shippy-cli-consignment

#	export PATH=$PATH:$HOME/go/bin
#	export PATH=$PATH:/usr/local/go/bin