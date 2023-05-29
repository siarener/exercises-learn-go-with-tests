package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"

}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.somedave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":         true,
		"http://blog.somedave5.com": true,
		"waat://furhurterwe.geds":   false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwanted\t%v, \ngot\t%v", want, got)
	}
}
