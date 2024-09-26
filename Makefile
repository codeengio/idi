APP=idi
MODULE := github.com/codeengio/idi
VERSION := 1.0.0

.PHONY: clean bin test

all: clean zip

clean:
	rm -rf bin release

test:
	go test -p 1 -count=1 -v ./...

zip: release/$(APP)_$(VERSION)_linux_x86_64.tar.gz release/$(APP)_$(VERSION)_osx_x86_64.tar.gz

binaries: binaries/linux_x86_64/$(APP) binaries/osx_x86_64/$(APP)

release/$(APP)_$(VERSION)_linux_x86_64.tar.gz: binaries/linux_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_linux_x86_64.tar.gz -C bin/linux_x86_64 $(APP)_$(VERSION)

binaries/linux_x86_64/$(APP):
	GOOS=linux GOARCH=amd64 go build -o bin/linux_x86_64/$(APP)_$(VERSION) main.go

release/$(APP)_$(VERSION)_osx_x86_64.tar.gz: binaries/osx_x86_64/$(APP)
	mkdir -p release
	tar cfz release/$(APP)_$(VERSION)_osx_x86_64.tar.gz -C bin/osx_x86_64 $(APP)_$(VERSION)

binaries/osx_x86_64/$(APP):
	GOOS=darwin GOARCH=amd64 go build -o bin/osx_x86_64/$(APP)_$(VERSION) main.go
