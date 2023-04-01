FROM golang:1.20-alpine

WORKDIR /app/cmd/lavanderia

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
# RUN go mod download

COPY *.go ./

# RUN go build -o /lavanderia.pro

EXPOSE 8080

CMD [ "go", "run", "." ]