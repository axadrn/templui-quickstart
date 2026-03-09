FROM golang:1.24-alpine AS build
WORKDIR /app

COPY . .

RUN apk add --no-cache gcc musl-dev wget

RUN arch="$(apk --print-arch)" && \
    case "$arch" in \
      x86_64) tailwind_arch="x64" ;; \
      aarch64) tailwind_arch="arm64" ;; \
      *) echo "unsupported architecture: $arch" && exit 1 ;; \
    esac && \
    wget -O /usr/local/bin/tailwindcss "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-${tailwind_arch}" && \
    chmod +x /usr/local/bin/tailwindcss

RUN TEMPLUI_PATH="$(go list -m -f '{{.Dir}}' github.com/templui/templui)" && \
    printf '%s\n' \
      '@source "./**/*.templ";' \
      "@source \"$TEMPLUI_PATH/components/**/*.templ\";" \
      > ./assets/css/sources.generated.css && \
    tailwindcss -i ./assets/css/input.css -o ./assets/css/output.css --minify

RUN go tool templ generate
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./main.go

FROM alpine:3.20.2
WORKDIR /app

RUN apk add --no-cache ca-certificates

ENV GO_ENV=production

COPY --from=build /app/main .

EXPOSE 8090

CMD ["./main"]
