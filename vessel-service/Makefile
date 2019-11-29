build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/vessel/vessel.proto
	docker build -t shippy-vessel-service .

run:
	docker run -p 5052:5052 -e MICRO_SERVER_ADDRESS=:5052 shippy-vessel-service