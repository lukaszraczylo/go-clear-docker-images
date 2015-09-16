# go-clear-docker-images

- Removes all stopped containers
- Removes all untagged images ( if possible )
- Preserves X of images tagged in the same way
- Uses whitelist to ignore images which you want to keep at all cost

### Requires following ENV variables:
```
    DOCKER_HOST : tcp://192.168.99.100:2376 | unix://var/run/docker.sock
    DOCKER_TLS_VERIFY : 1|0
    DOCKER_CERT_PATH : `pwd` | /home/potato
```

### Usage
```
  -older-than int
        Removes images older than X seconds. (default 1209600)
  -preserve int
        Numbers of images to preserve even if older than required (default 3)
  -whitelist string
        Whitelisted images, comma separated (default "postgres,ubuntu,golang")
```

##### Latest release available [here](https://github.com/lukaszraczylo/go-clear-docker-images/releases/latest).
