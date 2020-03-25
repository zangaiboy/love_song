# Start from the latest golang base image
FROM golang:latest AS builder
ENV GOPROXY=https://goproxy.cn,direct
# Add Maintainer Info
LABEL maintainer="zangaiboy@outlook.com"
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# Final image.
FROM scratch
COPY --from=builder /app .
EXPOSE 8000
CMD ["./main"]
