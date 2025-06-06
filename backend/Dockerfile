FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

ENV PATH=$PATH:/go/bin

RUN go install github.com/air-verse/air@latest

CMD ["air"]
