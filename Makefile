build:
	protoc -I. --go_out=plugins=micro:. \
		proto/discount.proto
	sudo docker build -t discount-api .
run:
	sudo docker run -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		discount-api
