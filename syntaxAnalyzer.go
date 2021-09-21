package main

import (
    "fmt"
    "log"
    "os"
    "unicode"
)

var next_char rune

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
