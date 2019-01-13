#build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o app cmd/web/*

# iron/go is the alpine image with only ca-certificates added
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ["./app"]