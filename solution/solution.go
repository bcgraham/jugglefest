package solution

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
	"sort"
)

type Solution struct {
	Circuits   Circuits
	Unassigned Jugglers
	Rejected   Jugglers
	Capacity   int
}

type Circuit struct {
	Name string
	*skills
	Accepted Jugglers
}

type Circuits map[string]*Circuit

type Juggler struct {
	Name               string
	CircuitScores      map[string]int
	CircuitPreferences *list.List
	CurrentCircuit     *list.Element
	Promoted			bool
}

type Jugglers []*Juggler

type skills struct {
	h, e, p int
}

func (j *Juggler) String() string {
	s := j.Name
	s += " "
	for e := j.CircuitPreferences.Front(); e != nil; e = e.Next() {
		s += e.Value.(string)
		s += ":"
		s += strconv.Itoa(j.CircuitScores[e.Value.(string)])
		s += " "
	}
	s = strings.TrimSpace(s)
	return s
}

func (js *Jugglers) Pop()  (j *Juggler) {
	old := *js
	n := len(old)
	j = old[n-1]
	*js = old[0 : n-1]
	return j
}

func (c *Circuit) String() string {
	s := c.Name
	s += " "
	for _, j := range c.Accepted {
		s += j.String()+", "
	}
	s = strings.TrimSuffix(s, ", ")
	return s
}

func (c *Circuit) Len() int				{ return len(c.Accepted) }
func (c *Circuit) Less(i, j int) bool { 
	if c.Accepted[i].CircuitScores[c.Name] == c.Accepted[j].CircuitScores[c.Name] {
		return c.Accepted[i].Promoted && !c.Accepted[j].Promoted
	}
	return c.Accepted[i].CircuitScores[c.Name] > c.Accepted[j].CircuitScores[c.Name] 
}
func (c *Circuit) Swap(i, j int)		{ c.Accepted[i], c.Accepted[j] = c.Accepted[j], c.Accepted[i] }

// MakeCircuit takes an input string and returns a pointer
// to a new Circuit struct. Per the specification: 1. Names
// should never have spaces. 2. Skill and the rating for
// that skill should be separated by a colon. 3. The order
// for attributes will always be Name, Skills (order unspecified).
func MakeCircuit(input string) (c *Circuit, err error) {
	// Sanity checks before we get started. Not empty, right
	// number of elements in attribute string, identifies as
	// a circuit.
	if input == "" {
		return nil, fmt.Errorf("solution: input was empty string")
	}
	attributes := strings.Fields(input)
	if len(attributes) != 5 {
		return nil, fmt.Errorf("solution: not enough/too many fields in input string; expected 5, got: %v", len(attributes))
	}
	if attributes[0] != "C" {
		return nil, fmt.Errorf("solution: input does not identify itself as circuit, expected C, got: %v", attributes[0])
	}

	// Setup
	c = new(Circuit)
	c.Name = attributes[1]
	c.Accepted = make(Jugglers, 0)

	// Skills
	c.skills = new(skills)
	sk, err := makeSkills(attributes[2:5])
	if err != nil {
		return nil, fmt.Errorf("solution: error in MakeCircuit when creating skills, error: %v", err)
	}
	c.h, c.e, c.p = sk.h, sk.e, sk.p

	return c, nil
}

// MakeJuggler takes in a string and a Circuits and returns a pointer
// to a new Juggler struct. Per the specification: 1. Names should
// never have spaces. 2. Skill and the rating for that skill should
// be separated by a colon. 3. Circuit preferences should always
// be a list of comma-separated circuit names. 4. The order for
// attributes will always be Name, Skills (order unspecified),
// Circuit preferences.
func MakeJuggler(input string, cs Circuits) (j *Juggler, err error) {

	// Sanity checks before we get started.
	attributes := strings.Fields(input)
	if len(attributes) != 6 {
		return nil, fmt.Errorf("solution: not enough/too many fields in input string; expected 6, got: %v", len(attributes))
	}
	if attributes[0] != "J" {
		return nil, fmt.Errorf("solution: input does not identify itself as circuit, expected J, got: %v", attributes[0])
	}

	// Setup
	j = &Juggler{}
	j.Name = attributes[1]
	j.CircuitScores = make(map[string]int)
	j.CircuitPreferences = list.New()

	//  Juggler skills; these are not saved with the juggler but
	//  discarded after the calculation of CircuitScores.
	sk, err := makeSkills(attributes[2:5])
	if err != nil {
		return nil, fmt.Errorf("solution: error in MakeJuggler when creating skills, error: %v", err)
	}

	// Circuit scores. Calculate the dot product of each
	// juggler with each of the circuits in its preference
	// list.
	for _, circuitName := range strings.Split(attributes[5], ",") {
		c, ok := cs[circuitName]
		if !ok {
			return nil, fmt.Errorf("solution: juggler prefers a circuit not present in its circuits argument: %v prefers %v", j.Name, circuitName)
		}
		j.CircuitPreferences.PushBack(circuitName)
		j.CircuitScores[circuitName] = sk.h*c.h + sk.e*c.e + sk.p*c.p
	}
	j.CurrentCircuit = j.CircuitPreferences.Front()
	return j, nil
}

func makeSkills(attributes []string) (sk *skills, err error) {
	sk = &skills{}
	var (
		hSet bool
		eSet bool
		pSet bool
	)
	for _, a := range attributes {
		switch strings.Split(a, ":")[0] {
		case "H":
			h, err := strconv.Atoi(strings.Split(a, ":")[1])
			if err != nil {
				return sk, fmt.Errorf("solution: error converting H skill value to int, got: %v", strings.Split(a, ":")[1])
			}
			if hSet {
				return sk, fmt.Errorf("solution: H is being set twice in skillset: %v", attributes)
			}
			hSet = true
			sk.h = h
		case "E":
			e, err := strconv.Atoi(strings.Split(a, ":")[1])
			if err != nil {
				return sk, fmt.Errorf("solution: error converting E skill value to int, got: %v", strings.Split(a, ":")[1])
			}
			if eSet {
				return sk, fmt.Errorf("solution: E is being set twice in skillset: %v", attributes)
			}
			eSet = true
			sk.e = e
		case "P":
			p, err := strconv.Atoi(strings.Split(a, ":")[1])
			if err != nil {
				return sk, fmt.Errorf("solution: error converting P skill value to int, got: %v", strings.Split(a, ":")[1])
			}
			if pSet {
				return sk, fmt.Errorf("solution: P is being set twice in skillset: %v", attributes)
			}
			pSet = true
			sk.p = p
		default:
			return sk, fmt.Errorf("solution: expected skill H, E, or P, got: %v", strings.Split(a, ":")[0])
		}
	}
	if !hSet || !eSet || !pSet {
		return sk, fmt.Errorf("solution: h, e, or p wasn't assigned from input skillset: %v", attributes)
	}
	return sk, nil
}

func (s *Solution) AssignJugglers() {
	for ; len(s.Unassigned)>0; {
		j := s.Unassigned.Pop()
		c := j.CurrentCircuit.Value.(string)
		s.Circuits[c].Accepted = append(s.Circuits[c].Accepted, j)
		if len(s.Circuits[c].Accepted)>s.Capacity {
			sort.Stable(s.Circuits[c])
			r := s.Circuits[c].Accepted.Pop() // r = rejected juggler
			r.CurrentCircuit = r.CurrentCircuit.Next()
			if r.CurrentCircuit == nil {
				r.Promoted = true
				r.CurrentCircuit = r.CircuitPreferences.Front()
				s.Rejected = append(s.Rejected, r)
			} else {
				s.Unassigned = append(s.Unassigned, r)
			}
		}
	}
}

func (cs Circuits) Publish() (ss []string) {
	ss = make([]string, 0)
	for _, c := range cs {
		ss = append(ss, fmt.Sprintf("%v", c))
	}
	return ss
}