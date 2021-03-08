FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /root/app cmd/main.go
FROM alpine
COPY --from=build /root/app /

# Command to run the executable
CMD ["/app"]