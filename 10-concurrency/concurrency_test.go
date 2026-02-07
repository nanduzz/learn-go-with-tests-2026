package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://www.google.com",
		"waat://furhurterwe.geds",
		"http://www.facebook.com",
	}

	want := map[string]bool{
		"http://www.google.com":   true,
		"waat://furhurterwe.geds": false,
		"http://www.facebook.com": true,
	}

	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
