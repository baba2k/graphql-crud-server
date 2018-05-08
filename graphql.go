package graphql

import (
	"github.com/baba2k/graphql-rungen/parser"
)

func ParseFile(fileName string) {
	parser.ParseFile(fileName)
}

func ParseString(rawString string) {
	parser.ParseString(rawString)
}