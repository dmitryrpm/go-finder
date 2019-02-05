// +build integration

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dmitryrpm/go-finder/config"
	"github.com/pborman/uuid"
)

type mockLog struct {
	buf bytes.Buffer
}

func (m *mockLog) Write(p []byte) (n int, err error) {
	m.buf.Write(p)
	return 0, nil
}

func TestSuccessUrl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Go Go")
	}))
	defer server.Close()
	buf := bytes.NewBuffer([]byte(fmt.Sprintf("%s\n%s\n%s\n", server.URL, server.URL, server.URL)))
	mockL := mockLog{buf: bytes.Buffer{}}
	stdout := log.New(&mockL, "", 0)
	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	run(buf, stdout, cfg)
	res := fmt.Sprintf("Count for %s: 3\nCount for %s: 3\nCount for %s: 3\nTotal: 9\n", server.URL, server.URL, server.URL)
	if mockL.buf.String() != res {
		t.Fatalf("\n'%s'\n!=\n'%s'", res, mockL.buf.String())
	}
}

func TestSuccessFile(t *testing.T) {
	filename := "/tmp/file_test" + uuid.New()

	b := []byte("hello\nGo\n")
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	buf := bytes.NewBuffer([]byte(fmt.Sprintf("%s\n%s\n%s\n", filename, filename, filename)))
	mockL := mockLog{buf: bytes.Buffer{}}
	stdout := log.New(&mockL, "", 0)
	if err != nil {
		log.Fatal(err)
	}
	run(buf, stdout, &config.Config{K: 5, Type: "file"})
	res := fmt.Sprintf("Count for %s: 1\nCount for %s: 1\nCount for %s: 1\nTotal: 3\n", filename, filename, filename)
	if mockL.buf.String() != res {
		t.Fatalf("\n'%s'\n!=\n'%s'", res, mockL.buf.String())
	}
}
