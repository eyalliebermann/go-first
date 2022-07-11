package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Jakob"
	//	want := regexp.MustCompile("\b" + name + "\b")//FAILS " vs '
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello(name)
	match := want.MatchString(msg)
	if !match || err != nil {
		t.Fatalf(`Hello("") = %q, match= %v, error= %v, want "", error`, msg, match, err)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf("Hello should return error when input is empty")
	}
}
