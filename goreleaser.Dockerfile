FROM alpine:3.21.3
WORKDIR /
RUN apk add --no-cache tzdata
COPY webos-dev-mode /usr/local/bin
ENTRYPOINT ["webos-dev-mode"]
