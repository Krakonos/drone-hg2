FROM alpine:3.2
RUN apk update && apk add ca-certificates mercurial git openssh curl perl bash && rm -rf /var/cache/apk/*
ADD drone-hg2 /bin/
ENTRYPOINT ["/bin/drone-hg2"]
