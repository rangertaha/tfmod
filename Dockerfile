# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/rangertaha/tfmod
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary.
# RUN go build -o /go/bin/tfmod
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/bin/tfmod

# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /usr/bin/tfmod /usr/bin/tfmod

# Expose tfmod on port 8080
EXPOSE 8080

# Run the tfmod binary.
ENTRYPOINT ["/go/bin/tfmod"]

# CMD ["tfmod", "server"]
