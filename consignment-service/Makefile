protoc -I. --go_out=plugins=micro:F:/learnmico/consignment-service proto/consignment/consignment.proto
go build
docker build -t consignment-service .
docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns -d consignment-service

