FROM debian:9-slim

ADD ./orderservice app/bin/
WORKDIR /app

EXPOSE 8000

CMD [ "/app/bin/orderservice" ]