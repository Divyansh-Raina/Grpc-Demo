BIN = bin
PROTO = proto
SERVER = server
CLIENT = client
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
.DEFAULT_GOAL := help
.PHONY: v1 v2
project := v1 v2

all: $(project)

v1:$@
v2:$@

$(project):
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. \
	--go-grpc_opt=module=${PACKAGE} \
	--go-grpc_out=. \
	$@/${PROTO}/*.proto
	go build -o ${BIN}/$@/${SERVER} ./$@/${SERVER}
	go build -o ${BIN}/$@/${CLIENT} ./$@/${CLIENT}

test: all
	go test ./...


clean: clean_v1 clean_v2
	rm -rf ssl/*.crt
	rm -rf ssl/*.cse
	rm -rf ssl/*.key
	rm -rf ssl/*.pem

clean_v1:
	rm -f v1/${PROTO}/*.pb.go

clean_v2:
	rm -f v2/${PROTO}/*.pb.go
	
clean_all: clean all

bump: all
	go get -u ./..