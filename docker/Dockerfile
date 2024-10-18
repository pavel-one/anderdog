FROM golang:alpine AS builder
ENV CGO_ENABLED=1
WORKDIR /app
COPY . .
RUN apk add --no-cache --update git build-base
RUN go mod tidy \
	&& go build -o app .


FROM alpine:latest AS runner
ENV TZ=Europe/Moscow
RUN apk --no-cache add ca-certificates tzdata libc6-compat libgcc libstdc++
WORKDIR /app

COPY --from=builder /app/app .

ENTRYPOINT ["./app"]