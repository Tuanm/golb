ARG PKG=github.com/Tuanm/golb

FROM golang:1.19-alpine
WORKDIR /go/src/$PKG/
COPY . .
RUN go install && go build -o golb-server .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/$PKG/golb-server .
CMD ["./golb-server"]