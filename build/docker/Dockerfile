FROM golang:latest AS build

WORKDIR /app

# download dependencies - low frequency of changes
COPY go.mod go.sum ./
RUN go mod download

# copy sources - higher frequency of changes
COPY internal/ internal/
COPY cmd/ cmd/

RUN CGO_ENABLED=0 GOOS=linux \
    go build \
      -ldflags="-w -s" \
      -installsuffix 'static' \
      -o ./wac-tg-jv-backend ./cmd/wac-tg-jv-backend

############################################
FROM scratch

COPY --from=build /app/wac-tg-jv-backend ./

EXPOSE 8080
ENTRYPOINT ["./wac-tg-jv-backend"]
