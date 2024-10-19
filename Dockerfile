# Build.
FROM golang:1.23 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

COPY . /app
RUN CGO_ENABLED=1 GOOS=linux go build -o /main cmd/main.go

# Deploy.
FROM ubuntu:latest 
WORKDIR /
COPY --from=build-stage /main /main
COPY --from=build-stage /app/assets /assets
EXPOSE 8080
ENTRYPOINT ["/main"]
