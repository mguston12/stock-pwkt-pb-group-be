FROM golang:alpine AS builder
RUN apk add --no-cache git make bash
RUN mkdir /go/src/skeleton
WORKDIR /go/src/skeleton
COPY . .
RUN make build

FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN apk --no-cache add tzdata
ENV TZ Asia/Jakarta
EXPOSE 8080
COPY --from=builder /go/src/skeleton/bin/skeleton /
COPY --from=builder /go/src/skeleton/files/etc/skeleton /
ENTRYPOINT ["/skeleton"]
