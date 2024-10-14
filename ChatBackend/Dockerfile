FROM golang:1.23 AS build
COPY . /src
WORKDIR /src
ARG GOOS=linux
RUN CGO_ENABLED=0 GOOS=${GOOS} go build -o BackendServer cmd/main.go

FROM scratch
COPY --from=build /src/BackendServer .
EXPOSE 5050
CMD [ "/BackendServer" ]