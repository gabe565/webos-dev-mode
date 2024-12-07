FROM alpine:3.20.3
WORKDIR /
COPY lg-dev-mode /usr/local/bin
ENTRYPOINT ["lg-dev-mode"]
