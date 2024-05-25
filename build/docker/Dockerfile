FROM golang:latest AS build

WORKDIR /app

# download dependencies - low frequency of changes
COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux \
    go build \
      -o ./wac-tg-jv-backend ./cmd/wac-tg-jv-backend

############################################
FROM scratch

COPY --from=build /app/wac-tg-jv-backend ./

EXPOSE 8080
ENTRYPOINT ["./wac-tg-jv-backend"]
