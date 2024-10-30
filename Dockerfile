# ใช้ golang image เป็น base image (runtime เริ่มต้น)
FROM golang:latest

# กำหนด path /app เป็น directory เริ่มต้น
WORKDIR /app

# Copy ไฟล์ go.mod และ go.sum เข้าไป ที่ /app
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy ไฟล์ทั้งหมดจาก current directory สู่ Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable # ทำงานเมื่อมีการใช้ Container หรือ Docker run
CMD ["./main"]