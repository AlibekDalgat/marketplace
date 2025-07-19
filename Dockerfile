FROM golang:1.23.1

RUN go version

COPY . /marketplace/
WORKDIR /marketplace/

RUN apt-get update && apt-get -y install postgresql-client

RUN go mod tidy
RUN go mod download
RUN GOOS=linux go build -o ./bin/app ./cmd/main.go

RUN sed -i -e 's/\r$//' *.sh
RUN chmod +x wait-for-postgres.sh

CMD ["./bin/app"]