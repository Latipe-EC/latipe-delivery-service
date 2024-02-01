#setup > wire > clean > build > run


SERVICE_NAME = delivery_service_v1
WORKER_MAIN_FILE = server_app
BUILD_DIR = $(PWD)/build
GO= go

setup:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

wire:
	cd internal/ && wire

wire-linux:
	cd internal/ && ~/go/bin/wire

protoc:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative .\internal\grpc-service\protobuf\deliveryGrpc\delivery.proto

gen-cert:
	openssl genrsa -out /cert/server.key 2048
	openssl req -new -x509 -days 3650 -subj "/C=VN/ST=HoChiMinh/L=HoChiMinh/O=UTE/OU=UTE Education/CN=delivery-services.vn" -addext "subjectAltName=DNS:*.delivery-services.vn,IP:0.0.0.0" -key /cert/server.key -out /cert/server.crt
	openssl genrsa -out /cert/client.key 2048
	openssl req -new -x509 -days 3650 -subj "/C=VN/ST=HoChiMinh/L=HoChiMinh/O=UTE/OU=UTE Education/CN=delivery-services.vn" -addext "subjectAltName=DNS:*.delivery-services.vn,IP:0.0.0.0" -key /cert/client.key -out /cert/client.crt

#linux
# clean build file
cleanl:
	echo "remove bin exe"
	rm -rf $(BUILD_DIR)

# build binary
buildl:
	echo "build binary execute file"
	make wire-linux
	cd cmd/ && GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(WORKER_MAIN_FILE)_linux .

runl:
	make buildl

	echo "Run service application"
	cd $(BUILD_DIR) && ./$(WORKER_MAIN_FILE)_linux


#windows
cleanw:
	echo "remove bin exe"
	rd /s build

buildw:
	echo "build binary execute file"
	make wire
	cd cmd/ &&  $(GO) build -o ..$(BUILD_DIR)/$(WORKER_MAIN_FILE)_win.exe .

runw:
	.\$(BUILD_DIR)\$(WORKER_MAIN_FILE)_win.exe

startw:
	make buildw
	.\$(BUILD_DIR)\$(WORKER_MAIN_FILE)_win.exe