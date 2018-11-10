IMGNAME="packetfire/immigrant"

build: | depend fmt test
	go build

depend:
	glide update ; glide install

build-docker: | depend fmt test
	docker build -t ${IMGNAME}:`cat "version.txt"` .

test: | depend
	go test ./...

fmt: | depend
	go fmt ./...

clean:
	rm -f immigrant; \
	docker rmi -f ${IMGNAME}:`cat "version.txt"`
