FROM golang:latest as builder

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /hello

EXPOSE 10000

CMD [ "/hello" ]