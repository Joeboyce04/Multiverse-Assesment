package main

import (
	"errors"
	"testing"
)

func TestFilterClientsByRevenue_ValidInput(t *testing.T) {
	clients := []Client{
		{"Client A", 10000, true},
		{"Client B", 200, false},
		{"Client C", 15000, true},
	}

	minRevenue := 1000

	want := []Client{
		{"Client A", 10000, true},
		{"Client C", 15000, true},
	}
	
	got, _ := FilterClientsByRevenue(clients, minRevenue)
	
	if len(got) != len(want) {
		t.Error("got",len(got),"does not equal want",len(want))
	}

	for i, client := range got {
		if client != want[i] {
			t.Error("got",client," want ",want[i])
		}
	}
}


func TestFilterClientsByRevenue_NegativeRevenue(t *testing.T) {
	
	clients := []Client{
		{"Client A", 10000, true},
		{"Client B", 5000, false},
	}
	minRevenue := -5000
	wantErr := errors.New("revenue cannot be negative")
	
	_, gotErr := FilterClientsByRevenue(clients, minRevenue)
	
	if gotErr == nil || gotErr.Error() != wantErr.Error() {
		t.Error("got error",gotErr,"want error",wantErr,)
	}
}

func TestFilterClientsBySignedStatus(t *testing.T) {
	
	clients := []Client{
		{"Client A", 10000, true},
		{"Client B", 5000, false},
		{"Client C", 15000, true},
	}
	signed := true
	want := []Client{
		{"Client A", 10000, true},
		{"Client C", 15000, true},
	}
	
	got := FilterClientsBySignedStatus(clients, signed)
	
	if len(got) != len(want) {
		t.Error("got",len(got)," clients want", len(want),)
	}
	for i, client := range got {
		if client != want[i] {
			t.Error("got",client,"want", want[i],)
		}
	}
}