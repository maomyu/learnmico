FROM debian:latest
RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli
CMD ["./consignment-cli"]