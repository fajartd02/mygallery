package main

import (
	"testing"
)

const (
	healthURL = "http://20.70.234.83/health"
)

func hello() string {
	return "Hello Golang"
}

func TestHello(t *testing.T) {

	want := "Hello Golang"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}

func TestHealth(t *testing.T) {

	//response, err := http.Get(healthURL)
	//if err != nil {
	//	log.Fatal(err.Error())
	//	fmt.Print(err.Error())
	//	os.Exit(1)
	//}
	//
	//responseData, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal("Failed to read response body")
	//}
	//
	//if string(responseData) != "ok" {
	//	log.Fatal("Response health is not equal")
	//}
}
