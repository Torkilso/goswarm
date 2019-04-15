package main

func main() {
	problem := parseFile(2)


	operations := particleSwarmOptimization(problem, 1000, 100)

	drawGannt(problem, operations)

	//http.HandleFunc("/", drawChart)
	//http.ListenAndServe(":5000", nil)
}


