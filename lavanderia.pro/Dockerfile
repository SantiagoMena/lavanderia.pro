FROM golang:1.20-alpine

WORKDIR /app/cmd

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
# RUN go mod download

COPY *.go ./

# RUN go build -o /lavanderia.pro

EXPOSE 1002

CMD [ "go", "run", "." ]