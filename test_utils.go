package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestResponse(method string, url string) (string, int, error) {
	req, _ := http.NewRequest(method, url, nil)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", 0, err
	}

	return strings.Trim(string(b), "\n"), res.StatusCode, nil
}

func testServer(srv *httptest.Server, t *testing.T) {
	b, code, _ := getTestResponse("GET", srv.URL+"/create?url=invalidurl")
	t.Log("ak:invalid url")
	t.Log(b)
	assert.Equal(t, http.StatusInternalServerError, code)

	key, code, _ := getTestResponse("GET", srv.URL+"/create?url=https://www.leetcode.com")
	assert.Equal(t, http.StatusCreated, code)
	assert.Len(t, key, keyLength)
	t.Log("ak:key created")
	t.Log(key)

	b, code, _ = getTestResponse("GET", srv.URL+"/get?key=invalidkey")
	assert.Equal(t, http.StatusNotFound, code)	
	t.Log("ak:invalid key")
	t.Log(b)

	t.Log("ak:url retrieval by key")
	

	
	url, code, _ := getTestResponse("GET", srv.URL+"/get?key="+key)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "https://www.leetcode.com", url)
	t.Log(url)
	t.Log("ak:end of testServer")
}
