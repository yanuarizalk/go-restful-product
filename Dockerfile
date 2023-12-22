FROM telkomindonesia/alpine:go-1.19 as builder

RUN apk update && apk add --no-cache git
RUN apk add --no-cache tzdata
RUN apk add build-base

WORKDIR /src/go-restful-test
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN GIT_TERMINAL_PROMPT=0 go mod download

COPY . .
RUN go build -v -o product .


# Final Image
# ---------------------------------------------------
FROM alpine:3.16.0

RUN apk add --no-cache ca-certificates openssl
RUN apk add --no-cache tzdata

WORKDIR /app
COPY --from=builder /src/go-restful-test/product .

CMD ["./product"]