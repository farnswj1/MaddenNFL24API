FROM golang:1.22.2-alpine AS build

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./
RUN go build -o build -tags jsoniter

FROM scratch

EXPOSE 8080

COPY --from=build /opt/app/build /
COPY --from=build /opt/app/templates /templates/

ENTRYPOINT ["/build"]
