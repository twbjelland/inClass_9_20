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
		nextToken = 0 //change to eof
		break
	}
	return nextToken
}

func addChar() {
	lexeme[lexLen+1] = nextChar
	lexeme[lexLen] = 0
}

func getChar(ch rune) {
	//to do : implement getChar
	nextChar = ch
}

func getNonBlank(ch rune) {
	for unicode.IsSpace(nextChar) {
		getChar(ch)
	}
}

func lex(ch rune) int {
	lexLen = 0
	getNonBlank(ch)
	switch charClass {
	case LETTER:
		addChar()
		getChar(ch) //what should I pass in here?
		for charClass == LETTER || charClass == DIGIT {
			addChar()
			getChar(ch) //what should I pass in here?
		}
		nextToken = IDENT
		break
	case DIGIT:
		addChar()
		getChar(ch) //what should I pass in here?
		for charClass == DIGIT {
			addChar()
			getChar(ch) //what should I pass in here?
		}
		nextToken = INT_LIT
		break
	case UNKNOWN:
		lookup(nextChar)
		getChar(ch) //what should I pass in here?
		break
	case EOF:
		nextToken = EOF
		lexeme[0] = 'E'
		lexeme[1] = 'O'
		lexeme[2] = 'F'
		lexeme[3] = 0
		break
	}

	fmt.Println("Next token is: ")
	fmt.Println("%v", nextToken)
	return nextToken
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

	for in_fp.Scan() {
		lex(in_fp.Text())
	}
}

//implement lex function and figure out how to pass char throughout
