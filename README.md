# ali-ddns
ali-ddns api wapper


## Usage

### Install as a service
edit **/etc/systemd/system/ali-ddns.service**
```
[Unit]
Description=Ali-ddns
[Service]
Type=simple
PIDFile=/var/run/ali-ddns.pid
ExecStart=/path/to/your/file/ali-ddns server -a {bind_address} -p {bind_port} -i {AccessKeyId} -s {AccessKeySecret} -z {zone}
User=root
Group=root
[Install]
WantedBy=multi-user.target
```
