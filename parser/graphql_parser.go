package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	graphql "github.com/baba2k/graphql-rungen/parser/antlr"
)

// ParseFile parses a graphql schema from file
func ParseFile(fileName string) {
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		panic("could not parse file " + fileName + ": " + err.Error())
	}
	parse(input)
}

// ParseFile parses a graphql schema from string
func ParseString(rawString string) {
	parse(antlr.NewInputStream(rawString))
}

func parse(input interface{}) {
	// create lexer
	lexer := graphql.NewGraphQLLexer(input.(antlr.CharStream))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create parser
	p := graphql.NewGraphQLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	// parse expression
	antlr.ParseTreeWalkerDefault.Walk(NewGraphqlListener(), p.Document())
}
