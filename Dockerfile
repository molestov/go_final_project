FROM golang:1.21.5

USER root

WORKDIR /app

COPY / ./

EXPOSE ${TODO_PORT}

RUN ${BUILD_PARAMS} go build -o /go_final_project
CMD ["/go_final_project"]