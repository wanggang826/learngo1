protoc --go_out=./  person/person.proto
protoc --go_out=./  home/home.proto

protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  person/person.proto
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  home/home.proto