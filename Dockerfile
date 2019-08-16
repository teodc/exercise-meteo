FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o meteo .

EXPOSE 8080

ENTRYPOINT ["./meteo"]
