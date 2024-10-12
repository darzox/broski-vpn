# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./



RUN go build -o /bot -v ./cmd/bot

##
## Deploy
##
FROM golang:1.22

WORKDIR /

COPY --from=build /bot /bot
COPY --from=build /app/migrations /migrations
EXPOSE 8085

ENTRYPOINT ["/bot"]
