package examples

import (
	"fmt"
	"testing"
	"time"
)

type result struct {
	id  int
	res int
}

func worker1(id int, jobs <-chan int, results chan<- result) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(1 * time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- result{id, j * 2}
	}
}

func TestWork(t *testing.T) {

	t.Run("worker pool demo", func(t *testing.T) {
		const numJobs = 5
		jobs := make(chan int, numJobs)
		results := make(chan result, numJobs)

		for w := 1; w <= 3; w++ {
			go worker1(w, jobs, results)
		}

		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
		for a := 1; a <= numJobs; a++ {
			res := <-results
			fmt.Println(res)
		}
	})
}
