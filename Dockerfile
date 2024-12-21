FROM alpine:3.21

COPY depctl /usr/local/bin/depctl

CMD ["/usr/local/bin/depctl"]
