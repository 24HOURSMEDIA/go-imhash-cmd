BUILD_BASEDIR=./build
BUILD_FILE=${BUILD_DIR}/imhash

default: info
build: default_build

info:
	@echo Use \'make build\'

default_build: OUT_FILE = ${BUILD_FILE}
default_build: BUILD_DIR=${BUILD_BASEDIR}
default_build:
	mkdir -p ${BUILD_DIR}
	go build -o "${OUT_FILE}" -ldflags="-s -w" imhash.go
	chmod +x ${OUT_FILE}
