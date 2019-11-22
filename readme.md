# Meteo

> WIP

Tells you the weather in a given city.

> Exercice based on the one on howistart.org

## Usage

1. Prepare the `.env` & add your Open Weather API key:

```
cp .env.dist .env
vim .env
```

2. Run the app:

```
go build
./meteo
curl http://localhost:8080/weather/tokyo
```

or

```
docker build -t meteo .
docker run --rm -it -p 8080:8080 meteo
curl http://localhost:8080/weather/tokyo
```

## ToDo

- Add tests
- Error handling
- Add new provider
