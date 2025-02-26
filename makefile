.DEFAULT_GOAL := all

all: test lint build

tool:
	@go install gotest.tools/gotestsum@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/gobuild@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/golint@latest
	@GOPROXY=direct go install github.com/funtimecoding/go-library/cmd/goupdate@latest

test:
	@gotestsum --format standard-quiet -- ./...

lint:
	@golint
	@golangci-lint run

update:
    # k8s.io/apimachinery: Cilium is not ready for 0.31.0
    # sigs.k8s.io/structured-merge-diff/v4: IgnoredFields error in structuredmerge.go
	@goupdate --downgrade k8s.io/apimachinery@v0.30.8 \
		--downgrade sigs.k8s.io/structured-merge-diff/v4@v4.4.1

build:
	@go run cmd/gobuild/main.go gobuild
	@go run cmd/gobuild/main.go gobump
	@go run cmd/gobuild/main.go goclean
	@go run cmd/gobuild/main.go gocommit
	@go run cmd/gobuild/main.go godebian
	@go run cmd/gobuild/main.go golint
	@go run cmd/gobuild/main.go goupdate
	@go run cmd/gobuild/main.go godownload
	@go run cmd/gobuild/main.go gopackage
	@go run cmd/gobuild/main.go goupload

clean:
	@rm --force --recursive tmp

install:
	@go run cmd/gobuild/main.go --copy-to-bin gobuild
	@gobuild --copy-to-bin gobump
	@gobuild --copy-to-bin goclean
	@gobuild --copy-to-bin gocommit
	@gobuild --copy-to-bin godebian
	@gobuild --copy-to-bin golint
	@gobuild --copy-to-bin goupdate
	@gobuild --copy-to-bin godownload
	@gobuild --copy-to-bin gopackage
	@gobuild --copy-to-bin goupload

monitor:
	@gobuild --copy-to-bin gomonitor
	@gobuild --copy-to-bin gosensor
	@gobuild --copy-to-bin gosentry
	@gobuild --copy-to-bin gomonitord

chk:
	@gobuild --copy-to-bin gochk

deploy: chk
	tmp/deploy.sh
