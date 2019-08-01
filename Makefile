build:
	protoc -I. --go_out=plugins=micro:. \
		proto/discount.proto
	docker build -t discount-api .
run:
	docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		discount-api
