FROM golang:1.24-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /my-app

EXPOSE 8080

CMD [ "/my-app" ]