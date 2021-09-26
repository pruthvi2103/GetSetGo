package kurir

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string,error){
	
	if name==""{
		return "",errors.New("Hey there! wont you introdice yourself now :(")
	}

	message := fmt.Sprintf("Hello there %v nice that you greeted me :) \n I am Kurir, named after messenger in Indonesian (or Bahasa)",name)
	return message,nil
}

func Greet(name string) (string,error){
	if name==""{
		return "",errors.New(":( give me your name pls")
	}
	message := fmt.Sprintf(randomGreeting(),name)
	return message,nil
}

func GreetMultiple(names []string) (map[string]string,error){
	messages := make(map[string]string)
	
	for _,name := range names{
		message,err := Greet(name)
		if err !=nil{
			return nil,err
		}
		messages[name]=message
	}
	return messages,nil
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomGreeting() string{
	formats:= []string{
		"Hi %v welcome",
		"Great to see you, %v",
		"General %v",
		"Hello there %v",
	}
	return formats[rand.Intn(len(formats))]
}
