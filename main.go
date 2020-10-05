package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/tarm/serial"
)

func main() {
	flag.Parse()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	c := &serial.Config{Name: flag.Arg(0), Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			fmt.Fprintln(s, sc.Text())
		}
	}()
	go func() {
		sc := bufio.NewScanner(s)
		for sc.Scan() {
			log.Println(sc.Text())
		}
		close(sig)
	}()
	<-sig
	fmt.Println()
	fmt.Println("terminated")
	s.Close()
}
