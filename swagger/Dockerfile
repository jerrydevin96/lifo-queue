FROM quay.io/goswagger/swagger

RUN mkdir -p /swagger-client
WORKDIR /swagger-client

COPY swagger.json /swagger-client

EXPOSE 8083

CMD [swagger, "serve", " --flavor swagger --no-open --base-path=/swagger --port=8083 swagger.json"]