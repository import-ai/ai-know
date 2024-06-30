#!/bin/bash

export API_JWT_SECRET_KEY=123
export API_LISTEN_ADDR="127.0.0.1:5433"
export API_DATA_SOURCE_NAME='host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai'
go run main.go
