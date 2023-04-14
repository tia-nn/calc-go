package main

import (
	"errors"
	"regexp"
)

func Tokenize(code string) []TokenInfo {
	target := tokenizingCode{code: code, pos: 0}
	tokens := make([]TokenInfo, 0, 100)

	for {
		if target.code == "" {
			break
		}

		if target.code[0] == ' ' {
			target.code = target.code[1:]
			target.pos += 1
		} else if t, err := target.tokenizeOperator(); err == nil {
			tokens = append(tokens, t)
		} else if t, err := target.tokenizeNum(); err == nil {
			tokens = append(tokens, t)
		} else {
			panic("unknown token")
		}

	}

	return tokens
}

func (code *tokenizingCode) tokenizeOperator() (TokenInfo, error) {
	re := regexp.MustCompile(`^[+-/*]`)
	tokenStr := re.FindString(code.code)

	if tokenStr != "" {
		token := TokenInfo{
			Token: Token{
				TokenType: int(tokenStr[0]),
				Value:     tokenStr,
				Position:  Position{CharAt: code.pos, length: len(tokenStr)},
			},
		}
		code.code = code.code[len(tokenStr):]
		code.pos += len(tokenStr)
		return token, nil
	} else {
		return nilToken, errors.New("operator unmatched")
	}
}

func (code *tokenizingCode) tokenizeNum() (TokenInfo, error) {
	re := regexp.MustCompile(`^[1-9][0-9]*`)
	tokenStr := re.FindString(code.code)

	if tokenStr != "" {
		token := TokenInfo{
			Token: Token{
				TokenType: NUMBER,
				Value:     tokenStr,
				Position:  Position{CharAt: code.pos, length: len(tokenStr)},
			},
		}
		code.code = code.code[len(tokenStr):]
		code.pos += len(tokenStr)
		return token, nil
	} else {
		return nilToken, errors.New("number unmatced")
	}
}

type tokenizingCode struct {
	code string
	pos  int
}

type TokenInfo struct {
	Token  Token
	Remain string
}

type Token struct {
	TokenType int
	Value     string
	Position  Position
}

const (
	NUMBER = 256
)

var nilToken = TokenInfo{
	Token: Token{
		TokenType: 0,
		Value:     "",
		Position:  Position{CharAt: 0},
	},
	Remain: "",
}
