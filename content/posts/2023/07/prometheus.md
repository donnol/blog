
## Question

prometheus start failed:

```sh
...
ts=2023-07-06T01:34:43.871Z caller=repair.go:57 level=info component=tsdb msg="Found healthy block" mint=1688493607159 maxt=1688515200000 ulid=01H4J6TB6NCBZFNR9XZ1R2P67H
ts=2023-07-06T01:34:43.871Z caller=repair.go:57 level=info component=tsdb msg="Found healthy block" mint=1688529607159 maxt=1688536800000 ulid=01H4JDM4YR6J3TJVBY6P6EGZS4
ts=2023-07-06T01:34:43.872Z caller=repair.go:57 level=info component=tsdb msg="Found healthy block" mint=1688536807159 maxt=1688544000000 ulid=01H4JMFGKWY1MAM7GBFDQ89FRV
ts=2023-07-06T01:34:43.872Z caller=main.go:696 level=info msg="Stopping scrape discovery manager..."
ts=2023-07-06T01:34:43.872Z caller=main.go:710 level=info msg="Stopping notify discovery manager..."
ts=2023-07-06T01:34:43.872Z caller=main.go:732 level=info msg="Stopping scrape manager..."
ts=2023-07-06T01:34:43.872Z caller=manager.go:946 level=info component="rule manager" msg="Stopping rule manager..."
ts=2023-07-06T01:34:43.872Z caller=manager.go:956 level=info component="rule manager" msg="Rule manager stopped"
ts=2023-07-06T01:34:43.872Z caller=notifier.go:601 level=info component=notifier msg="Stopping notification manager..."
ts=2023-07-06T01:34:43.872Z caller=main.go:726 level=info msg="Scrape manager stopped"
ts=2023-07-06T01:34:43.872Z caller=main.go:692 level=info msg="Scrape discovery manager stopped"
ts=2023-07-06T01:34:43.872Z caller=main.go:907 level=info msg="Notifier manager stopped"
ts=2023-07-06T01:34:43.872Z caller=main.go:706 level=info msg="Notify discovery manager stopped"
ts=2023-07-06T01:34:43.872Z caller=tls_config.go:195 level=info component=web msg="TLS is disabled." http2=false
ts=2023-07-06T01:34:43.872Z caller=main.go:916 level=error err="opening storage failed: get segment range: segments are not sequential"
```

## Answer

```sh
rm -rf /var/lib/prometheus/metrics2/
```

Path is from `prometheus -h | grep '--storage.tsdb.path'`.

