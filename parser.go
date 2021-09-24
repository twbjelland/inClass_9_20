package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

//assign each operator a token
const INT_LIT = 10
const IDENT = 11
const ASSIGN_OP = 20
const ADD_OP = 21
const SUB_OP = 22
const MULT_OP = 23
const DIV_OP = 24
const LEFT_PAREN = 25
const RIGHT_PAREN = 26

var charClass int
var nextChar rune
var lexeme [100]rune
var lexLen int
var token int
var nextToken int
var in_fp *bufio.Scanner

var isAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

const LETTER = 0
const DIGIT = 1
const UNKNOWN = 99
const EOF = 98

//function to lookup operators and return the token that they were assigned
func lookup(ch rune) int {
	switch ch {
	case '(':
		addChar()
		nextToken = LEFT_PAREN
		break
	case ')':
		addChar()
		nextToken = RIGHT_PAREN
		break
	case '+':
		addChar()
		nextToken = ADD_OP
		break
	case '-':
		addChar()
		nextToken = SUB_OP
		break
	case '*':
		addChar()
		nextToken = MULT_OP
		break
	case '/':
		addChar()
		nextToken = DIV_OP
		break
	default:
		addChar()
		nextToken = EOF
		break
	}
	return nextToken
}

//a function to add Char to the lexeme
func addChar() {
	lexeme[lexLen] = nextChar
	lexLen++
	lexeme[lexLen] = 0
}

//function to get next character of input and determine its character class
func getChar() {
	in_fp.Scan()
	nextString := in_fp.Text()
	if len(nextString) > 0 {
		nextChar = []rune(nextString)[0]
		if (65 <= nextChar && nextChar <= 90) || (97 <= nextChar && nextChar <= 122) {
			charClass = LETTER
		} else if unicode.IsDigit(nextChar) {
			charClass = DIGIT
		} else {
			charClass = UNKNOWN
		}
	} else {
		charClass = EOF
	}

}

//function to call getChar until it returns a non-whitespace character
func getNonBlank() {
	for unicode.IsSpace(nextChar) && charClass != EOF {
		getChar()
	}
}

//a lexical analyzer for arithmetic expressions
func lex() int {
	lexLen = 0
	getNonBlank()
	switch charClass {
	case LETTER:
		addChar()
		getChar()
		for charClass == LETTER || charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = IDENT
		break
	case DIGIT:
		addChar()
		getChar()
		for charClass == DIGIT {
			addChar()
			getChar()
		}
		nextToken = INT_LIT
		break
	case UNKNOWN:
		lookup(nextChar)
		getChar()
		break
	case EOF:
		nextToken = EOF
		lexLen = 3
		lexeme[0] = 'E'
		lexeme[1] = 'O'
		lexeme[2] = 'F'
	}
	fmt.Print("Next token is: ")
	fmt.Print(nextToken)
	fmt.Print(", Next lexeme is: ")
	for i := 0; i < lexLen; i++ {
		fmt.Print(string(lexeme[i]))
	}
	fmt.Println()

	return nextToken
}

// parses strings in the language generated by the rule:
// <expr> -> <term> {(+ | -) <term>}
func expr() {
	fmt.Printf("Enter <expr>\n")

	// parse the first term
	term()

	for nextToken == ADD_OP || nextToken == SUB_OP {
		lex()
		term()
	}
	fmt.Printf("Exit <expr>\n")
}

/*
parses strings in the language generated by the rule :
<term> -> <factor> {(* | /) <factor>}
*/
func term() {
	fmt.Printf("Enter <term>\n")

	factor()

	for nextToken == MULT_OP || nextToken == DIV_OP {
		lex()
		factor()
	}
	fmt.Printf("Exit <term>\n")
}

/*
parses strings in the language generated by the rule:
<factor> -> id | int_constant | ( <expr> )
*/
func factor() {
	fmt.Printf("Enter <factor>\n")

	// determine which RHS
	if nextToken == IDENT || nextToken == INT_LIT {
		lex()

		/*
			if the RHS is ( <expr> ), call lex to pass over the left parentheses, call expr, and check for the right parentheses
		*/

	} else {
		if nextToken == LEFT_PAREN {
			lex()
			expr()
			if nextToken == RIGHT_PAREN {
				lex()
			} else {
				errors.New("oops")
			}
		} else {
			errors.New("oops")
		}
	}
	fmt.Printf("Exit <factor>\n")
}

func parse() {
	expr()
}

func main() {
	filename := "test.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	in_fp = bufio.NewScanner(strings.NewReader(inputdata))
	in_fp.Split(bufio.ScanRunes)

	for charClass != EOF {
		lex()
		parse()
	}

}
