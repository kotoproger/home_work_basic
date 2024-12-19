FROM golang:1.22.10-alpine AS build

WORKDIR /app

COPY ./hw02_fix_app/go.mod ./hw02_fix_app/go.sum ./hw02_fix_app/
COPY ./hw03_chessboard/go.mod ./hw03_chessboard/go.sum ./hw03_chessboard/
COPY ./hw04_struct_comparator/go.mod ./hw04_struct_comparator/go.sum ./hw04_struct_comparator/
COPY ./hw05_shapes/go.mod ./hw05_shapes/go.sum ./hw05_shapes/
COPY ./hw06_testing/go.mod ./hw06_testing/go.sum ./hw06_testing/
COPY ./hw07_word_counter/go.mod ./hw07_word_counter/go.sum ./hw07_word_counter/
COPY ./hw08_binary_search/go.mod ./hw08_binary_search/go.sum ./hw08_binary_search/
COPY ./hw09_serialize/go.mod ./hw09_serialize/go.sum ./hw09_serialize/
COPY ./hw10_motion_sensor/go.mod ./hw10_motion_sensor/go.sum ./hw10_motion_sensor/
COPY ./hw11_worker_pool/go.mod ./hw11_worker_pool/go.sum ./hw11_worker_pool/
COPY ./hw12_log_util/go.mod ./hw12_log_util/go.sum ./hw12_log_util/
COPY ./hw13_http/go.mod ./hw13_http/go.sum ./hw13_http/
COPY ./hw15_go_sql/go.mod ./hw15_go_sql/go.sum ./hw15_go_sql/
COPY ./config/go.mod ./config/
COPY ./go.work ./

RUN go mod download

COPY ./config/*.go ./config
COPY ./hw15_go_sql ./hw15_go_sql

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./hw15_go_sql/main.go


FROM alpine:latest AS server

WORKDIR /
COPY --from=build /server /server
COPY --from=build /app/hw15_go_sql/sql /sql

ENTRYPOINT [ "/server" ]