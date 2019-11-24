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

clean-docker:
	@type docker >/dev/null 2>&1 && \
	docker rmi -f ${IMGNAME}:latest || \
	true

clean: clean-docker
	@rm -f ${APP_NAME} || true
