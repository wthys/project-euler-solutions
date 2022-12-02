FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN go get
RUN go build -o /out/euler

FROM scratch AS bin
COPY --from=build /out/euler /
