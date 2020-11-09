package main

import(
	"log"
	"os"
	"os/signal"
)

func main() {
	s := &Server{}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for sig := range c {
			s.Database.Con.Close()
			log.Println(sig)
			os.Exit(1)
			close(c)
		}
	}()

	s.Database.InitData()
}