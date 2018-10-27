FROM golang:1.9
RUN bash -c "curl https://glide.sh/get | sh"
WORKDIR /go/src/github.com/BonexIO/go

COPY glide.lock /go/src/github.com/BonexIO/go
COPY glide.yaml /go/src/github.com/BonexIO/go
RUN glide install

COPY . .
RUN go install github.com/BonexIO/go/tools/...
RUN go install github.com/BonexIO/go/services/...