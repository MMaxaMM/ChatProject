FROM golang:1.23 AS build
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o Auth cmd/auth/main.go

FROM scratch
COPY --from=build /src/Auth .
COPY --from=build /src/.open.env .
EXPOSE 5040
CMD [ "/Auth" ]