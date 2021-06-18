proto:
	rm -rf protobuftmp
	mkdir protobuftmp
	protoc --go_out=plugins=grpc:protobuftmp  protobuf/service/*.proto
	protoc --go_out=plugins=grpc:protobuftmp  protobuf/request/*.proto
	protoc --go_out=plugins=grpc:protobuftmp  protobuf/response/*.proto
	rm -rf micro/protobuf/*
	mv protobuftmp/JupiterBlog/micro/* micro/
	rm -rf protobuftmp
article:
	rm -rf 	exe/article
	go build -o exe/article main/article.go
	exe/article   --config=config/article.toml
blog:
	rm -rf 	exe/blog
	go build -o exe/blog main/blog.go
	exe/blog   --config=config/blog.toml
