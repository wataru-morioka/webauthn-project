FROM  golang:1.16-rc-alpine3.12

COPY ./.aws /root/.aws

WORKDIR /root
RUN mkdir app
WORKDIR /root/app

RUN go mod init webauthn-project && \
    go get -u github.com/cosmtrek/air && \
    go get github.com/99designs/gqlgen && \
    export GOPATH=$HOME/go && \
    export PATH=$PATH:$GOPATH/bin   

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]