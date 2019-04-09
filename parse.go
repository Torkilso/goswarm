package main

import (
	"bufio"
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
	machine int
	duration int

}

type Job = []*WorkLoad

func readNumber(lineNumber string) int {
	number, err := strconv.ParseInt(lineNumber, 10, 64)
	check(err)

	return int(number)
}

func parseFile(problem int) (numMachines int, jobs []Job) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/data/" + strconv.Itoa(problem) + ".txt")
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
	numMachines = readNumber(counts[1])


	// Read data
	jobInfoCount := 0

	for scanner.Scan() {
		jobInfo := strings.Fields(scanner.Text())
		log.Println("jobInfo", jobInfo)

		for i := 0 ; i < len(jobInfo) - 1 ; i++ {
			var job Job

			job = append(job, &WorkLoad{
				machine: readNumber(jobInfo[i]),
				duration: readNumber(jobInfo[i+1]),
			})
			jobs = append(jobs, job)
		}
		jobInfoCount++
		if jobInfoCount >= numJobs {
			break
		}
	}

	return
}
