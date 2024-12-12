FROM golang:1.22.10-alpine AS build

WORKDIR /app

COPY ./hw13_http/go.mod ./hw13_http/go.sum ./hw13_http/
COPY ./config/go.mod ./config/
COPY ./go.work ./

RUN go mod download

COPY ./config/*.go ./config
COPY ./hw13_http ./hw13_http

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./hw13_http/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /client ./hw13_http/client/main.go


FROM alpine:latest AS server

WORKDIR /
COPY --from=build /server /server

ENTRYPOINT [ "/server" ]


FROM alpine:latest AS client

WORKDIR /
COPY --from=build /client /client
