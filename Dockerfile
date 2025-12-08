FROM alpine:3.23

COPY depctl /usr/local/bin/depctl

CMD ["/usr/local/bin/depctl"]
