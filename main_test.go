package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestGetShoplink(t *testing.T) {
	var jsonStr = []byte(`{"productUrl":"https://www.dior.com/bag"`)
	req, err := http.NewRequest("POST", "localhost:8000/api/narrativ/getShoplink", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatalf("could not POST: %v", err)
	}

	rec := httptest.NewRecorder()

	getShoplink(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK, got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response %v", err)
	}

	matched, err := regexp.MatchString(`https://shop-links.co/`, string(bytes.TrimSpace(b)))
	if !matched {
		t.Fatalf("expected a shoplink: got %v", string(b))
	}
}
