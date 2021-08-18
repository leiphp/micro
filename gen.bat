cd protos
protoc  --micro_out=. --go_out=. blog.proto
protoc-go-inject-tag -input=blog.pb.go
cd ..

