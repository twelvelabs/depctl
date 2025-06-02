FROM alpine:3.22

COPY depctl /usr/local/bin/depctl

CMD ["/usr/local/bin/depctl"]
