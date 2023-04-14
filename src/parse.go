package main

import (
	"errors"
	"strconv"
)

func Parse(tokens []TokenInfo) (Node, error) {
	p := parsing{tokens: tokens}

	return p.parseAdd()
}

func (tokens *parsing) parseAdd() (Node, error) {
	lhs, err := tokens.parseMul()
	if err != nil {
		return nilNode, err
	}

	for {
		if tokens.consume('+') {
			rhs, err := tokens.parseMul()
			if err != nil {
				panic("require add rhs")
			}
			lhs = AddNode{
				lhs:      lhs,
				rhs:      rhs,
				addType:  int('+'),
				position: Position{CharAt: lhs.getPosition().CharAt, length: rhs.getPosition().CharAt - lhs.getPosition().CharAt + rhs.getPosition().length},
			}
		} else if tokens.consume('-') {
			rhs, err := tokens.parseMul()
			if err != nil {
				panic("require sub rhs")
			}
			lhs = AddNode{
				lhs:      lhs,
				rhs:      rhs,
				addType:  int('-'),
				position: Position{CharAt: lhs.getPosition().CharAt, length: rhs.getPosition().CharAt - lhs.getPosition().CharAt + rhs.getPosition().length},
			}
		} else {
			return lhs, nil
		}
	}
}

func (tokens *parsing) parseMul() (Node, error) {
	lhs, err := tokens.parseNumber()
	if err != nil {
		return nilNode, err
	}

	for {
		if tokens.consume('*') {
			rhs, err := tokens.parseNumber()
			if err != nil {
				panic("require mul rhs")
			}
			lhs = MulNode{
				lhs:      lhs,
				rhs:      rhs,
				mulType:  int('*'),
				position: Position{CharAt: lhs.getPosition().CharAt, length: rhs.getPosition().CharAt - lhs.getPosition().CharAt + rhs.getPosition().length},
			}
		} else if tokens.consume('/') {
			rhs, err := tokens.parseNumber()
			if err != nil {
				panic("require div rhs")
			}
			lhs = MulNode{
				lhs:      lhs,
				rhs:      rhs,
				mulType:  int('/'),
				position: Position{CharAt: lhs.getPosition().CharAt, length: rhs.getPosition().CharAt - lhs.getPosition().CharAt + rhs.getPosition().length},
			}
		} else {
			return lhs, nil
		}
	}
}

func (tokens *parsing) parseNumber() (Node, error) {
	head, err := tokens.getToken()
	if err != nil {
		return nilNode, err
	}
	if head.Token.TokenType == NUMBER {
		val, err := strconv.Atoi(head.Token.Value)
		if err != nil {
			panic("unreachable")
		}
		tokens.tokens = tokens.tokens[1:]
		return NumNode{value: val, position: head.Token.Position}, nil
	} else {
		return NumNode{0, nilPosition}, errors.New("number unmatched")
	}
}

func (tokens *parsing) consume(tokenType int) bool {
	if len(tokens.tokens) > 0 && tokens.tokens[0].Token.TokenType == tokenType {
		tokens.tokens = tokens.tokens[1:]
		return true
	} else {
		return false
	}
}

func (tokens *parsing) getToken() (TokenInfo, error) {
	if len(tokens.tokens) > 0 {
		return tokens.tokens[0], nil
	} else {
		return nilToken, errors.New("EOF")
	}
}

type parsing struct {
	tokens []TokenInfo
}
