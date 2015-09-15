# go-clear-docker-images

- Removes all stopped containers
- Removes all untagged images ( if possible )

### Requires following ENV variables:
```
    DOCKER_HOST : lolcathost
    DOCKER_TLS_VERIFY : 1|0
    DOCKER_CERT_PATH : `pwd` | /home/potato
```

##### Check release for OSX binary ( if you don't want to compile it on your own )
