# hive
Network packet metadata logging. Depends on [libpcap-dev](https://packages.debian.org/search?keywords=libpcap-dev).

## quickstart

add logstash filter
```
filter {
  json {
    source => "message"
  }
}
```

send packets to logstash
```
go build
sudo ./hive -i <interface> | nc <host> <port>
```
