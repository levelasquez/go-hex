# Hexagonal Architecture Example Using GO

## build project

`go build main.go`

## run project with redis

default run with redis

`./main`

## run project with mongo

`./main --database mongo`

## create example

```
curl -X POST http://localhost:3000/campaigns -H 'Cache-Control: no-cache' -H 'Content-Type: application/json' -d '{
  "name": "test campaign",
  "code": "TEST;0;0;INTERNET",
  "method": "CMR",
  "start": "2020-02-26",
  "end": "2020-02-29"
}'
```

## find all example

```
curl http://localhost:3000/campaings
```

## get by id example

```
curl http://localhost:3000/campaings/id
```
