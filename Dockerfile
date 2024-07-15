FROM golang:alpine AS build-env
RUN apk --no-cache add build-base gcc git ca-certificates
ENV GOPROXY=direct
ADD . /src
RUN cd /src && go build -o pg_api

# final stage
FROM alpine
WORKDIR /app

#graph db
ENV PG_URI=localhost:6234
ENV PG_USER=dbuser
ENV PG_PASSWORD=mypassword
ENV PG_DATA=widgetdb
ENV KEY =
ENV PORT=4000
ENV GIN_MODE=debug
COPY --from=build-env /src/pg_api /app/
EXPOSE 4000
CMD ["./pg_api"]