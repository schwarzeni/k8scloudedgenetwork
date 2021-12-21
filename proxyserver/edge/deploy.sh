 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o proxyedge .
 scp proxyedge root@10.211.55.52:~/
 scp proxyedge root@10.211.55.65:~/
 scp proxyedge root@10.211.55.66:~/
