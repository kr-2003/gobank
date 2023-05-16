FROM golang:1.20.0

WORKDIR /app

COPY . /app

RUN go mod tidy

EXPOSE 3000

CMD ["make", "run"]