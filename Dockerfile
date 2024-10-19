FROM alpine:latest
LABEL org.opencontainers.image.source https://github.com/pteich/clai

COPY clai /usr/bin/clai

ENTRYPOINT ["/usr/bin/clai"]