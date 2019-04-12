package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type WorkLoad struct {
	machine  int
	duration int
}

type Job []WorkLoad

func readNumber(lineNumber string) int {
	number, err := strconv.ParseInt(lineNumber, 10, 64)
	check(err)

	return int(number)
}

func parseFile(problem int) *Problem {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/Test data/" + strconv.Itoa(problem) + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Get counts to know number of iterations
	scanner.Scan()
	counts := strings.Fields(scanner.Text())
	numJobs := readNumber(counts[0])
	numMachines := readNumber(counts[1])

	fmt.Println(counts)

	// Read Test data
	jobs := make([]Job, numJobs)

	for i := 0; i < numJobs; i++ {
		scanner.Scan()
		jobInfo := strings.Fields(scanner.Text())

		job := make([]WorkLoad, 0)

		for i := 0; i < len(jobInfo)-1; i += 2 {
			job = append(job, WorkLoad{
				machine:  readNumber(jobInfo[i]),
				duration: readNumber(jobInfo[i+1]),
			})
		}
		jobs[i] = job
	}

	return &Problem{
		numJobs:     numJobs,
		numMachines: numMachines,
		jobs:        jobs,
	}
}
