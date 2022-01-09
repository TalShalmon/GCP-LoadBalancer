FROM golang:1.9-alpine
ADD . /go/src/gke-nodeapp
WORKDIR /app
COPY . ./usr/local/go/src/gke-nodeapp
COPY go.* /usr/local/
RUN go install gke-nodeapp
WORKDIR /go
FROM alpine:latest
ENV GO111MODULE=on
COPY --from=0 /go/bin/gke-nodeapp .
CMD ["./gke-nodeapp"]
