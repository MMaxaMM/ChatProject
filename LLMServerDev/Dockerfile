FROM golang:1.23 AS build
COPY . /src
WORKDIR /src
ARG GOOS
ENV GOOS $GOOS
RUN CGO_ENABLED=0 GOOS=$GOOS go build -o LLMServer cmd/main.go

FROM scratch
COPY --from=build /src/LLMServer .
EXPOSE 5000
CMD [ "/LLMServer" ]