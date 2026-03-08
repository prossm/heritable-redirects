FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod main.go ./
RUN go build -o redirect .

FROM alpine:3.19
COPY --from=build /app/redirect /redirect
ENV PORT=8080
EXPOSE 8080
CMD ["/redirect"]
