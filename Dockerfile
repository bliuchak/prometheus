FROM golang:1.17.2-alpine3.13 as builder
WORKDIR /go/src/app
RUN go get github.com/cespare/reflex
COPY . .
RUN go build -o /api -v

FROM scratch
COPY --from=builder /api /
USER 9000
ENTRYPOINT [ "/api" ]
