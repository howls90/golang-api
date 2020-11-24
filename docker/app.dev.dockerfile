FROM golang

ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /src

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN go get -u github.com/mitranim/gow

# Download all the dependencies
RUN go mod download

# Run the executable
CMD ["/go/bin/gow", "run", "main.go"]