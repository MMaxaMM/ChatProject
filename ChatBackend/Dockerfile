FROM golang:1.23 AS build
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o ChatBackend cmd/chat-backend/main.go

FROM scratch
COPY --from=build /src/ChatBackend .
EXPOSE 5050
CMD [ "./ChatBackend" ]