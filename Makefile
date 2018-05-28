BUILDIMG="golang:1.10.1-alpine"
PKGNAME="github.com/ncatelli/immigrant"
IMGNAME="ncatelli/immigrant"

build: | fmt test
	go build

build-docker: | fmt test
	docker run --rm -v ${PWD}:/go/src/${PKGNAME} -e GOPATH=/go -w /go/src/${PKGNAME}/ ${BUILDIMG} go build; \
  docker build -t ${IMGNAME}:`cat "version.txt"` .

test:
	go test ./...

fmt: 
	go fmt ./...

clean:
	rm -f immigrant; \
	docker rmi -f ${IMGNAME}:`cat "version.txt"`
