FROM golang:1.19.0-alpine AS build

WORKDIR /usr/src/app

COPY . .
RUN go build ./...

FROM alpine
WORKDIR /app
COPY --from=build /usr/src/app/receipt_processor /app

EXPOSE 8080

ENTRYPOINT ["./receipt_processor"]
