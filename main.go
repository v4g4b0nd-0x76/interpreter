package main

import (
	"fmt"
	"os"
	"os/user"

	"custom-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

// this is the way and interpreter works
// source code -> tokens -> abs(abstract syntax tree)
// the first transformation for source code to token is called lexical analysis or lexing that is done by tokenizer or scanner or lexer(all are same thing)
// token ar small easily ctegorizable data structures that are then fed to the parser which does the second transformation adn turns the token into an abstract syntax tree
// parsers take source code as input (either text or tokens) and produce a data structure which represents this source code while building the data structs they unavoidably analyse the input checking that it conforms to be expected
