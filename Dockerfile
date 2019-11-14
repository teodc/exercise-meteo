FROM golang:1.12

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o meteo .

EXPOSE 8080

ENTRYPOINT ["./meteo"]
