package finder

import (
	"bytes"
	"errors"
	"log"
	"testing"
)

type mockLog struct {
	buf bytes.Buffer
}

func (m *mockLog) Write(p []byte) (n int, err error) {
	m.buf.Write(p)
	return 0, nil
}

type mockTask struct {
	Source    string
	Error     error
	Count     int
	countCall int
}

func (task *mockTask) Run() (int, error) {
	task.countCall++
	return task.Count, task.Error
}

func (task *mockTask) GetSource() string {
	return task.Source
}

func TestPool(t *testing.T) {
	mockL := mockLog{buf: bytes.Buffer{}}
	stdout := log.New(&mockL, "", 0)

	concurrency := 3
	p := NewPool(concurrency, stdout)

	// add success 5 tasks
	tasks := []Tasker{}
	for i := 0; i < 5; i++ {
		task := &mockTask{
			Source: "http://google.com",
			Count: 1,
		}
		tasks = append(tasks, task)
		p.Put(task)
	}

	// add one error task
	errTask := mockTask{
		Source: "http://google.com",
		Error:  errors.New("fuck"),
		Count:  0,
	}
	p.Put(&errTask)

	if p.currentSize != p.size {
		t.Fatalf("tasks count incorrect after create PollWorkers %d != %d", p.currentSize, p.size)
	}

	p.StopPoolAndWait()

	if p.currentSize != 0 {
		t.Fatalf("tasks count after stop incorrect %d != 0", p.currentSize)
	}

	call := 0
	for _, t := range tasks {
		call += t.(*mockTask).countCall
	}
	if call != len(tasks) {
		t.Fatalf("run tasks: %d, need %d", call, len(tasks))
	}

	if p.Total != 5 {
		t.Fatalf("incorrect total")
	}

}
