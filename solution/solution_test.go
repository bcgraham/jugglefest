package solution

import (
	"reflect"
	"testing"
	"fmt"
	"sort"
)

func mockAttributes() []string {
	attributes := make([]string, 6)
	attributes[0] = ""
	attributes[1] = ""
	attributes[2] = "H:1"
	attributes[3] = "E:1"
	attributes[4] = "P:1"
	attributes[5] = ""
	return attributes
}

func mockCircuitsCase1() Circuits {
	cs := make(Circuits)
	ss := make([]string, 0)
	ss = append(ss, "C C0 H:1 E:1 P:1")
	ss = append(ss, "C C1 H:2 E:2 P:2")
	for _, s := range ss {
		c, err := MakeCircuit(s)
		if err != nil {
			panic(err)
		}
		cs[c.Name] = c
	}
	return cs
}

func mockJugglersCase1() Jugglers {
	cs := mockCircuitsCase1()
	jugglers := make(Jugglers, 0)
	ss := make([]string, 0)
	ss = append(ss, "J J0 H:1 E:1 P:1 C0,C1")
	ss = append(ss, "J J1 H:2 E:2 P:2 C0,C1")
	ss = append(ss, "J J2 H:3 E:3 P:3 C0,C1")
	ss = append(ss, "J J3 H:4 E:4 P:4 C1,C0")
	for _, s := range ss {
		j, err := MakeJuggler(s, cs)
		if err != nil {
			panic(err)
		}
		jugglers = append(jugglers, j)
	} 
	return jugglers
}

func mockCircuitsCase2() Circuits {
	cs := make(Circuits)
	ss := make([]string, 0)
	ss = append(ss, "C C1 H:1 E:1 P:1")	
	ss = append(ss, "C C2 H:2 E:2 P:2")	
	ss = append(ss, "C C3 H:3 E:3 P:3")	
	ss = append(ss, "C C4 H:4 E:4 P:4")	
	for _, s := range ss {
		c, err := MakeCircuit(s)	
		if err != nil {
			panic(err)
		}
		cs[c.Name] = c
	}
	return cs
}

func mockJugglersCase2() Jugglers {
	cs := mockCircuitsCase2()
	jugglers := make(Jugglers, 0)
	ss := make([]string, 0)
	ss = append(ss, "J J1 H:1 E:1 P:1 C1,C4")
	ss = append(ss, "J J2 H:2 E:2 P:2 C2,C4")
	ss = append(ss, "J J3 H:3 E:3 P:3 C3,C1")
	ss = append(ss, "J J4 H:4 E:4 P:4 C3,C2")
	ss = append(ss, "J J5 H:5 E:5 P:5 C3,C4")
	ss = append(ss, "J J6 H:6 E:6 P:6 C4,C1")
	ss = append(ss, "J J7 H:7 E:7 P:7 C4,C2")
	ss = append(ss, "J J8 H:8 E:8 P:8 C4,C3") 
	for _, s := range ss {
		j, err := MakeJuggler(s, cs)
		if err != nil {
			panic(err)
		}
		jugglers = append(jugglers, j)
	} 
	return jugglers
}

func mockCircuitsCase3() Circuits {
	cs := make(Circuits)
	ss := make([]string, 0)
	ss = append(ss, "C C1 H:1 E:1 P:1")	
	ss = append(ss, "C C2 H:1 E:1 P:1")	
	ss = append(ss, "C C3 H:1 E:1 P:1")	
	ss = append(ss, "C C4 H:1 E:1 P:1")	
	for _, s := range ss {
		c, err := MakeCircuit(s)	
		if err != nil {
			panic(err)
		}
		cs[c.Name] = c
	}
	return cs
}

func mockJugglersCase3() Jugglers {
	cs := mockCircuitsCase3()
	jugglers := make(Jugglers, 0)
	ss := make([]string, 0)
	ss = append(ss, "J J1 H:1 E:1 P:1 C1,C4")
	ss = append(ss, "J J2 H:1 E:1 P:1 C2,C4")
	ss = append(ss, "J J3 H:2 E:2 P:2 C3,C1")
	ss = append(ss, "J J4 H:2 E:2 P:2 C3,C2")
	ss = append(ss, "J J5 H:3 E:3 P:3 C3,C4")
	ss = append(ss, "J J6 H:3 E:3 P:3 C4,C1")
	ss = append(ss, "J J7 H:4 E:4 P:4 C4,C2")
	ss = append(ss, "J J8 H:4 E:4 P:4 C4,C3") 
	for _, s := range ss {
		j, err := MakeJuggler(s, cs)
		if err != nil {
			panic(err)
		}
		jugglers = append(jugglers, j)
	} 
	return jugglers
}

func TestMakeCircuitReturnsTypePointerCircuits(t *testing.T) {
	c, err := MakeCircuit("C C1999 H:9 E:5 P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if reflect.TypeOf(c).String() != "*solution.Circuit" {
		t.Errorf("Expected type *solution.Circuit, got %v", reflect.TypeOf(c).String())
	}
}

func TestMakeCircuitNameSingleSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C C1999 H:9 E:5 P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.Name != "C1999" {
		t.Errorf("Expected c.Name = C1999, got: %v", c.Name)
	}
}

func TestMakeCircuitNameTabSeparators(t *testing.T) {
	c, err := MakeCircuit("C	C1999	H:9	E:5	P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.Name != "C1999" {
		t.Errorf("Expected c.Name = C1999, got: %v", c.Name)
	}
}

func TestMakeCircuitHSingleSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C C1999 H:9 E:5 P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.h != 9 {
		t.Errorf("Expected c.h = 9, got: %v", c.h)
	}
}
func TestMakeCircuitHMultiSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C  C1999  H:9  E:5  P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.h != 9 {
		t.Errorf("Expected c.h = 9, got: %v", c.h)
	}
}

func TestMakeCircuitHVariableOrder(t *testing.T) {
	c, err := MakeCircuit("C  C1999  E:9  P:5  H:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.h != 9 {
		t.Errorf("Expected c.h = 9, got: %v", c.h)
	}
}

func TestMakeCircuitESingleSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C C1999 H:9 E:5 P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.e != 5 {
		t.Errorf("Expected c.e = 5, got: %v", c.h)
	}
}

func TestMakeCircuitEMultiSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C  C1999  H:9  E:5  P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.e != 5 {
		t.Errorf("Expected c.e = 5, got: %v", c.h)
	}
}

func TestMakeCircuitEVariableOrder(t *testing.T) {
	c, err := MakeCircuit("C  C1999  H:9  P:9  E:5")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.e != 5 {
		t.Errorf("Expected c.e = 5, got: %v", c.h)
	}
}

func TestMakeCircuitPSingleSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C C1999 H:9 E:5 P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.p != 9 {
		t.Errorf("Expected c.p = 9, got: %v", c.h)
	}
}

func TestMakeCircuitPMultiSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C C1999  H:9  E:5  P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.p != 9 {
		t.Errorf("Expected c.p = 9, got: %v", c.h)
	}
}

func TestMakeCircuitPVariableOrder(t *testing.T) {
	c, err := MakeCircuit("C C1999  P:9  E:5  H:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.p != 9 {
		t.Errorf("Expected c.p = 9, got: %v", c.h)
	}
}

func TestMakeCircuitNameMultiSpaceSeparators(t *testing.T) {
	c, err := MakeCircuit("C  C1999  H:9  E:5  P:9")
	if err != nil {
		t.Errorf("MakeCircuit returned error, got %v", err)
	}
	if c.Name != "C1999" {
		t.Errorf("Expected c.Name = C1999, got: %v", c.Name)
	}
}

func TestMakeCircuitWithBadName(t *testing.T) {
	c, err := MakeCircuit("C  H:9  E:5  P:9")
	if err == nil {
		t.Errorf("MakeCircuit did not return an error when no circuit name was present in the input string, got: %v", c)
	}
}

func TestMakeCircuitWithNoH(t *testing.T) {
	c, err := MakeCircuit("C C1999 E:5  P:9")
	if err == nil {
		t.Errorf("MakeCircuit did not return an error when no circuit name was present in the input string, got: %v", c)
	}
}

func TestMakeCircuitWithNoE(t *testing.T) {
	c, err := MakeCircuit("C C1999 H:5  P:9")
	if err == nil {
		t.Errorf("MakeCircuit did not return an error when no circuit name was present in the input string, got: %v", c)
	}
}

func TestMakeCircuitWithNoP(t *testing.T) {
	c, err := MakeCircuit("C C1999 E:5  H:9")
	if err == nil {
		t.Errorf("MakeCircuit did not return an error when no circuit name was present in the input string, got: %v", c)
	}
}

func TestMakeCircuitEmptyString(t *testing.T) {
	c, err := MakeCircuit("")
	if err == nil {
		t.Errorf("MakeCircuit did not return an error when an empty string was passed to it, got: %v", c)
	}
}

func TestMakeCircuitUnidentified(t *testing.T) {
	c, err := MakeCircuit("J C1999 H:9 E:5 P:9")
	if err == nil {
		t.Errorf("MakeCircuit created a circuit not identified as a circuit, made: %v", c)
	}
}

func TestMakeCircuitUnidentifiedAttribute(t *testing.T) {
	c, err := MakeCircuit("C C1999 H:9 E:5 F:9")
	if err == nil {
		t.Errorf("MakeCircuit created a circuit with an attribute not identified in the specification, made: %+v", c)
	}
}

func TestMakeCircuitRepeatedAlternatingWhitespace(t *testing.T) {
	_, err := MakeCircuit("C	 	 	C101	 	   			 P:909	 	 	 	 	E:100002	  	 	   H:1 	 		  	 	")
	if err != nil {
		t.Errorf("MakeCircuit did not parse valid input, returned error: %v", err)
	}
}

func TestMakeJugglerReturnsTypePointerJuggler(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J0 H:7 E:6 P:0 C0,C1", cs)
	if err != nil {
		t.Errorf("MakeJuggler did not parse valid input, returned error: %v", err)
	}
	if reflect.TypeOf(j).String() != "*solution.Juggler" {
		t.Errorf("Expected type *solution.Juggler, got %v", reflect.TypeOf(j).String())
	}
}

func TestMakeJugglerWithNonexistentCircuitInPrefs(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J0 H:1 E:1 P:1 C0,C3", cs)
	if err == nil {
		t.Errorf("MakeJuggler did not error out when passed a preference for a circuit not present in its Circuits argument, created: %+v", j)
	}
}

func TestMakeJugglerEmptyString(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("", cs)
	if err == nil {
		t.Errorf("MakeJuggler did not error out when passed an empty string, created: %+v", j)
	}
}

func TestMakeJugglerUnidentified(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("C J0 H:1 E:1 P:1 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler not identified as a juggler, made: %+v", j)
	}
}

func TestMakeJugglerWithoutName(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J H:1 E:1 P:1 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed input string without name in it, made: %+v", j)
	}
}

func TestMakeJugglerWithoutPreferences(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 H:1 E:1 P:1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed input string without comma-separated preferences in it, made: %+v", j)
	}
}

func TestMakeJugglerWithoutH(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 E:1 P:1 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed input string without H in it, made: %+v", j)
	}
}
func TestMakeJugglerWithoutE(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 H:1 P:1 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed input string without E in it, made: %+v", j)
	}
}
func TestMakeJugglerWithoutP(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 E:1 H:1 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed input string without P in it, made: %+v", j)
	}
}

func TestMakeJugglerWithMultipleH(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 H:1 E:1 P:1 H:2 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed multiple H attributes in it, made: %+v", j)
	}
}

func TestMakeJugglerWithMultipleE(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 H:1 E:1 P:1 E:2 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed multiple E attributes in it, made: %+v", j)
	}
}

func TestMakeJugglerWithMultipleP(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 H:1 E:1 P:1 P:2 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed multiple P attributes in it, made: %+v", j)
	}
}

func TestMakeJugglerWithUnidentifiedAttribute(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 H:1 E:1 P:1 F:2 C0,C1", cs)
	if err == nil {
		t.Errorf("MakeJuggler created a juggler when passed an attribute unidentified by the specification, made: %+v", j)
	}
}

func TestMakeJugglerCurrentCircuitSet(t *testing.T) {
	cs := mockCircuitsCase1()
	j, err := MakeJuggler("J J1 H:1 E:1 P:1 C0,C1", cs)
	if err != nil {
		t.Errorf("Error when creating juggler from valid input, returned error: %v", err)
	}
	e := j.CircuitPreferences.Front()
	if e.Value != j.CurrentCircuit.Value {
		t.Errorf("Juggler's first preferred circuit not the front of its CircuitPreferences list; front = %v, while preferred = %v", e.Value, j.CurrentCircuit.Value)
	}
}

func TestMakeSkillsReturnsTypeSkills(t *testing.T) {
	attributes := mockAttributes()
	sk, err := makeSkills(attributes[2:5])
	if err != nil {
		t.Errorf("makeSkills did not parse valid input, returned error: %v", err)
	}
	if reflect.TypeOf(sk).String() != "*solution.skills" {
		t.Errorf("Expected type *solution.skills, got %v", reflect.TypeOf(sk).String())
	}
}

func TestMakeSkillsNonNumericSkillValueH(t *testing.T) {
	attributes := mockAttributes()
	attributes[2] = "H:One"
	sk, err := makeSkills(attributes[2:5])
	if err == nil {
		t.Errorf("makeSkills did not throw an error when it was passed a non-numeric value in H, got: %+v", sk)
	}
}
func TestMakeSkillsNonNumericSkillValueE(t *testing.T) {
	attributes := mockAttributes()
	attributes[3] = "E:One"
	sk, err := makeSkills(attributes[2:5])
	if err == nil {
		t.Errorf("makeSkills did not throw an error when it was passed a non-numeric value in E, got: %+v", sk)
	}
}
func TestMakeSkillsNonNumericSkillValueP(t *testing.T) {
	attributes := mockAttributes()
	attributes[4] = "P:One"
	sk, err := makeSkills(attributes[2:5])
	if err == nil {
		t.Errorf("makeSkills did not throw an error when it was passed a non-numeric value in P, got: %+v", sk)
	}
}

func TestPopRemoves(t *testing.T) {
	js := mockJugglersCase1()
	length := len(js)
	_ = js.Pop()
	if length-len(js)!=1 {
		t.Errorf("Pop() does not remove elements; length was %v and became %v", length, len(js))
	}
}

func TestPopReturns(t *testing.T) {
	js := mockJugglersCase1()
	jstring := js[len(js)-1].String()
	j := js.Pop()
	if jstring != fmt.Sprintf(j.String()) {
		t.Errorf("Pop() does not return element; expected %v, got: %v", jstring, j.String())
	}
}

func TestJugglerStringFirstExample(t *testing.T) {
	js := mockJugglersCase1() 
	if fmt.Sprintf(js[0].String()) != "J0 C0:3 C1:6" {
		t.Errorf("Juggler default printing does not meet specified format, expected \"J0 C0:3 C1:6\", got: %v", js[0])
	}
}
func TestJugglerStringSecondExample(t *testing.T) {
	js := mockJugglersCase1() 
	if fmt.Sprintf(js[1].String()) != "J1 C0:6 C1:12" {
		t.Errorf("Juggler default printing does not meet specified format, expected \"J1 C0:6 C1:12\", got: %v", js[1])
	}
}


func TestCircuitStringFirstExample(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	if fmt.Sprintf(cs["C0"].String()) != "C0 J3 C1:24 C0:12, J2 C0:9 C1:18" {
		t.Errorf("Circuit default printing does not meet specified format, expected \"C0 J3 C1:24 C0:12, J2 C0:9 C1:18\", got: %v", cs["C0"])
	}
}

func TestCircuitStringSecondExample(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	
	cs["C1"].Accepted = append(cs["C1"].Accepted, js.Pop())
	cs["C1"].Accepted = append(cs["C1"].Accepted, js.Pop())
	
	if fmt.Sprintf(cs["C1"].String()) != "C1 J3 C1:24 C0:12, J2 C0:9 C1:18" {
		t.Errorf("Circuit default printing does not meet specified format, expected \"C1 J3 C1:24 C0:12, J2 C0:9 C1:18\", got: %v", cs["C1"])
	}
}

func TestCircuitEmptyLen(t *testing.T) {
	cs := mockCircuitsCase1()
	
	if cs["C0"].Len() != 0 {
		t.Errorf("Circuit with empty Accepted returns non-zero Len(), returned: %v", cs["C0"].Len())
	}
}

func TestCircuitSomeLen(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()

	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	if cs["C0"].Len() != 2 {
		t.Errorf("Circuit with Accepted of length 2 returns Len()!=2, returned: %v", cs["C0"].Len())
	}
}
	
func TestCircuitSwap(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	J3 := cs["C0"].Accepted[0].Name
	J2 := cs["C0"].Accepted[1].Name
	
	cs["C0"].Swap(0, 1)
	
	if (J3 != cs["C0"].Accepted[1].Name || J2 != cs["C0"].Accepted[0].Name) {
		t.Errorf("Swap did not swap; expected cs.Accepted[1].Name == \"J3\" and cs.Accepted[0].Name == \"J2\" but actually cs.Accepted[1].Name == \"%v\" and cs.Accepted[0].Name == \"%v\"", cs["C0"].Accepted[1].Name, cs["C0"].Accepted[0].Name)
	}
}

func TestCircitLessTrue(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()

	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	// For C0, J3(index 0) > J2 (index 1); J3 should 
	// be sorted before J2 (i.e. Pop() should encounter
	// J2 before J3), so "Less" in this case is "More".
	if !cs["C0"].Less(0, 1) {
		t.Errorf("Less reported J2 should be sorted before J3; J3 = %v and J2 = %v", cs["C0"].Accepted[0].CircuitScores["C0"], cs["C0"].Accepted[1].CircuitScores["C0"])
	}
}

func TestCircitLessFalse(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	// For C0, J3(index 0) > J2 (index 1); J3 should 
	// be sorted before J2 (i.e. Pop() should encounter
	// J2 before J3), so "Less" in this case is "More".
	if cs["C0"].Less(1, 0) {
		t.Errorf("Less reported J2 should be sorted before J3; J3 = %v and J2 = %v", cs["C0"].Accepted[0].CircuitScores["C0"], cs["C0"].Accepted[1].CircuitScores["C0"])
	}
}

func TestCircitLessEqualPromotedI(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()

	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	cs["C0"].Accepted[0].CircuitScores["C0"] = 0
	cs["C0"].Accepted[1].CircuitScores["C0"] = 0
	cs["C0"].Accepted[0].Promoted = true
	cs["C0"].Accepted[1].Promoted = false
	
	// For C0, J3(index 0) > J2 (index 1); J3 should 
	// be sorted before J2 (i.e. Pop() should encounter
	// J2 before J3), so "Less" in this case is "More".
	if !cs["C0"].Less(0, 1) {
		t.Errorf("Promoted i, unpromoted j should return i \"less\" than j as true, but returns: %v", cs["C0"].Less(0, 1))
	}
}

func TestCircitLessEqualPromotedJ(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	cs["C0"].Accepted[0].CircuitScores["C0"] = 0
	cs["C0"].Accepted[1].CircuitScores["C0"] = 0
	cs["C0"].Accepted[0].Promoted = false
	cs["C0"].Accepted[1].Promoted = true
	
	// For C0, J3(index 0) > J2 (index 1); J3 should 
	// be sorted before J2 (i.e. Pop() should encounter
	// J2 before J3), so "Less" in this case is "More".
	if cs["C0"].Less(0, 1) {
		t.Errorf("Unpromoted i, promoted j should return i \"less\" than j as false, but returns: %v", cs["C0"].Less(0, 1))
	}
}

func TestCircitLessEqualBothPromoted(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()

	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	cs["C0"].Accepted[0].CircuitScores["C0"] = 0
	cs["C0"].Accepted[1].CircuitScores["C0"] = 0
	cs["C0"].Accepted[0].Promoted = true
	cs["C0"].Accepted[1].Promoted = true
	
	// For C0, J3(index 0) > J2 (index 1); J3 should 
	// be sorted before J2 (i.e. Pop() should encounter
	// J2 before J3), so "Less" in this case is "More".
	if cs["C0"].Less(0, 1) {
		t.Errorf("Promoted i, promoted j should return i \"less\" than j as false, but returns: %v", cs["C0"].Less(0, 1))
	}
}

func TestCircitLessEqualNeitherPromoted(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	
	cs["C0"].Accepted[0].CircuitScores["C0"] = 0
	cs["C0"].Accepted[1].CircuitScores["C0"] = 0
	cs["C0"].Accepted[0].Promoted = false
	cs["C0"].Accepted[1].Promoted = false
	
	// For C0, J3(index 0) > J2 (index 1); J3 should 
	// be sorted before J2 (i.e. Pop() should encounter
	// J2 before J3), so "Less" in this case is "More".
	if cs["C0"].Less(0, 1) {
		t.Errorf("Unpromoted i, unpromoted j should return i \"less\" than j as false, but returns: %v", cs["C0"].Less(0, 1))
	}
}

// Case1: Every circuit appears in every juggler's list 
// of preferences. No random assignment. 
func TestAssignJugglersCase1Length(t *testing.T) {
	
	// In the mock examples, J3>J2>J1>J0. As there are 
	// four jugglers and two circuits, each circuit will
	// accept a total of two jugglers. J3 wants C1 and 
	// will be assigned to this circuit. J2 & J1 both
	// want C0 and will be assigned to this circuit. J0
	// wants C0, but C0 has been filled, so will take his
	// second pick of C1. 
	
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	r := make(Jugglers, 0)
	capacity := len(js) / len(cs)
	s := &Solution{cs, js, r, capacity}
	
	s.AssignJugglers()
	
	if len(s.Rejected)>0 {
		t.Errorf("Case 1 Solution's list of rejected jugglers should be empty, but contains %v jugglers.", len(s.Rejected))
	}
	
}

func TestCircuitsPublish(t *testing.T) {
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	
	// To simplify the test case, I manually sort the
	// jugglers as if they had been assigned appropriately.
	cs["C1"].Accepted = append(cs["C1"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C0"].Accepted = append(cs["C0"].Accepted, js.Pop())
	cs["C1"].Accepted = append(cs["C1"].Accepted, js.Pop())
	
	results := cs.Publish()
	
	sort.Strings(results)
	if results[0] != "C0 J2 C0:9 C1:18, J1 C0:6 C1:12" || results[1] != "C1 J3 C1:24 C0:12, J0 C0:3 C1:6" {
		t.Errorf("Publish() published wrong results; expected results[0] == \"C0 J2 C0:9 C1:18, J1 C0:6 C1:12\" but got %v and expected results[1] == \"C1 J3 C1:24 C0:12, J0 C0:3 C1:6\" but got %v", results[0], results[1])
	}
}

// Case1: Every circuit appears in every juggler's list 
// of preferences. No random assignment. 
 func TestAssignJugglersCase1Content(t *testing.T) {
	
	// In the mock examples, J3>J2>J1>J0. As there are 
	// four jugglers and two circuits, each circuit will
	// accept a total of two jugglers. J3 wants C1 and 
	// will be assigned to this circuit. J2 & J1 both
	// want C0 and will be assigned to this circuit. J0
	// wants C0, but C0 has been filled, so will take his
	// second pick of C1. 
	
	cs := mockCircuitsCase1()
	js := mockJugglersCase1()
	r := make(Jugglers, 0)
	capacity := len(js) / len(cs)
	s := &Solution{cs, js, r, capacity}
	
	s.AssignJugglers()
	results := s.Circuits.Publish()
	
	sort.Strings(results)
	if results[0] != "C0 J2 C0:9 C1:18, J1 C0:6 C1:12" || results[1] != "C1 J3 C1:24 C0:12, J0 C0:3 C1:6" {
		t.Errorf("Assign failed to assign properly; expected results[0] == \"C0 J2 C0:9 C1:18, J1 C0:6 C1:12\" but got %v and expected results[1] == \"C1 J3 C1:24 C0:12, J0 C0:3 C1:6\" but got %v", results[0], results[1])
	}
	
} 

func TestAssignJugglersCase2Content(t *testing.T) {
	cs := mockCircuitsCase2()
	js := mockJugglersCase2()
	r := make(Jugglers, 0)
	capacity := len(js) / len(cs)
	s := &Solution{cs, js, r, capacity}
	
	s.AssignJugglers()
	results := s.Circuits.Publish()
	
	sort.Strings(results)
	
	C1Wrong := (results[0] != "C1 J6 C4:72 C1:18, J3 C3:27 C1:9")
	C2Wrong := (results[1] != "C2 J2 C2:12 C4:24")
	C3Wrong := (results[2] != "C3 J5 C3:45 C4:60, J4 C3:36 C2:24")
	C4Wrong := (results[3] != "C4 J8 C4:96 C3:72, J7 C4:84 C2:42")
	
	if (C1Wrong || C2Wrong || C3Wrong || C4Wrong) {
		t.Errorf("Assign failed to assign properly.")
	}
}

func TestAssignJugglersCase3(t *testing.T) {
	cs := mockCircuitsCase3()
	js := mockJugglersCase3()
	r := make(Jugglers, 0)
	capacity := len(js) / len(cs)
	s := &Solution{cs, js, r, capacity}
	
	s.AssignJugglers()
	results := s.Circuits.Publish()
	
	sort.Strings(results)
	
//	for _, v := range(results) {
//		fmt.Println(v)
//	}
	
}

