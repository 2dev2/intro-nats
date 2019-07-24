# intro-nats
An Introduction to NATS

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
docker run -p 4222:4222 nats
```

References:
- https://nats.io/
- https://www.youtube.com/watch?v=rGI5J0b1deI&t=510s
- https://www.amazon.com/Practical-NATS-Beginner-Waldemar-Quevedo/dp/148423569X
