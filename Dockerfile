FROM golang:latest

# Metadata
LABEL maintainer="Aymen <aymen.rebouh@gmail.com>"

# Indicate the working directory
WORKDIR /app

# Copy the list of dependencies
COPY go.mod .
# .. and hash
COPY go.sum .

# Install all the dependencies
RUN go mod download

# Copy all the files to the container
COPY . .

ENV GO_PRACTICE_PORT 8000

RUN go build

CMD [ "./go-practice" ]