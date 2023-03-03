migrate-up:
	migrate -path ./migrations \
	-database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' \
	up

migrate-down:
	migrate -path ./migrations \
	-database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' \
	down

protoc-gen: gen-server gen-web

PROTO_FILEPATH="service.proto"
OUT_DIR="./pitch/src/generated"
gen-web:
	mkdir -p ${OUT_DIR}
	#protoc --ts_out ${OUT_DIR} --proto_path "." ${PROTO_FILEPATH}

#	protoc -I=. ${PROTO_FILEPATH} \
#  	--js_out=import_style=commonjs:${OUT_DIR} \
#  	--grpc-web_out=import_style=typescript,mode=grpcwebtext:${OUT_DIR}

	protoc service.proto \
    --js_out=import_style=commonjs,binary:${OUT_DIR} \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${OUT_DIR}

gen-service:
	mkdir -p pb
	protoc --go_out=./pb --go_opt=paths=source_relative \
		--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
		./service.proto

start-server:
	go run github.com/mokan-r/pitch/server/cmd