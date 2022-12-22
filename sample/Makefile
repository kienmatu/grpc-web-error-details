.PHONY: go
go:
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/sample.proto
.PHONY: js
js:
#PATH="${PATH}:${HOME}/go/bin"
		protoc \
		--proto_path="./proto" \
		--grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:./client-js/lib \
		--js_out=import_style=commonjs,binary:./client-js/lib \
		--plugin=protoc-gen-grpc-web=../node_modules/.bin/protoc-gen-grpc-web \
		./proto/sample.proto
.PHONY: clean
clean:
		rm ./proto/*.go
