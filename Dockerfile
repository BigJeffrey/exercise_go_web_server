# The base go-image
FROM golang:1.17-alpine AS build

# Set working directory
WORKDIR /app

# Copy all files from the current directory to the app directory
COPY --chown=nobody ./ ./

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o webserver .

##############
##production##
##############

FROM alpine:3.15.0 AS app

# Set working directory
WORKDIR /app

COPY --from=build --chown=nobody /app/webserver /app/webserver
EXPOSE 8080
# Run the server executable
CMD [ "/app/webserver" ]
