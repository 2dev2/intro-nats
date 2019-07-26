# An Introduction to NATS

## Try Using Telnet
You can try it right away using telnet onto demo server.

```
telnet demo.nats.io 4222
INFO {"server_id":"EiRJABZmVpWQDpriVqbbtw",..., "max_payload":1048576}
SUB greetings 1
+OK
PUB greetings 12
Hello World!
+OK
MSG greetings 1 12
Hello World!
```

or you can run it using docker

```
docker run --name -p 4222:4222 nats
```

## Try Using Simple Subscriber
- Run your nats server `docker run --name -p 4222:4222 nats`
- make run-simple

## Try Using anime-service and search-service
- Run your nats server `docker run --name -p 4222:4222 nats`
- Run your elasticsearch server `docker run -d --name elasticsearch --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:6.8.1`
- make subscriber
- make search-service
- make anime-service
- make es-mapping
- Try It!

## TODO:
- Make docker-compose for easier demo

## References:
- https://nats.io/
- https://www.youtube.com/watch?v=rGI5J0b1deI&t=510s
- https://www.amazon.com/Practical-NATS-Beginner-Waldemar-Quevedo/dp/148423569X
