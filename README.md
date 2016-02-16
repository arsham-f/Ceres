# Ceres

Ceres, pronounced "series," is a time-series data store, exposed via TCP. It's an abstraction of a key-value store where your key is always some increasing integer (i.e. a timestamp) and your values will be returned to you in ascending order of key. 

You can create "buckets" to store these keys. Each bucket has a lifetime. Entries which are too old will be pruned if their key is too old.

This is all built on Redis ZSET.


## Installation

```
go get "github.com/mediocregopher/radix.v2/redis"
go get "github.com/arsham-f/Ceres/ceres"
make
```

## Start up the server

```
./ceres-server -redis_addr 111.222.333.444:6379 -address 0.0.0.0 -port 1234
INFO: Waiting for connections on 127.0.0.1:7171
```

##### Options

* redis_addr: redis address to connect to (with port)
* redis_prot: redis protocol (default TCP)
* adddress: IP to bind TCP server to (default 127.0.0.1)
* port: TCP port to bind server to (default 7171)


## Connect to server

```
nc 127.0.0.1 1234
>CREATE my_bucket 86400 # Bucket with length of 86400 seconds (1 day)
Successfully created
>INSERT my_bucket 1455648831 this is some data
Successfully added
>INSERT my_bucket 1455648830 this is some older data
Successfully added
>GET my_bucket
[{"time":1455648830,"data":"this is some older data"},{"time":1455648831,"data":"this is some data"}]
```

##### Commands
`CREATE [bucket name] [length in seconds]` creates a new bucket

`DEL [bucket_name]` deletes a bucket

`INSERT [bucket_name] [time] [data]` inserts data into a bucket

`GET [bucket_name]` gets all data in a bucket
