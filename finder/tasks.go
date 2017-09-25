package finder

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Tasker interface {
	Run() (int, error)
	GetSource() string
}

/*
URL-task implementation
*/

type WebTask struct {
	Source string
}

func (task *WebTask) Run() (count int, err error) {
	resp, err := http.Get(task.Source)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return count, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return count, err
	}
	return strings.Count(string(body), "Go"), err
}

func (task *WebTask) GetSource() string {
	return task.Source
}

/*
FILE-task implementation
*/

type FileTask struct {
	Source string
}

func (task *FileTask) Run() (count int, err error) {
	data, err := ioutil.ReadFile(task.Source)
	if err != nil {
		return count, err
	}
	return strings.Count(string(data), "Go"), nil
}

func (task *FileTask) GetSource() string {
	return task.Source
}
