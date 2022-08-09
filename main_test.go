package main

import (
	"strconv"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {

	t.Run("add user", func(t *testing.T) {
		got := AddUser(501, "Yasin", "Bozat", "admin@yasinbozat.com", "123456789", "+90 (531) 833 2425", "Turkey", "Sivas", "99:34:YB:23:BZ:58", db())
		want := "501:Yasin"
		assertCorrectMessage(t, got, want)
	})
}

func TestLogin(t *testing.T) {

	t.Run("login", func(t *testing.T) {
		got := strconv.FormatBool(Login("admin@yasinbozat.com", "123456789"))
		want := "true"
		assertCorrectMessage(t, got, want)
	})
}

func TestSelectUser(t *testing.T) {

	t.Run("select user", func(t *testing.T) {
		got := SelectUserName(501, db())
		want := "501:Yasin"
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
