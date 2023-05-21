FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

COPY . .

RUN go build -o /appparser cmd/main.go

EXPOSE ${PORT}

CMD ["/appparser"]