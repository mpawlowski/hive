# hive
Go packet metadata capture. Depends on `libpcap-dev`.

## build
```
go build
sudo ./hive -i <interface> &
sudo tail -f /var/log/hive/hive.log
```
