FROM golang:1.24-alpine3.22 AS build
WORKDIR /app/
COPY ./ /app/
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o /app/sso-jrione

FROM gcr.io/distroless/base
COPY --from=build /app/sso-jrione /app/sso-jrione
COPY ./config.json /

EXPOSE 8000
CMD [ "./app/sso-jrione" ]
