# hive
Network packet metadata logging. Depends on `libpcap-dev`.

## build
```
go build
sudo ./hive -i <interface> &
sudo tail -f /var/log/hive/hive.log
```
