FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod init github.com/golangci_test
RUN go mod tidy 
RUN go build -o math 

CMD ["./math"]