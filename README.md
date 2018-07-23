# iptables-healthchecker
It is a Web API that can health check the iptables.

## Description
Provide WebAPI to check if iptables is running normally on CentOS 7.

## Installation/Usage

`go build` and just launch the output binary.

```
$ git clone https://github.com/HirokiSakonju/iptables-healthchecker.git
$ env GOOS=linux GOARCH=amd64 go build iptables_checker.go
$ ./iptables_ckecker 
```
