package runner

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	log.Println("starting work")
	timeout := 3 * time.Second
	r := New(timeout)
	r.Add(createTask(), createTask(), createTask())

	time.Sleep(3 * time.Second)
	log.Println("sleep end")

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminating duo to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process end.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
