PKG = github.com/gilliek/go-opml/opml

all: install

install: check test
	go install ${PKG}

build:
	go build ${PKG}

check:
	go vet ${PKG}
	golint ${GOPATH}/src/${PKG}

test:
	go test ${PKG}

cover:
	@go test -coverprofile=c.out ${PKG}
	@go tool cover -html=c.out -o coverage.html
	@go tool cover -func=c.out
	@rm c.out
	@echo ""
	@echo "See coverage.html for more details"

