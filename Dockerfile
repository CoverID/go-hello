FROM golang:latest as builder
COPY . /app/go
WORKDIR /app/go
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o go
#second stage
FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache tzdata
COPY --from=builder /app/go .

EXPOSE 8086

CMD ["./go"]