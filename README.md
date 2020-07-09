# ali-ddns
ali-ddns api wapper


## Usage

- provide the key when starting the program

`/path/to/your/file/ali-ddns server -a {bind_address} -p {bind_port} -z {zone}`
then
visit
`http://ip:port/update?domain=your.domian&record=example&address=8.8.8.8`

- provide key when visiting the web api

`/path/to/your/file/ali-ddns server -a {bind_address} -p {bind_port} -i {AccessKeyId} -s {AccessKeySecret} -z {zone}`
then
visit
`http://ip:port/update?domain=your.domian&record=example&address=8.8.8.8&id=yoru_access_key_id&sec=your_access_key_secret`

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
