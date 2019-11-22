ARG BASEIMG="alpine:3.10"
ARG BUILDIMG="golang:1.13.4-alpine3.10"
FROM $BUILDIMG as builder

ENV GOPATH=""
ENV APP_NAME=immigrant

COPY . /go/

RUN cd /go && go build -o /${APP_NAME}

FROM $BASEIMG
LABEL maintainer="Nate Catelli <ncatelli@packetfire.org>"
LABEL description="Container for immigrant"

ENV SERVICE_USER "immigrant"
ENV APP_NAME="immigrant"

RUN addgroup ${SERVICE_USER} && \
    adduser -D -G ${SERVICE_USER} ${SERVICE_USER}

COPY --from=builder /${APP_NAME} /opt/${APP_NAME}/bin/

RUN chown -R ${SERVICE_USER}:${SERVICE_USER} /opt/${APP_NAME} && \
    chmod +x /opt/${APP_NAME}/bin/${APP_NAME}

WORKDIR "/opt/$APP_NAME/"
USER ${SERVICE_USER}

ENTRYPOINT [ "/opt/immigrant/bin/immigrant" ]
CMD [ "-h" ]
