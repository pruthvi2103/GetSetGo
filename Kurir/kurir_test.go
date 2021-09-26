package kurir

import (
	"regexp"
	"testing"
)

func TestGreeting(t *testing.T){
	name := "Kanobi"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Greet("Kanobi")
	if !want.MatchString(msg) || err!=nil{
		t.Fatalf(`Greet("Kanobi") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestGreetEmpty(t *testing.T){
	msg, err := Greet("")
	if msg!="" || err == nil{
		t.Fatalf(`Greet("") = %q, %v, want "", error`, msg, err)
	}
}