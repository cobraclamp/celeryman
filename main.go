package main

import (
	"fmt"

	"github.com/cobraclamp/celeryman/bot"
	"github.com/cobraclamp/celeryman/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})

}
