.PHONY: build clean deploy

build:
	export GO111MODULE=on GOPROXY=https://proxy.golang.org
	env GOOS=linux go build -ldflags="-s -w" -o bin/service cmd/service/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --stage ${STAGE:-dev} --verbose
