FROM golang:1.24.5-bookworm
LABEL authors="darius"

COPY . .

RUN go build .

ENTRYPOINT ["./kleinCommand"]