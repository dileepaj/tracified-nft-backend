FROM golang:1.18-alpine
RUN apk add --update cmake gcc g++ git  make  tar wget python3
RUN go version
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
RUN go build github.com/dileepaj/tracified-nft-backend
COPY . ./
RUN chmod +x tracified-nft-backend
CMD ["./tracified-nft-backend"]