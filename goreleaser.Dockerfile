FROM alpine:3.20.3
WORKDIR /
RUN apk add --no-cache tzdata
COPY lg-dev-mode /usr/local/bin
ENTRYPOINT ["lg-dev-mode"]
