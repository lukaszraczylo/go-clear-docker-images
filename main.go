package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/fsouza/go-dockerclient"
)

var safeImages []string

func main() {
	var RemoveOlderThan time.Duration = 1209600 // time in seconds, default 2 weeks.
	imagesWhitelist := []string{"postgres", "ubuntu", "golang"}

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
	currentTime := time.Now().Unix()
	for _, img := range imgs {
		if float64(img.Created) < float64(currentTime)-(time.Second*RemoveOlderThan).Seconds() {
			var toRemove = true
			for _, rpt := range img.RepoTags {
				for _, protected := range imagesWhitelist {
					re, _ := regexp.Compile(protected)
					if re.Match([]byte(rpt)) {
						toRemove = false
					}
				}
				if toRemove || rpt == "<none>:<none>" {
					fmt.Printf("Trying to remove image: %s (%+v)\n", img.ID, img.RepoTags)
					client.RemoveImage(img.ID)
				}
			}
		}
	}
}
