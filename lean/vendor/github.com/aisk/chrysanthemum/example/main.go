package main

import (
	"github.com/aisk/chrysanthemum"
	"time"
)

func main() {
	c := chrysanthemum.New("I'll be ok").Start()
	time.Sleep(5 * time.Second)
	c.End()
	c = chrysanthemum.New("I'll be error").Start()
	time.Sleep(5 * time.Second)
	c.Failed()
}
