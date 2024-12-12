FROM golang:1.22.10-alpine AS build

WORKDIR /app

COPY ./hw10_motion_sensor/go.mod ./hw10_motion_sensor/go.sum ./hw10_motion_sensor/
COPY ./config/go.mod ./config/
COPY ./go.work ./
RUN go mod download

COPY ./config/*.go ./config
COPY ./hw13_http ./hw131_http

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./hw131_http/server/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /client ./hw131_http/client/main.go


FROM alpine:latest AS server

WORKDIR /
COPY --from=build /server /server
EXPOSE ${APP_PORT}
ENTRYPOINT [ "/server" ]


FROM alpine:latest AS client

WORKDIR /
COPY --from=build /client /client
