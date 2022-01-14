## 生成grpc代码

cd server/proto/v1

执行

```bigquery
 protoc --go_out=. --go-grpc_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_opt=paths=source_relative \
    hello.proto
```