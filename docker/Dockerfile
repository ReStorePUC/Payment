
FROM alpine:latest

RUN apk update && apk add --no-cache libc6-compat
COPY ./payment /go/src/
COPY ./config.yaml /

CMD ["/go/src/payment"]