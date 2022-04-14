package main

import (
	"bufio"
	"errors"
	"fmt"
	_ "github.com/mattn/go-tty"
	"github.com/muesli/termenv"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

type terminal struct {
	p termenv.Profile
}

func initTerminal() *terminal {
	return &terminal{
		p: termenv.EnvColorProfile(),
	}
}

type Promptable interface {
	~uint64 | ~int64 | ~float64 | ~rune | ~string | ~bool
	String() string
	TypeOf() string
}

type Prompt[T Promptable] struct {
	prompt     string
	defaultVal T
	maxVal     T // maximum value the user can provide (optional)
	minVal     T // minimum value the user can provide (optional)
	optimalVal T // value the user must provide to avoid problems (optional)
}

// NewPrompt returns a pointer to a Prompt and to a ResultFromUser
// it accepts a Type param for both these structs, with R being passed to the constructor for ResultFromUser
// it requires a prompt, the defaultVal of type Resultable, and a pointer to the TurnLog
// it may have additional params of type Promptable passed, in order to fill out the optional values of Prompt:
// 1st member of otherVals is the max value, 2nd is the min value, 3rd is the required value
// func NewPrompt[T Promptable, R Resultable](prompt string, defaultVal T, turnLog *TurnLog, otherVals ...T) *Prompt[T, R] {

// func NewPrompt[T Promptable](prompt string, defaultVal T, turnLog *TurnLog, otherVals ...T) *Prompt[T] {
func NewPrompt[T Promptable](prompt string, defaultVal T, turnLog *TurnLog, otherVals ...T) (string, T) {
	var max, min, optimal T
	switch len(otherVals) {
	case 3:
		max = otherVals[0]
		min = otherVals[1]
		optimal = otherVals[2]
	case 2:
		max = otherVals[0]
		min = otherVals[1]
	case 1:
		max = otherVals[0]
	}

	promptStruct := &Prompt[T]{
		prompt:     prompt,
		defaultVal: defaultVal,
		maxVal:     max,
		minVal:     min,
		optimalVal: optimal,
	}
	return buildPrompt[T](promptStruct)
}

// type Resultable interface {
//	~uint64 | ~int64 | ~float64 | ~rune | ~bool | ~string | ~PromptBool
// }

type PromptUInt64 uint64
type PromptInt64 int64
type PromptFloat float64
type PromptRune rune
type PromptBool bool
type PromptString string

func (r PromptUInt64) String() string {
	return fmt.Sprintf("%d", r)
}

func (r PromptUInt64) TypeOf() string {
	return fmt.Sprintf("%s", reflect.TypeOf(r).Kind())
}

func (r PromptInt64) String() string {
	return fmt.Sprintf("%d", r)
}

func (r PromptInt64) TypeOf() string {
	return fmt.Sprintf("%s", reflect.TypeOf(r).Kind())
}

func (r PromptFloat) String() string {
	return fmt.Sprintf("%d", r)
}

func (r PromptFloat) TypeOf() string {
	return fmt.Sprintf("%s", reflect.TypeOf(r).Kind())
}

func (r PromptRune) String() string {
	return fmt.Sprintf("%d", r)
}

func (r PromptRune) TypeOf() string {
	return fmt.Sprintf("%s", reflect.TypeOf(r).Kind())
}

func (r PromptBool) String() string {
	return fmt.Sprintf("%t", r)
}

func (r PromptBool) TypeOf() string {
	return fmt.Sprintf("%s", reflect.TypeOf(r).Kind())
}

func (r PromptString) String() string {
	return fmt.Sprintf("%s", r)
}

func (r PromptString) TypeOf() string {
	return fmt.Sprintf("%s", reflect.TypeOf(r).Kind())
}

// buildPrompt returns a string of the prompt, and the type of the value the prompt is requesting from the user
// a prompt is composed of the following parts:
//
// 1: the basic prompt string, typically a question like "do you wish to do X?" or "how many credits will you use?"
//
// 2: a default value which will be accepted as input if the player simply hits enter, signified by `[]`
//
// 3: (optional) a maximum value, for example in cases where the player can only purchase so many of an item, if present signified by `()`
//
// 4: (optional) a minimum value // CONSIDER:  possibly not necessary
//
// 5: (optional) an optimal value representing the best choice for the player to make, this will often be represented
// by the default value but where present separately, signified by `<>`
//
// example using 1, 2, 3, 5:
// 	"How much energy will you pay to import? (254,214) <5000> [0] => "
// this tells a player that they can import 254214 units of energy, that optimally they will import 5000, and that if they simply hit ENTER, they will import 0
func buildPrompt[T Promptable](p *Prompt[T]) (string, T) {
	defaultVal := p.defaultVal.String()
	optimalVal := p.optimalVal.String()
	if fmt.Sprintf("%s", p.optimalVal) == "0" || fmt.Sprintf("%s", p.optimalVal) == "" {
		optimalVal = fmt.Sprintf("(%s)", p.defaultVal)
	}

	prompt := fmt.Sprintf("%s (%s) [%s] => ", p.prompt, optimalVal, defaultVal)

	switch p.defaultVal.TypeOf() {
	case "string":
		return prompt, p.defaultVal
	case "rune":
		return prompt, p.defaultVal
	case "uint64":
		return prompt, p.defaultVal
	case "int64":
		return prompt, p.defaultVal
	case "float64":
		return prompt, p.defaultVal
	case "bool":
		// FIXME: This sucks badly
		var dVal string
		if p.defaultVal.String() == "true" {
			dVal = "y"
		} else {
			dVal = "n"
		}
		prompt = fmt.Sprintf("%s (%s) [%s] => ", p.prompt, optimalVal, dVal)
		return prompt, p.defaultVal
	}
	return "", p.defaultVal // should never reach here
}

/*
func promptWithOptimalVal() string {

}

func promptWithOptimalAndMaxVal() string {

}

func promptNoOptimalVal() string {

}
*/
func getString[T Promptable](p string, defaultVal T) string {
	// fmt.Printf("%s\n", p)
	i := bufio.NewReader(os.Stdin)
	input, _ := i.ReadString('\n')
	return input
}

func getChar() rune {
	var i rune
	if OperatingSystem == "windows" {
		_, _ = fmt.Scanf("%c\n", &i)
	} else {
		_, _ = fmt.Scanf("%c", &i)
	}
	return i
}

func getUInt[T Promptable](p string, defaultVal T) uint64 {
	var i string
	var err error
	var n int

	fmt.Printf("%s", p)

	if OperatingSystem == "windows" {
		n, err = fmt.Scanf("%s\n", &i)
	} else {
		n, err = fmt.Scanf("%s", &i)
	}

	if n == 0 { // player hit ENTER
		x, _ := strconv.ParseUint(defaultVal.String(), 10, 64)
		return x
	}
	var x uint64
	x, err = strconv.ParseUint(i, 10, 64)
	if errors.Is(err, strconv.ErrSyntax) { // invalid input
		return getUInt(p, defaultVal)
	}

	return x
}

func getBool[T Promptable](p string, defaultVal T) bool {
	var i string

	fmt.Printf("%s", p)

	if OperatingSystem == "windows" {
		_, _ = fmt.Scanf("%s\n", &i)
	} else {
		_, _ = fmt.Scanf("%s", &i)
	}

	// player hits <RET>, return the default specified in p
	if i == "" {
		x := defaultVal.String()
		y, _ := strconv.ParseBool(x)
		return y
	}
	// anything other than <RET> or y/n, try again
	rx, _ := regexp.Compile("([yYnN])")
	if !rx.MatchString(i) {
		termenv.ClearLines(2)
		return getBool(p, defaultVal)
	}

	if i == "y" || i == "Y" {
		// r.result = true
		return true
	} else {
		return false
	}
}
