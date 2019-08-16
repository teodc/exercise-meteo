# Meteo

> WIP

Go exercice based on the one on howistart.org

## Usage

```
cp .env.dist .env
vim .env
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
- Add comments
- ...
