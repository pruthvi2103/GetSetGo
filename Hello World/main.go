package main

import (
	"fmt"
	"log"
	"pruthvi2103/kurir"
)

func main()  {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)
	names:= []string{"Pruthvi","Shubham","Gandalf","Aragon"}
	message,err:=kurir.Greet("Pruthvi")
	messages,err:=kurir.GreetMultiple(names)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(message)
	for _,singleMessage := range messages{
		fmt.Println(singleMessage)
	} 
}
