package main

import (
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {

	t.Run("add user", func(t *testing.T) {
		got := AddUser(7, "Bozat", db())
		want := "7:Bozat"
		assertCorrectMessage(t, got, want)
	})
}

func TestSelectUser(t *testing.T) {

	t.Run("select user", func(t *testing.T) {
		got := SelectUser(1, db())
		want := "1:Yasin"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestCurrentTime(t *testing.T) {

	t.Run("current time", func(t *testing.T) {
		got := CurrentTime().Format(TimeFormat)
		want := time.Now().Format(TimeFormat)
		assertCorrectMessage(t, got, want)
	})
}
