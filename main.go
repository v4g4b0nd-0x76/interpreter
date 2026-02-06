package main

import "fmt"

func main() {
	fmt.Println("custom interpreter")
}

// this is the way and interpreter works
// source code -> tokens -> abs(abstract syntax tree)
// the first transformation for source code to token is called lexical analysis or lexing that is done by tokenizer or scanner or lexer(all are same thing)
// token ar small easily ctegorizable data structures that are then fed to the parser wich does the second transformaiton adn turns the token into an abstract syntax tree
