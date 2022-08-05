package main

import (
	"testing"
)

func TestAddUser(t *testing.T) {

	t.Run("add user to db", func(t *testing.T) {
		got := AddUser(6, "Yusuf", DBConnect())
		want := "6:Yusuf"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
