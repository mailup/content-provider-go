FROM alpine:latest

MAINTAINER Massimo Zerbini <zerbini@mailup.com>

WORKDIR "/opt"

ADD .docker_build/content-provider-go /opt/bin/content-provider-go

CMD ["/opt/bin/content-provider-go"]

