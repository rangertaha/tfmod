
############################# Stage 1
FROM golang:alpine AS builder

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/github.com/rangertaha/tfmod
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary.
# RUN go build -o /go/bin/tfmod
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/tfmod

############################### Stage 2
#FROM golang:latest
## Copy our static executable.
#COPY --from=builder /go/bin/tfmod /go/bin/tfmod
#
## Expose tfmod on port 8080
#EXPOSE 8080
#
## Run the tfmod binary.
##ENTRYPOINT ["/tfmod"]

CMD ["/go/bin/tfmod", "server"]
