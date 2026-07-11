FROM golang:1.26
COPY . .
RUN go build -o server .
CMD ["./server"]
