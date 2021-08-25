# https://docs.docker.com/language/golang/build-images/

FROM golang:alpine3.14

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

# This "myapp" can be any name
# RUN go build -o /myapp
RUN go build -o /myapp

EXPOSE 8081

CMD [ "/myapp" ]
