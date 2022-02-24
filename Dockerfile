FROM golang:1.17-alpine

# Set destination for COPY
RUN mkdir -p /go/src/github.com/dileepaj/tracified-nft-backend/
WORKDIR /go/src/github.com/dileepaj/tracified-nft-backend/

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY . ./

# Build
RUN CGO_ENABLED=0 go build github.com/dileepaj/tracified-nft-backend
COPY . ./
RUN chmod +x tracified-nft-backend
CMD ["./tracified-nft-backend"]