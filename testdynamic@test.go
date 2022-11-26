package main

import (
	"github.com/aura-studio/dynamic"
)

/*
-----------------------------------------------------------------------------------

	github.com/aura-studio/testdynamic@test

	require(
		github.com/aura-studio/dynamic@test
	)

-----------------------------------------------------------------------------------

	Test:
		go install github.com/aura-studio/dynamic/dynamic-cli@test
		dynamic-cli build github.com/aura-studio/testdynamic1@test
		dynamic-cli build github.com/aura-studio/testdynamic2@test
		go run ./testdynamic@test.go

-----------------------------------------------------------------------------------
*/
func main() {
	testdynamic1, err := dynamic.GetTunnel("testdynamic1_test")
	if err != nil {
		panic(err)
	}
	testdynamic1.Invoke("", "")

	testdynamic2, err := dynamic.GetTunnel("testdynamic2_test")
	if err != nil {
		panic(err)
	}
	testdynamic2.Invoke("", "")
}
