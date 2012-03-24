package using_template_embed

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func ExampleEmbedTemplate() {
	ts := httptest.NewServer(http.HandlerFunc(Hello))
	defer ts.Close()

	res, _ := http.Get(ts.URL + "?name=Felix")
	got, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(got))
	//Output: <h1>Felix</h1>

}
