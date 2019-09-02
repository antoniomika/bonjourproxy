Bonjour Proxy
=============

Creates a way to proxy bonjour services across subnets without broadcast access between them.

Setup
-----

Place a `services.toml` file within the directory of the binary (`/app` with docker).
An example of it is here:

```toml
[[proxyservice]]
name = "ecobee"
servicetype = "_hap._tcp"
domain = ""
port = 1200
host = "ecobee"
ip = "192.168.0.1"
textdata = [
    "MFG=ecobee Inc."
]
```

You can have multiple proxyservice entries within the service definition, they will all be proxied by this one process

Raspberry PI
------------

This project uses multiarch builds for docker to create images for different platforms. They can be accessed normally via the docker hub. I personally use this on a raspberry pi running home assistant.

An example:

```bash
docker run -itd \
    --restart=always \
    --net=host \
    --name=bonjourproxy \
    -v /mnt/data/supervisor/share/bonjourproxy/services.toml:/app/services.toml:ro \
    antoniomika/bonjourproxy:latest
```

CLI Flags
---------

```bash
Usage of bonjourproxy
  -services string
        The config file defining services to proxy (default "services.toml")
```
