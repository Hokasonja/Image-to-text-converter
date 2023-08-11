package main

import "testing"

func TestHello(t *testing.T) {

	want := "Hello HokaSonja"

	got := hello()

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}

func TestPlus(t *testing.T) {

	want := 5

	got := plus(2, 3)

	if want != got {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}
