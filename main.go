package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	portString := os.Environ()
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Println("Hello World your port ",portString," is exposed")
}