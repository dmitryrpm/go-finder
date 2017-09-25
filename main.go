package main

import (
	"github.com/dmitryrpm/go-finder/finder"
	"github.com/dmitryrpm/go-finder/config"
	"log"
	"os"
	"bufio"
	"io"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	run(os.Stdin, log.New(os.Stderr, "", 0), cfg)
}

func run(reader io.Reader, stdout *log.Logger, cfg *config.Config){
	p := finder.NewPool(cfg.K, stdout)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if cfg.Type == "url" {
			p.Put(&finder.WebTask{
				Source: scanner.Text(),
			})
		} else {
			p.Put(&finder.FileTask{
				Source: scanner.Text(),
			})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	p.StopPoolAndWait()

	stdout.Printf("Total: %d", p.Total)
}
