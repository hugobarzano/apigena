# Dockerfile2 from https://github.com/hugobarzano/LProductosSoftware.git
#FROM golang:latest
FROM golang:onbuild
RUN mkdir /gen
ADD ./gen /gen/
WORKDIR /gen
#RUN go build -o main cmd/gen/main.go
CMD ["/gen/main"]
