生成 proto 文件
```shell
protoc --go_out=. --go_opt=paths=source_relative user.proto  
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative user.proto
```