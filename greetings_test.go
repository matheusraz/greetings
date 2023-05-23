package greetings

import (
	"regexp"
	"testing"
)

func TesteHelloName(t *testing.T) {
	name := "Xablau"
	must := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !must.MatchString(msg) || err != nil {
		t.Fatalf("%v != %v or something wrong has happened! :(", msg, name)
	}
}
