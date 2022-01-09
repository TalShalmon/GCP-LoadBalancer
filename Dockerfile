FROM golang:1.9-alpine
ADD . /go/src/gke-nodeapp
WORKDIR /app
COPY . ./usr/local/go/src/gke-nodeapp
#COPY /go/bin/GKE-nodeapp .

COPY go.* /usr/local/
#RUN go mod downlaod
#RUN go get -d -v ./
#RUN go install -v ./
RUN go install gke-nodeapp
# go build -o /GKE-nodeapp 
# go get gopkg.in/mgo.v2 && \
 #go get gopkg.in/mgo.v2/bson && \	
# go get github.com/gin-gonic/gin 
WORKDIR /go
FROM alpine:latest
ENV GO111MODULE=on
COPY --from=0 /go/bin/gke-nodeapp .
CMD ["./gke-nodeapp"]
#WORKDIR /go/src/easycast/src
#COPY src /go
#EXPOSE 27017
#CMD ["go", "run", "main.go"]

