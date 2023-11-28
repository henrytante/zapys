package main

import (
	"zapys/src/config"
	"zapys/src/router"
)

func init(){
	config.LoadDotEnv()
}

func main() {
	router.INITSERVER()
}
