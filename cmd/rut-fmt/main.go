package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mnavarrocarter/go-rut"
	"log"
	"os"
)

func main() {
	skipErr := flag.Bool("skip-error", false, "When present, invalid ruts are skipped")

	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		str := s.Text()
		r, err := rut.Parse(str)
		if err != nil && !*skipErr {
			log.Fatalf("error parsing %s: %s", str, err.Error())
		}
		fmt.Println(r.String())
	}
}
