- docker-compose up
- create topic for deposits docker-compose exec kafka1 kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 3 --partitions 1 --topic deposits
- redis
- cp env.example.yaml env.yaml
- make run-api
- installation guide protoc `https://grpc.io/docs/protoc-installation/`
- install go-protoc `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
