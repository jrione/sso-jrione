FROM golang:1.21.4-alpine3.18 as build
WORKDIR /app/
COPY ./ /app/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/gin-crud

FROM scratch
COPY --from=build /app/gin-crud /app/gin-crud
COPY ./config.json /
EXPOSE 8000
CMD [ "./app/gin-crud" ]