package main

import "os/exec"
import "log"

func main() {

	// send two pings and send the ouput to STDOUT
	output1, err := exec.Command("ping", "-c", "2", "8.8.8.8").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The results are ----->  %s\n", output1)

	// run the date command and store the output
	output2, err := exec.Command("date").Output()
	if err != nil {
		log.Print(err)
	}
	// another wat to print
	log.Printf("The date is --> %s ", output2)
}
