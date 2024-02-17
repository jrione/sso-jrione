FROM golang:1.21.4

WORKDIR /app/

COPY go.mod go.sum /app/
RUN go mod download

COPY ./ /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o ./app/gin-crud

EXPOSE 6969
CMD [ "./app/gin-crud" ]