## Description

Homemade non persistent in memory Redis made in GO, for learning essentially : https://redis.io/docs/latest/


## Run redis

```
make run
```

## Test it

```
netcat 127.0.0.1 6379
```

This redis is RESP compilant so you can also use redis-cli :

```
redis-cli
```

## Commands implemented

- [PING]([url](https://redis.io/docs/latest/commands/ping/))
