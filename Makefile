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

