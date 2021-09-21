package main

import (
    "fmt"
    "log"
    "os"
    "unicode"
    "regexp"
)

var charClass int
var next_char rune
var lexeme string
var lexLen int
var token int
var nextToken int
var in_fp

var isAlpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

const LETTER = 0
const DIGIT = 1
const UNKNOWN = 99

const INT_LIT = 10
const IDENT = 11
const ASSIGN_OP = 20
const ADD_OP = 21
const SUB_OP = 22
const MULT_OP = 23
const DIV_OP = 24
const LEFT_PAREN = 25
const RIGHT_PAREN = 26


//define grammar

//int literal
//identifier
//assign_op
//add_op
//sub_op
//mult_op
//div_op
//left_paren
//right_paren

func addChar() {
    lexeme[lexLen++] = nextChar
    lexeme[lexLen] = 0
}

func getChar(){
    if((nextChar = getc(in_fp)) = EOF) {
        if (isAlpha(nextChar))
            charClass = LETTER
        else if (unicode.isDigit(nextChar))
            charClass = DIGIT
            else charClass = UKNOWN;
    }
    else
        charClass = EOF;
}

func getNonBlank(){
    if unicode.IsSpace(next_char){
        getChar()
    }

}

func main() {
    
    //read in file
    content, err := os.ReadFile("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))
    
    //iterate over all characters
    
        //find lexemes
    
        //convert to tokens
    
        //enforce rules
    
}
