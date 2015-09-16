package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/fsouza/go-dockerclient"
)

type Params struct {
	OlderThan int
	Whitelist string
	Preserve  int
	Debug     bool
}

// Requires following ENV variables:
// DOCKER_HOST :
// DOCKER_TLS_VERIFY : 1|0
// DOCKER_CERT_PATH : `pwd` | /home/potato
func main() {
	var p Params

	f := func(c rune) bool {
		return c == ':'
	}

	flag.IntVar(&p.OlderThan, "older-than", 2*7*86400, "Removes images older than X seconds.")
	flag.StringVar(&p.Whitelist, "whitelist", "postgres,ubuntu,golang", "Whitelisted images, comma separated")
	flag.IntVar(&p.Preserve, "preserve", 3, "Numbers of images to preserve even if older than required")
	flag.BoolVar(&p.Debug, "debug", false, "Print out what's going to be remove without touching stuff")
	flag.Parse()

	imagesWhiteList := strings.Split(p.Whitelist, ",")
	client, _ := docker.NewClientFromEnv()
	containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})
	fmt.Println("Checking containers...")
	for _, cnt := range containers {
		re, _ := regexp.Compile("Exited.*")
		if re.Match([]byte(cnt.Status)) {
			fmt.Printf("Removing container: %s (image: %s)\n", cnt.ID, cnt.Image)
			opts := docker.RemoveContainerOptions{ID: cnt.ID}
			if !p.Debug {
				client.RemoveContainer(opts)
			}
		}
	}

	fmt.Println("Checking images...")
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	currentTime := time.Now().Unix()
	var imagesArray []string
	for _, img := range imgs {
		if float64(img.Created) < float64(currentTime)-(time.Second*time.Duration(p.OlderThan)).Seconds() {
			var toRemove = true
			for _, rpt := range img.RepoTags {
				for _, protected := range imagesWhiteList {
					re, _ := regexp.Compile(protected)
					if re.Match([]byte(rpt)) {
						toRemove = false
					} else {
						imagesArray = append(imagesArray, strings.FieldsFunc(rpt, f)[0])
						if strings.Count(strings.Join(imagesArray, " "), rpt) < p.Preserve {
							toRemove = false
						}
					}
				}
				if toRemove || rpt == "<none>:<none>" {
					fmt.Printf("Trying to remove image: %s (%+v)\n", img.ID, img.RepoTags)
					if !p.Debug {
						client.RemoveImage(img.ID)
					}
				}
			}
		}
	}
}
