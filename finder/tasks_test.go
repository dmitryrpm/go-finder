// +build integration
package finder

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"os"
	"github.com/pborman/uuid"
	"io/ioutil"
)

func TestSuccessUrl(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Go Go")
	}))
	defer server.Close()

	task := WebTask{
		Source: server.URL,
	}
	count, err := task.Run()
	if err != nil {
		t.Fatalf("Request incorrect: %s", err)
	}

	if count != 3 {
		t.Fatalf("Count incorrect: %d != 3", count)
	}
}

func TestSuccessUrlWithoutGoWord(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "=) trolololo")
	}))
	defer server.Close()

	task := WebTask{
		Source: server.URL,
	}

	if task.GetSource() != task.Source {
		t.Fatalf("Source incorrect: %s != %s", task.GetSource(), task.Source)
	}

	count, err := task.Run()
	if err != nil {
		t.Fatalf("Task.Run incorrect: %s", err)
	}

	if count != 0 {
		t.Fatalf("Count incorrect: %d != 3", count)
	}
}

func TestFailUrl(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Go Go")
	}))
	server.Close()

	task := WebTask{
		Source: server.URL,
	}
	count, err := task.Run()
	res := fmt.Sprintf("Get %s: dial tcp %s: getsockopt: connection refused", server.URL, server.Listener.Addr())
	if err.Error() != res {
		t.Fatalf("Request incorrect: \n%s\n != \n%s", err, res)
	}

	if count != 0 {
		t.Fatalf("Count incorrect: %d != 3", count)
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

	task := FileTask{
		Source: filename,
	}

	if task.GetSource() != task.Source {
		t.Fatalf("Source incorrect: %s != %s", task.GetSource(), task.Source)
	}

	count, err := task.Run()
	if err != nil {
		t.Fatalf("Task.Run incorrect: %s", err)
	}

	if count != 1 {
		t.Fatalf("Count incorrect: %d != 1", count)
	}
}

func TestEmptyFile(t *testing.T) {

	filename := "/tmp/file_test" + uuid.New()

	b := []byte{}
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(filename)

	task := FileTask{
		Source: filename,
	}

	count, err := task.Run()
	if err != nil {
		t.Fatalf("Task.Run incorrect: %s", err)
	}

	if count != 0 {
		t.Fatalf("Count incorrect: %d != 0", count)
	}
}

func TestFileDoesNotExist(t *testing.T) {

	filename := "/tmp/file_test" + uuid.New()

	task := FileTask{
		Source: filename,
	}

	count, err := task.Run()
	res := fmt.Sprintf("open %s: no such file or directory", filename)
	if err.Error() != res {
		t.Fatalf("Task.Run incorrect: %s", err)
	}

	if count != 0 {
		t.Fatalf("Count incorrect: %d != 0", count)
	}
}

