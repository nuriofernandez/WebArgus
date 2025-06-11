##############################
# Build stage
##############################
FROM golang:1.21-alpine as build

### Copy only src files
WORKDIR /app
COPY ./ .

# Prepare env
RUN apk add git

### Build the binary
RUN go build -o ./argus .

##############################
# Final stage
##############################
FROM alpine

### Copy binary file from build stage
COPY --from=build /app/argus /usr/bin/argus
CMD ["argus"]