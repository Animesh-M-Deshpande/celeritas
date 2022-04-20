package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "no error rendering non-existent template, while one is expected"},
	{"jet_page", "jet", "home", false, "rendering go template"},
	{"jet_page_no_template", "jet", "no-file", true, "no error rendering non-existent template, while one is expected"},
}

func TestRender_Page(t *testing.T) {

	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/some-url", nil)

		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = "go"
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)

		if e.errorExpected {

			if err == nil {
				t.Errorf("%s:%s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s:%s:%s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/url", nil)

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err := testRenderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error("Error rendering page")
	}

	testRenderer.Renderer = ""
	err = testRenderer.Page(w, r, "home", nil, nil)

	if err == nil {
		t.Error("Error should be thrown if we do not specifiy a valid engine")
	}

}

func TestRender_JetPage(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/url", nil)

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"

	err := testRenderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Error("Error rendering page")
	}

}
