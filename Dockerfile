FROM golang:1.19-alpine

# set the working directory
WORKDIR /app

# copy the source code to the container
COPY . .

# build the binary
RUN go build -o main

# run the binary
CMD ["./main"]