FROM golang:1.9-alpine
ADD . /go/src/gke-nodeapp
RUN go install gke-nodeapp
FROM alpine:latest
COPY --from=0 /go/bin/gke-nodeapp .
ENV PORT 8080
CMD ["./gke-nodeapp"]
