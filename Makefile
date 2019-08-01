appName ?= "gapp"
echo:
	@echo $(appName)
win32:
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 GO111MODULE=on GOPROXY=https://goproxy.io go build -o bin/win/32/$(appName).exe
win64:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.io go build -o bin/win/64/$(appName).exe
mac32:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=386 GO111MODULE=on GOPROXY=https://goproxy.io go build -o bin/mac/32/$(appName)
mac64:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.io go build -o bin/mac/64/$(appName)
linux32:
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 GO111MODULE=on GOPROXY=https://goproxy.io go build -o bin/linux/32/$(appName)
linux64:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPROXY=https://goproxy.io go build -o bin/linux/64/$(appName)
clear:
	@rm -rf bin/mac/*
	@rm -rf bin/linux/*
	@rm -rf bin/win/*
all: clear win32 win64 mac32 mac64 linux32 linux64
dev:
	@GO111MODULE=on GOPROXY=https://goproxy.io go run -race main.go