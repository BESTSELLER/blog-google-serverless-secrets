FROM golang:1.15
WORKDIR $GOPATH/src/github.com/BESTSELLER/blog-google-serverless-secrets
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go install

CMD /go/bin/blog-google-serverless-secrets 
EXPOSE 8080