FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod init github.com/golangci_test \ 
    && go mod tidy && \
    go build -o math 

CMD ["./math"]