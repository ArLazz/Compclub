FROM golang:latest

COPY ./ ./
RUN go mod download
RUN go build -o /compclub 

CMD ["/compclub", "tests/test_file_1.txt"]
