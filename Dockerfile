FROM golang:1.19-alpine
 
RUN mkdir /app
 
COPY . /app
 
WORKDIR /app
 
EXPOSE 5050
RUN go build -o reshare-service . 
 
CMD [ "/app/reshare-service" ]
