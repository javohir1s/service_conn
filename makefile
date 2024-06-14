CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

gen-proto-module:
	sudo rm -rf ${CURRENT_DIR}/genproto
	./scripts/gen_proto.sh ${CURRENT_DIR}

migrateup:
	migrate -path ./migrations -database 'postgres://mirodil:1212@localhost:5432/catalog_service?sslmode=disable' up

migratedown:
	migrate -path ./migrations -database 'postgres://mirodil:1212@localhost:5432/catalog_service?sslmode=disable' down

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge
