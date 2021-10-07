# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /riddle

CMD [ "/riddle" ]