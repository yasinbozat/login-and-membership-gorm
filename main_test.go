package main

import (
	"strconv"
	"testing"
	"time"
)


func TestAddUser(t *testing.T) {

	t.Run("add user", func(t *testing.T) {
		got := AddUser(501, "Yasin", "Bozat", "admin@yasinbozat.com", "123456789", "+90 (543) 987 6543", "Turkey", "Sivas", "99:34:YB:23:BZ:58")
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
		got := SelectUserName(501)
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
		got := CurrentTime().Format(DDMMYYYYhhmmss)
		want := time.Now().Format(DDMMYYYYhhmmss)
		assertCorrectMessage(t, got, want)
	})
}
