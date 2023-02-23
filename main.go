package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(filename string) ([]problem, error) {

	//1 read all the problems from the quiz.csv file
	file, err := os.Open(filename)

	if err != nil {
		return nil, errors.New("A error has ocurred while open the file")
	}
	//2 we will create a new reader
	csvR := csv.NewReader(file)
	//3 it will need to read the file

	contcsv, err := csvR.ReadAll()

	if err == nil {
		return problemPaser(contcsv), nil
	}

	//4
	return nil, errors.New(fmt.Sprintf("A error has ocurred while read the file: %s", err.Error()))
}

func main() {

	//1. input the name of file
	fname := flag.String("f", "quiz.csv", "path of the csv file")

	timer := flag.Int("t", 30, "timer for the quiz")

	flag.Parse()

	problems, err := problemPuller(*fname)

	if err != nil {

		exit(fmt.Sprintf("Something went wrong: %s", err.Error()))
	}

	correctAns := 0

	//2. Set the duration of timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	//3. pull the problems from the file (calling our problem puller function)

	//4. handling the errorserrors.New("A error has ocurred")
	//5. create a variable to count our correct answers
	//6. using the duration of timer, we want to initialize it.
	//7. loop through the problems, print the questions that we`ll accept the answers.

problemloop:

	for i, s := range problems {
		var answer string

		fmt.Printf("Problem %d  %s : ", (i + 1), s.a)

		go func() {
			fmt.Scanln(&answer)

			ansC <- answer
		}()

		select {
		case <-tObj.C:
			break problemloop

		case ansi := <-ansC:

			if ansi == s.b {
				correctAns++
			}
		}
		if i == len(problems)-1 {
			close(ansC)
		}
	}

	fmt.Printf("correct answerd were %d of %d", correctAns, len(problems))
}

func problemPaser(lines [][]string) []problem {

	r := make([]problem, len(lines))

	for i := 0; i < len(lines); i++ {
		r[i] = problem{lines[i][0], lines[i][1]}
	}

	return r
}

type problem struct {
	a string
	b string
}

func exit(text string) {

	fmt.Print(text)
	os.Exit(1)
}
