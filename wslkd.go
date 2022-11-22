package main

import (
	"flag"
	"github.com/xiaoqidun/gowsl"
	"log"
	"strings"
	"sync"
)

func main() {
	var d string
	flag.StringVar(&d, "d", "Debian", "distribution")
	flag.Parse()
	waitGroup := &sync.WaitGroup{}
	distroList := strings.Split(d, ",")
	for i := 0; i < len(distroList); i++ {
		distro := strings.TrimSpace(distroList[i])
		if distro == "" {
			continue
		}
		waitGroup.Add(1)
		go exec(waitGroup, distro)
	}
	waitGroup.Wait()
}

func exec(done *sync.WaitGroup, distro string) {
	defer done.Done()
	if !gowsl.Registered(distro) {
		log.Printf("%v does not exist\n", distro)
		return
	}
	cmd := gowsl.Command(distro, "/bin/bash")
	if err := cmd.Run(); err != nil {
		log.Printf("%v cmd run err: %v", distro, err)
	}
}
