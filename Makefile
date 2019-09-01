TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

build:
	protoc -I. --go_out=plugins=micro:. \
		proto/discount.proto
pack:	
	sudo docker build -t supermarket.discount.api:$(TAG) .
run:
	sudo docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		supermarket.discount.api:$(TAG)

tag: 
	sudo docker tag supermarket.discount.api:$(TAG) sisimogangg/supermarket.discount.api:$(TAG)

push:
	sudo docker push sisimogangg/supermarket.discount.api
