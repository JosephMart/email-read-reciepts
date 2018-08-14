FROM golang:1.10 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
COPY . $GOPATH/
WORKDIR $GOPATH/src/receipts
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .
RUN mv $GOPATH/img.png /img.png

FROM scratch
COPY --from=builder /app ./
COPY --from=builder /img.png ./
ENTRYPOINT ["./app"]
