package demohttptest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(hello))
	defer ts.Close()

	res, _ := http.Get(ts.URL + "?name=Felix")
	got, _ := ioutil.ReadAll(res.Body)

	fmt.Println(ts.URL)

	if string(got) != "Hello, Felix" {
		t.Errorf("got %q", string(got))
	}
}
