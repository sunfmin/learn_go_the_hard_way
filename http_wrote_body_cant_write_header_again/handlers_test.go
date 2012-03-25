package main

import (
	"github.com/sunfmin/integrationtest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSessionValuesInRequests(t *testing.T) {
	ts := httptest.NewServer(Mux())
	defer ts.Close()

	integrationtest.Verbose = true
	s := integrationtest.NewSession()
	s.Get(ts.URL + "/update1")
	r1, _ := s.Get(ts.URL + "/show")
	assertBodyContains(t, r1, []string{"value1"}, []string{"value2"})

	s.Get(ts.URL + "/update2")
	r2, _ := s.Get(ts.URL + "/show")

	assertBodyContains(t, r2, []string{"value1"}, []string{"value2"})

}

func assertBodyContains(t *testing.T, r *http.Response, exists []string, notexists []string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	bodystring := string(b)
	for _, str := range exists {
		if !strings.Contains(bodystring, str) {
			t.Errorf("%v don't contain %v\n\n", bodystring, str)
		}
	}
	for _, str := range notexists {
		if strings.Contains(bodystring, str) {
			t.Errorf("%v contain %v\n\n", bodystring, str)
		}
	}
}
