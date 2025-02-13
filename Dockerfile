FROM golang:1.23.2-alpine AS builder

RUN apk update
RUN apk add git openssh tzdata build-base python3 net-tools

# Set the default value for the ENVIRONMENT build argument
ARG ENVIRONMENT=default
# Set the ENVIRONMENT environment variable to the value of the ENVIRONMENT build argument
ENV ENVIRONMENT=${ENVIRONMENT}



WORKDIR /app

COPY config.cold.json config.cold.json
COPY config.hot.json config.hot.json
COPY . .


RUN go install github.com/buu700/gin@latest
RUN GO111MODULE=auto
RUN go mod tidy
RUN go build main.go

FROM alpine:latest


RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    apk --no-cache add curl && \
    mkdir /app && \
    mkdir /database/

WORKDIR /app



COPY --from=builder /app /app


COPY . .

ENTRYPOINT ["sh","-c","/app/main"]