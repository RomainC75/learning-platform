FROM golang:1.26.3
RUN go install github.com/air-verse/air@latest
COPY . /app
WORKDIR /app
ENTRYPOINT ["air"]