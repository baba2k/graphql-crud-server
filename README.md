# graphql-rungen
GraphQL library for golang

## What is graphql-rungen?
graphql-rungen is a GraphQL runtime generator library written in golang  
for building graphql servers fast and simple. No need for extra models  
or any overhead. Just load a graphql schema, overwrite the resolvers  
and/or use a standard resolver interface and start the server. See section   
"Getting Started" for more details.

## Getting Started
To install the library, run:  
`go get github.com/baba2k/graphql-rungen`

## Examples
See "examples" dir for more examples

## Dependencies
ANTLR http://github.com/antlr/antlr4/runtime/Go/antlr 

## Additinal Information
### ANTLR lexer & parser
* ANTLR (ANother Tool for Language Recognition) is a wide spread and powerful parser generator for reading, processing, executing, or translating structured text or binary files

* EBNF (Extended Backus–Naur form) syntax to define the grammar

* ANTLR is an Adaptive LL(\*) or ALL(\*) parser which works top-down. The generated code for a LL(\*) parser is more understandable than a LALR parser (bottom-up and table driven)

* Listener mechanism works independently by an ANTLR-provided walker object (whereas old visitor methods must walk their children with explicit visit calls)
## License
See file "LICENSE" for licence information  