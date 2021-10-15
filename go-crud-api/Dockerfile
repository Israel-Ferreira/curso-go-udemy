FROM golang:1.17


ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod download


COPY . .

RUN GOOS=linux go build

EXPOSE 5000


ENTRYPOINT [ "./go-crud-api" ]