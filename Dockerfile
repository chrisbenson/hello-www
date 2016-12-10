FROM golang:alpine
RUN go get github.com/chrisbenson/hello-www
EXPOSE 8080
CMD ["hello-www"]
