FROM golang:alpine

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /app/cmd/product-http ./cmd/product-http
RUN apk add && apk add make

EXPOSE 8000

CMD [ "sh", "-c", "make db-init && make run-seed && /app/cmd/product-http/product-http"]