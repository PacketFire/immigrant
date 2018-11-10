IMGNAME="packetfire/immigrant"

depend:
	glide update ; glide install

build: | fmt test
	go build

build-docker: | fmt test
	docker build -t ${IMGNAME}:`cat "version.txt"` .

test:
	go test ./...

fmt: 
	go fmt ./...

clean:
	rm -f immigrant; \
	docker rmi -f ${IMGNAME}:`cat "version.txt"`
