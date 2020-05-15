# build image
FROM golang:1.14 as build

WORKDIR /app
COPY . .
RUN env CGO_ENABLED=0 go build -trimpath -tags netgo -ldflags '-w -s -extldflags "-static"' -o extenso

# run image
FROM scratch

WORKDIR /app
COPY --from=build /app/extenso .
CMD ["./extenso"]
