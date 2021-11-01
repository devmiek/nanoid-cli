#!/bin/bash

VERSION_TAG=0.1.1
OUTPUT_DIR_NAME=./dist
OUTPUT_FILE_NAME=nanoid

function build_package() {
    if [ $1 == "windows" ]
    then
        BIN_FILE_NAME="$OUTPUT_FILE_NAME.exe"
    else
        BIN_FILE_NAME="$OUTPUT_FILE_NAME"
    fi

    if [ ! -d $OUTPUT_DIR_NAME ]
    then
        mkdir $OUTPUT_DIR_NAME
    fi

    echo "Building package $1/$2..."

    GOOS=$1 GOARCH=$2 go build -ldflags "-s -w" -o "$OUTPUT_DIR_NAME/$BIN_FILE_NAME"
    tar -czf "$OUTPUT_DIR_NAME/$OUTPUT_FILE_NAME-$VERSION_TAG-$1-$2.tar.gz" \
        -C $OUTPUT_DIR_NAME $BIN_FILE_NAME
    rm -f "$OUTPUT_DIR_NAME/$BIN_FILE_NAME"
}

build_package "windows" "386"
build_package "windows" "amd64"
build_package "windows" "arm"
build_package "windows" "arm64"

build_package "linux" "386"
build_package "linux" "amd64"
build_package "linux" "arm"
build_package "linux" "arm64"

build_package "darwin" "amd64"
build_package "darwin" "arm64"
