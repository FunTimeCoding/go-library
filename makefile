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
	@gobuild cmd/gobuild/main.go
	@gobuild cmd/gobump/main.go
	@gobuild cmd/goclean/main.go
	@gobuild cmd/gocommit/main.go
	@gobuild cmd/godebian/main.go
	@gobuild cmd/golint/main.go
	@gobuild cmd/goupdate/main.go

clean:
	@rm --force --recursive tmp

install: build
	@cp tmp/gobuild $$HOME/bin/gobuild
	@cp tmp/gobump $$HOME/bin/gobump
	@cp tmp/goclean $$HOME/bin/goclean
	@cp tmp/gocommit $$HOME/bin/gocommit
	@cp tmp/godebian $$HOME/bin/godebian
	@cp tmp/golint $$HOME/bin/golint
	@cp tmp/goupdate $$HOME/bin/goupdate

monitor:
	@gobuild --copy-to-bin cmd/gomonitor/main.go
	@gobuild --copy-to-bin cmd/gosensor/main.go
	@gobuild --copy-to-bin cmd/gosentry/main.go
	@gobuild --copy-to-bin cmd/gomonitord/main.go
