## generate the grpc files
```bash
protoc --go_out=./ \    
--go-grpc_out=./ \
health/health.proto
```