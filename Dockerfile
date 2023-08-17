FROM golang:1.21-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o restaurantapi

CMD["/app/restaurantapi"]