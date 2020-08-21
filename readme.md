### Cluster Sim

A project to learn about simulating http traffic to a Go server and load balancing with NGINX using a Raspberry Pi 4 cluster

#### Tech
- [Golang](https://golang.org/) for a simple http server with a few endpoints serving very basic JSON data
- [wrk](https://github.com/wg/wrk) - HTTP benchmarking
- [NGINX](https://nginx.org/en/) - Load balancing

#### Current setup
- Desktop - WSL2 running Ubuntu 20.04 LTS
- x2 Raspberry Pi 4 Model B 2GB - Raspberry Pi OS Lite 2020-05-27

#### Current progress
- NGINX installed on Pi 1
- Go server binary running on both Pis in the cluster

![gif demonstration of NGINX working on Pi cluster](https://i.imgur.com/lVXFgMB.gif)

#### Notes
- server.go build command on Windows to run on Pi - `env GOOS=linux GOARCH=arm GOARM=6 go build`
Full setup instructions for cloning and contributions to come
- NGINX Pi start `sudo /etc/init.d/nginx start`
