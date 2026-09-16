FROM golang:1.22.2
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main.exe
RUN chmod +x main.exe
EXPOSE 3000
CMD ["./main.exe"]
