FROM golang:1.20

WORKDIR /src

COPY ./main.go ./
COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

RUN go build -o /docker-rest-api

EXPOSE 8000

CMD ["/docker-rest-api"]
