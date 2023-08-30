FROM golang:1.21-alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./

RUN go mod verify

COPY . .

RUN go build -o app

FROM alpine:3.14

RUN apk --no-cache add ca-certificates

COPY --from=build /app/app /app/app

CMD ["/app/app"]
