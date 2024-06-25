FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV PORT :8000

RUN CGO_ENABLED=1 GOOS=linux go build -o /adote-um-idoso-backend

CMD ["/adote-um-idoso-backend"]