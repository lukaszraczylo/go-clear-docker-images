package main

import (
	"fmt"
	"regexp"

	"github.com/fsouza/go-dockerclient"
)

var safeImages []string

func main() {
	// Requires following ENV variables:
	// DOCKER_HOST :
	// DOCKER_TLS_VERIFY : 1|0
	// DOCKER_CERT_PATH : `pwd` | /home/potato
	client, _ := docker.NewClientFromEnv()
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})
	fmt.Println("Checking containers...")
	for _, cnt := range containers {
		re, _ := regexp.Compile("Exited.*")
		if re.Match([]byte(cnt.Status)) {
			fmt.Printf("Removing container: %s (image: %s)\n", cnt.ID, cnt.Image)
			opts := docker.RemoveContainerOptions{ID: cnt.ID}
			client.RemoveContainer(opts)
		}
	}

	fmt.Println("Checking images...")
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		for _, rpt := range img.RepoTags {
			if rpt == "<none>:<none>" {
				fmt.Printf("Trying to remove image: %s (%+v)", img.ID, img.RepoTags)
				client.RemoveImage(img.ID)
			}
		}
	}
}
