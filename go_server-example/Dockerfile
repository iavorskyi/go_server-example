FROM golang:1.16-alpine as build

RUN apk add --no-cache git && \
    printf "machine github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n\
    \n\
    machine api.github.com\n\
    login ${ACCESS_TOKEN_USR}\n\
    password ${ACCESS_TOKEN_PWD}\n"\
    >> /root/.netrc && \
    chmod 600 /root/.netrc

WORKDIR /app

COPY . .
RUN GOOS=linux go build -o go_server-example

FROM alpine
WORKDIR /app
COPY --from=build /app/go_server-example .
CMD ["./go_server-example"]