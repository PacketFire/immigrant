ARG BASEIMG="alpine:3.7"
ARG BUILDIMG="golang:1.10.1-alpine3.7"
FROM $BUILDIMG as builder

ENV GOPATH=/go
ENV GIT_USER=ncatelli
ENV SCM_PROVIDER=github.com
ENV APP_NAME=immigrant

COPY . /go/src/${SCM_PROVIDER}/${GIT_USER}/${APP_NAME}/

RUN cd ${GOPATH} && go build -o /${APP_NAME} ${SCM_PROVIDER}/${GIT_USER}/${APP_NAME}

FROM $BASEIMG
LABEL maintainer="Nate Catelli <ncatelli@packetfire.org>"
LABEL description="Container for immigrant"

ENV SERVICE_USER "immigrant"
ENV APP_NAME=immigrant

RUN addgroup ${SERVICE_USER} && \
    adduser -D -G ${SERVICE_USER} ${SERVICE_USER}

COPY --from=builder /${APP_NAME} /opt/${APP_NAME}/bin/

RUN chown -R ${SERVICE_USER}:${SERVICE_USER} /opt/${APP_NAME} && \
    chmod +x /opt/${APP_NAME}/bin/${APP_NAME}

WORKDIR "/opt/$APP_NAME/"
USER ${SERVICE_USER}

ENTRYPOINT [ "/opt/immigrant/bin/immigrant" ]
CMD [ "-h" ]