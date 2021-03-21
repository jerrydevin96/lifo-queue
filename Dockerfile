FROM busybox

RUN mkdir -p /application
WORKDIR /application

COPY lifo-queue /application

RUN chmod 777 lifo-queue

EXPOSE 8080

CMD [sh, "./lifo-queue"]