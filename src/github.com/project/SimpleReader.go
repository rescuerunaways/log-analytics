package main

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
	"sync"
	"github.com/luci/luci-go/common/runtime/goroutine"

	"time"
)

var rgxp = regexp.MustCompile("event")

func openAndRead(done chan bool) {
	file, _ := os.Open("result.txt")
	jobs := make(chan string)
	results := make(chan string)

	wg := new(sync.WaitGroup)


	// get 3 workers,add them to the waiting groups
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go processJobs(jobs, results, wg)
	}


	// Go over a file line by line and queue up a ton of work
	go func() {
		//fmt.Println( " scanning goroutine:", goroutine.CurID())

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if rgxp.MatchString(scanner.Text()) {
				jobs <- scanner.Text()

			}

		}
		close(jobs)

	}()

	go func() {

		wg.Wait()
		close(results)
	}()

	counts := 0

	for range results {
		counts += 1
		//fmt.Println("result channel:", v)
	}

	fmt.Println("processed and matched events:", counts)
	done <- true
}

func processJobs(jobs <-chan string, results chan <- string, wg *sync.WaitGroup) {
	fmt.Println(" starting goroutine:", goroutine.CurID())

	defer wg.Done()

	for j := range jobs {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("processed job:", j, " by goroutine:", goroutine.CurID())

		results <- j
	}
}













