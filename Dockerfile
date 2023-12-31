FROM golang:1.21

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

ENV PORT=3000

# Copy local code to the container image.
COPY . ./

# Build the binary
RUN make build 

CMD [ "./go-mini-kv-store" ]
