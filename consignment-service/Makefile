build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/consignment/consignment.proto

	#   protoc --micro_out=. --go_out=. proto/consignment/consignment.proto


	# GOOS=linux GOARCH=amd64 go build
	docker build -t shippy-service-consignment .

run:
	docker run -p 5050:5050	\
	-e MICRO_SERVER_ADDRESS=:5050	\
	shippy-service-consignment