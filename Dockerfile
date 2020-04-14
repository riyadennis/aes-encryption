FROM golang:alpine
RUN mkdir aes-encryption
ADD . /aes-encryption/
WORKDIR /aes-encryption
EXPOSE 8082
RUN go build -o aes-encryption .
CMD ["./aes-encryption"]