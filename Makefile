APP := mkpasswd
DIR := bin

.PHONY: all clean build windows linux macos freebsd

all: windows linux macos freebsd

clean:
	@if [ -d ${DIR} ]; then rm -rf ${DIR}/*; else exit 0; fi

build:
	go build -o ${DIR}/${APP} main.go

windows:
	@# windows-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${DIR}/${APP}.exe main.go
	@cd ${DIR}/ && zip -qr ${APP}-windows_amd64.zip ${APP}.exe && rm -rf ${APP}.exe && cd ../
	@# windows-arm64:
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o ${DIR}/${APP}.exe main.go
	@cd ${DIR}/ && zip -qr ${APP}-windows_arm64.zip ${APP}.exe && rm -rf ${APP}.exe && cd ../

linux:
	@# linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -zcf ${APP}-linux_amd64.tar.gz ${APP} && rm -rf ${APP} && cd ../
	@# linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -zcf ${APP}-linux_arm64.tar.gz ${APP} && rm -rf ${APP} && cd ../

macos:
	@# macos-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -zcf ${APP}-darwin_amd64.tar.gz ${APP} && rm -rf ${APP} && cd ../
	@# macos-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -zcf ${APP}-darwin_arm64.tar.gz ${APP} && rm -rf ${APP} && cd ../

freebsd:
	@# freebsd-amd64:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -jcf ${APP}-freebsd_amd64.tar.bz2 ${APP} && rm -rf ${APP} && cd ../
	@# freebsd-arm64:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 go build -o ${DIR}/${APP} main.go
	@cd ${DIR}/ && tar -jcf ${APP}-freebsd_arm64.tar.bz2 ${APP} && rm -rf ${APP} && cd ../
