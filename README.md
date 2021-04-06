## VM Portal
A simple web portal for managing virtual machines.
Developed in golang on Ubuntu 20.04.5 LTS.

## Version


## Install golang and libvirt-go
**Install golang**
```
root@NUC10:~# wget -q https://golang.org/dl/go1.16.3.linux-amd64.tar.gz
root@NUC10:~# tar -C /usr/local/ -xzf go1.16.3.linux-amd64.tar.gz
root@NUC10:~# export PATH=$PATH:/usr/local/go/bin
root@NUC10:~# go version
go version go1.16.3 linux/amd64
```
**install required packages**
```
apt install pkg-config libvirt-dev
```
**install libvirt-go**
```
root@NUC10:~# cd /usr/local/go/src/
root@NUC10:/usr/local/go/src# git clone -q https://github.com/libvirt/libvirt-go.git
root@NUC10:/usr/local/go/src# cd libvirt-go
root@NUC10:/usr/local/go/src/libvirt-go# go install
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

## License
Apache License 2.0
