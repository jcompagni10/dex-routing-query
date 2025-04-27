FROM golang:1.21-bullseye

WORKDIR /app

# Install SQLite development files
RUN apt-get update && apt-get install -y \
    gcc \
    libsqlite3-dev \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod ./
RUN go mod download

COPY . .

# Enable CGO and build
ENV CGO_ENABLED=1
RUN go build -o /app/main

EXPOSE 8080

CMD ["/app/main"] 