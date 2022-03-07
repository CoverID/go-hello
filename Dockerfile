FROM golang:latest as builder

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /hello

EXPOSE 3000

CMD [ "/hello" ]