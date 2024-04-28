FROM golang:latest

USER root

WORKDIR /app

COPY / ./

ENV TODO_PORT="7540"
ENV TODO_DBFILE="/db/scheduler.db"
ENV TODO_PASSWORD="password"
ENV SECRET_KEY="my_secret_key"

EXPOSE 7540

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go_final_project
CMD ["/go_final_project"]