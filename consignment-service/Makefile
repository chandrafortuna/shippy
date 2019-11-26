build:
	protoc -I. --go_out=plugins=grpc:. \
	  proto/consignment/consignment.proto

	GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-service-consignment .

run:
	docker run -p 5050:5050 shippy-service-consignment