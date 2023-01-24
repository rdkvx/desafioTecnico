FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server


FROM scratch
COPY --from=builder /app/server app/server
ENTRYPOINT [ "/app/server" ]