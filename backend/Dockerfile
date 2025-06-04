FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
