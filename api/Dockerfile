FROM golang:1.22-alpine AS development

RUN mkdir /app

ADD go.mod go.sum /app/

WORKDIR /app

RUN go install golang.org/x/tools/gopls@latest

RUN go mod download

COPY . /app

RUN GOOS=linux go build -ldflags="-s -w" -o main main.go

FROM alpine AS production

LABEL "maintainer"="thuongtruong109"
LABEL "org.opencontainers.image.authors"="Tran Nguyen Thuong Truong <thuongtruongofficial@gmail.com>"
LABEL "org.opencontainers.image.version"="1.0"
LABEL "org.opencontainers.image.description"="Official image of onelink api"
LABEL "org.opencontainers.image.licenses"="MIT"
LABEL "org.opencontainers.image.source"="https://github.com/thuongtruong109/onelink/api"
LABEL "org.opencontainers.image.documentation"="https://github.com/thuongtruong109/onelink/blob/main/README.md"

RUN apk --no-cache add ca-certificates

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=development /app/main /app/

WORKDIR /app

EXPOSE 8080

CMD ["./main"]