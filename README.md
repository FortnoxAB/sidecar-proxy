# sidecar-proxy
Small proxy we use for exposing /metrics on specific kubernetes components. 
Only proxies specific path for security. So we can expose for example kube-controller-manager metrics.


### Example run
```
sidecar-proxy -listen-addr=10.1.1.1:10251 -proxy-addr=http://127.0.0.1:10251/metrics
```

