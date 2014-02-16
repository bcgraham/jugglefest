package main

import (
	"github.com/bcgraham/jugglefest/solution"
	"fmt"
	"os"
	"strings"
	"flag"
	"bufio"
	"bytes"
	//"sort"	
)

var input = flag.String("input", "", "path to the input file")
var output = flag.String("output", "", "path to the output file")

func main() {
	flag.Parse()
	cs, js, err := readLines(*input)
	if err != nil {
		panic(err)
	}

	r := make(solution.Jugglers, 0)
	capacity := len(js) / len(cs)
	s := &solution.Solution{cs, js, r, capacity}

	s.AssignJugglers()

	// Run through rejected jugglers, giving them preference in ties. 
	s.Unassigned = s.Rejected
	s.Rejected = make(solution.Jugglers, 0)
	s.AssignJugglers()

	fmt.Printf("%v unmatched jugglers.", len(s.Rejected))

	// Assign remaining jugglers at random. 
	for _, c := range s.Circuits {
		for len(c.Accepted) < 6 {
			c.Accepted = append(c.Accepted, s.Rejected.Pop())
		}
	}
	
	results := s.Circuits.Publish()	
	//sort.Strings(results)
	
	writeFile, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer writeFile.Close()

	w := bufio.NewWriter(writeFile)
	var line bytes.Buffer
	for _, r := range results {
		line.Reset()
		line.WriteString(r)
		fmt.Fprintln(w, line.String())
	}
	w.Flush()
}


func readLines(path string) (cs solution.Circuits, js solution.Jugglers, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	cs = make(solution.Circuits)
	js = make(solution.Jugglers, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line,"C") {
			c, err := solution.MakeCircuit(line)
			if err != nil {
				panic(err)
			}
			cs[c.Name] = c
		}
		if strings.HasPrefix(line,"J") {
			j, err := solution.MakeJuggler(line, cs)
			if err != nil {
				panic(err)
			}
			js = append(js, j)
		}
	}
	return cs, js, scanner.Err()
}
