package parser

import (
	"fmt"

	"github.com/UoCCS/project-GROS/lexer"
)

type Node any

type Program struct {
	Statements []Node
}

type Function struct {
	Name       string
	Parameters []string
	Body       []Node
}

type Variable struct {
	Name  string
	Value Node
}

type Expression struct {
	Left     Node
	Operator string
	Right    Node
}

type Literal struct {
	Value string
}

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() (*Program, error) {
	program := &Program{}
	for p.pos < len(p.tokens) {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		program.Statements = append(program.Statements, stmt)
	}
	return program, nil
}

func (p *Parser) parseStatement() (Node, error) {
	token := p.currentToken()
	switch token.Kind {
	case lexer.Ident:
		return p.parseFunction()
	default:
		return nil, fmt.Errorf("unexpected token: %v", token)
	}
}

func (p *Parser) parseFunction() (*Function, error) {
	name := p.currentToken().Literal
	p.nextToken() // consume function name

	if p.currentToken().Kind != lexer.OpenParen {
		return nil, fmt.Errorf("expected '(', got %v", p.currentToken())
	}
	p.nextToken() // consume '('

	parameters := []string{}
	for p.currentToken().Kind != lexer.CloseParen {
		if p.currentToken().Kind != lexer.Ident {
			return nil, fmt.Errorf("expected identifier, got %v", p.currentToken())
		}
		parameters = append(parameters, p.currentToken().Literal)
		p.nextToken()
		if p.currentToken().Kind == lexer.Comma {
			p.nextToken() // consume ','
		}
	}
	p.nextToken() // consume ')'

	if p.currentToken().Kind != lexer.OpenBrace {
		return nil, fmt.Errorf("expected '{', got %v", p.currentToken())
	}
	p.nextToken() // consume '{'

	body := []Node{}
	for p.currentToken().Kind != lexer.CloseBrace {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		body = append(body, stmt)
	}
	p.nextToken() // consume '}'

	return &Function{Name: name, Parameters: parameters, Body: body}, nil
}

func (p *Parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *Parser) nextToken() {
	p.pos++
}
