
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /pokemon-m-v1

EXPOSE 8080

CMD [ "/pokemon-m-v1" ]