# hive
Network packet metadata logging. Depends on [libpcap-dev](https://packages.debian.org/search?keywords=libpcap-dev).

## build
```
go build
sudo ./hive -i <interface> &
sudo tail -f /var/log/hive/hive.log
```
