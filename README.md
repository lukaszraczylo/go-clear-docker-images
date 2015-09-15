# go-clear-docker-images

- Removes all stopped containers
- Removes all untagged images ( if possible )
- Preserves X of images tagged in the same way
- Uses whitelist to ignore images which you want to keep at all cost

### Requires following ENV variables:
```
    DOCKER_HOST : lolcathost
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

##### Check release for OSX binary ( if you don't want to compile it on your own )
