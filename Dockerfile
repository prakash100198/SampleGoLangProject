# Build Stage
# First pull Golang image
FROM golang:latest as builder 
 
 RUN mkdir /app
# ADD /public ./app
WORKDIR /app
# COPY /public /app
# RUN go build -o main .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main App.go
# CMD ["/app/main"]


# Set environment variable
# ENV APP_NAME sams
# ENV CMD_PATH main.go
 
# Copy application data into image
# COPY . ./
# WORKDIR $GOPATH/src/$APP_NAME
 
# # Budild application
# RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/Package/$CMD_PATH
 
# # Run Stage
# FROM alpine:3.14
 
# # Set environment variable
# ENV APP_NAME sample-dockerize-app
 
# # Copy only required data into this image
# COPY --from=build-env /$APP_NAME .
 
# # Expose application port
# EXPOSE 8081
 
# # Start app
# CMD ./$APP_NAME
