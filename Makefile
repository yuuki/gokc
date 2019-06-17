export GO111MODULE=on

PKG = github.com/yuuki/gokc
COMMIT = $$(git describe --tags --always)
DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)
RELEASE_BUILD_LDFLAGS = -s -w $(BUILD_LDFLAGS)

all: build

yacc: deps
	goyacc -o parser/parser.go -v parser/parser.output parser/parser.go.y

build: yacc
	go build -ldflags="$(BUILD_LDFLAGS)" $(PKG)

test: yacc
	go test -v $$(go list ./... | grep -v vendor)

testfiles: build
	find ./testdata -type f | xargs -I{} ./$(BIN) -f {}
	find ./keepalived/doc/samples/keepalived.conf.* -type f | xargs -I{} ./$(BIN) -f {}

.PHONY: devel-deps
devel-deps:
	GO111MODULE=off go get -v \
	golang.org/x/tools/cmd/cover \
	github.com/mattn/goveralls \
	github.com/motemen/gobump/cmd/gobump \
	github.com/Songmu/ghch/cmd/ghch \
	github.com/Songmu/goxz/cmd/goxz \
	github.com/tcnksm/ghr \
	github.com/Songmu/gocredits/cmd/gocredits

deps:
	go get golang.org/x/tools/cmd/goyacc
	go get -d -v ./...

.PHONY: crossbuild
crossbuild: devel-deps
	$(eval ver = $(shell gobump show -r))
	goxz -pv=v$(ver) -os=linux,darwin -arch=386,amd64 -build-ldflags="$(RELEASE_BUILD_LDFLAGS)" \
	  -d=./dist/v$(ver)

.PHONY: release
release: devel-deps
	_tools/release
	_tools/upload_artifacts

clean:
	go clean

.PHONY: yacc build clean deps test buildlinux patch minor gobump
