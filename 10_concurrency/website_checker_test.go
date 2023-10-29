package concurrency

import (
	"reflect"
	"testing"
)

func MockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string {
		"https://google.com",
		"https://facebook.com",
		"waat://furhurterwe.geds",
	}

	got := CheckWebsites(MockWebsiteChecker, websites)

	want := map[string]bool {
		"https://google.com": true,
		"https://facebook.com": true,
		"waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	
}