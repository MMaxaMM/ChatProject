FROM golang:1.23 AS build
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o BackendServer cmd/main.go

FROM scratch
COPY --from=build /src/BackendServer .
COPY --from=build /src/.env .
EXPOSE 5050
CMD [ "/BackendServer" ]