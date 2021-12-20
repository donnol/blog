PHONY:genproto

# install protoc
# 在[protobuf release](https://github.com/protocolbuffers/protobuf/releases)页面下载protoc，里面会有protoc工具和一些基础依赖包

# 因为protoc原生不支持go，所以需要安装protoc-gen-go插件
install_protoc-gen-go:
	go install github.com/golang/protobuf/protoc-gen-go@latest

# grpc
install_protoc-gen-go-grpc:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# --go-grpc_out=. \
# --go-grpc_opt=paths=source_relative
genproto:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative \
    	./*.proto
