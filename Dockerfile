FROM golang:1.10 as builder
WORKDIR /go/src/github.com/klstr/reflector
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure
RUN make

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
EXPOSE 3000
COPY --from=builder /go/src/github.com/klstr/reflector/reflector .
CMD ["./reflector"]
