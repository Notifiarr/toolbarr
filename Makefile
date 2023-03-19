# The Makefile is not used for GitHub Actions or official builds. Only used for development and testing.
VERSION=$(shell git describe --abbrev=0 --tags $(shell git rev-list --tags --max-count=1) | tr -d v)
COMMIT="$(shell git rev-parse --short HEAD || echo 0)"
BRANCH="$(shell git rev-parse --abbrev-ref HEAD || echo unknown)"
ifeq ($(VERSION),)
	VERSION:=0.0.0
endif
ifeq ($(REVISION),)
	REVISION=$(shell git rev-list --count --all || echo 0)
endif

VERSION_LDFLAGS:= -X \"golift.io/version.Branch=$(BRANCH) ($(COMMIT))\" \
	-X \"golift.io/version.BuildDate=$(shell date -u +%Y-%m-%dT%H:%M:00Z)\" \
	-X \"golift.io/version.BuildUser=$(shell whoami || echo "unknown")\" \
	-X \"golift.io/version.Revision=$(REVISION)\" \
	-X \"golift.io/version.Version=$(VERSION)\"

.PHONY: all wailsjson build windows dev lint test

all: 
	echo "Use 'make build', 'make dev' or 'make windows'"

wailsjson:
	jq ".info.productVersion = \"$(VERSION)\"" wails.json > wails.json.new
	mv wails.json.new wails.json

build: wailsjson
	wails build -ldflags "$(VERSION_LDFLAGS)"

windows: wailsjson
	wails build -platform windows/amd64 -nsis -ldflags "$(VERSION_LDFLAGS)"

dev: wailsjson
	wails dev -nosyncgomod -race -ldflags "$(VERSION_LDFLAGS)"

# npm?? svelte? hm..
lint:
	codespell -S .git,node_modules,dist,catalog.go,package-lock.json
	GOOS=darwin golangci-lint run
	GOOS=linux golangci-lint run
	GOOS=windows golangci-lint run

test: lint
	go test -race ./...