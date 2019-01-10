IMGNAME="packetfire/immigrant"
APP_NAME="immigrant"
PKG="github.com/PacketFire/${APP_NAME}"

build: | fmt test
	go build

build-docker: | fmt test
	docker build -t ${IMGNAME}:latest .

test:
	go test -race -cover ./...

fmt:
	go fmt ./...

clean:
	rm -f ${APP_NAME}; \
	docker rmi -f ${IMGNAME}:latest
