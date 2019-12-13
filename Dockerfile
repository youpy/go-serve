FROM golang:latest as builder
ENV GOPATH=
ADD serve.go ./
RUN go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o ./serve serve.go

FROM alpine
COPY --from=0 /go/serve ./
CMD ["./serve"]
