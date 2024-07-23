FROM golang:1.22.5
#set workdir
WORKDIR /app

#Update 
RUN apt-get update
RUN apt install -y protobuf-compiler


#download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

#Get protoc compiler
RUN  GO111MODULE=on \
        go install google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN export PATH="$PATH:$(go env GOPATH)/bin"


RUN protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        pkg/handlers/grpc/thumbnail/thumbnail.proto 




EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd

CMD [ "/bin/app" ]
