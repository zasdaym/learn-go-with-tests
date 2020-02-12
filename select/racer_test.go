package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("should return error after specified timeout", func(t *testing.T) {
		serverA := makeDelayedServer(15 * time.Millisecond)
		serverB := makeDelayedServer(16 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		timeout := 1 * time.Millisecond
		_, err := ConfigurableRacer(serverA.URL, serverB.URL, timeout)

		if err == nil {
			t.Errorf("expected an error but didn't get that one")
		}
	})

	t.Run("should return faster server", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Errorf("not expected an error but get one %q", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
