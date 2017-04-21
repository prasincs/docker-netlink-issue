# Docker tunnel issue with 17.04 & Ubuntu 14.04.5


## Find a machine with Ubuntu 14.04.5 & docker 17.04

```
# uname -a

Linux hostname-redacted 3.13.0-112-generic #159-Ubuntu SMP Fri Mar 3 15:26:07 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
```

```

# docker version
Client:
 Version:      17.04.0-ce
 API version:  1.28
 Go version:   go1.7.5
 Git commit:   4845c56
 Built:        Mon Apr  3 18:01:08 2017
 OS/Arch:      linux/amd64

Server:
 Version:      17.04.0-ce
 API version:  1.28 (minimum version 1.12)
 Go version:   go1.7.5
 Git commit:   4845c56
 Built:        Mon Apr  3 18:01:08 2017
 OS/Arch:      linux/amd64
 Experimental: false
```


This is likely to happen in other versions too but this is the one I've tried so far, against Linux `3.13.0-112` and `3.13.0-116`

### Run the command


```
sudo docker run --privileged --rm golang:1.8 sh -c "go get github.com/prasincs/docker-netlink-issue/... && exec docker-netlink-issue"
```

Output:

```
2017/04/21 19:47:22 ip addr dev tun0 fd28:49be:758:7653:3cb1:6d:230c:60/64
2017/04/21 19:47:22 Failed to add fd28:49be:758:7653:3cb1:6d:230c:60/64 to dev "tun0": permission denied
<nil>
2017/04/21 19:47:22 Error opening tun. permission denied
```

When you downgrade to `docker 17.03` it works like a charm

```
sudo docker run --privileged --rm golang:1.8 sh -c "go get github.com/prasincs/docker-netlink-issue/... && exec docker-netlink-issue"
2017/04/21 19:48:08 ip addr dev tun0 fd28:49be:758:7653:3cb1:6d:230c:60/64
&{tun0 0xc42000e038}
```


The code works fine outside of docker.


Adding the following options for dockerd doesn't help either

```
DOCKER_OPTS="--ipv6=true --fixed-cidr-v6=fd28:49be::0/120"
```
