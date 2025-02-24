# build stage
FROM golang:1.22.3-alpine as build-stage
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/sitepulse ./cmd

# production stage
FROM alpine:3.19.1 as production-stage
RUN addgroup -g 1000 sitepulse && adduser -h /home/sitepulse -s /bin/sh -G sitepulse -D -u 1000 sitepulse
USER 1000:1000
COPY --from=build-stage --chown=1000:1000 --chmod=0770 /usr/local/bin/sitepulse /usr/local/bin/sitepulse
EXPOSE 3000
CMD ["sitepulse"]