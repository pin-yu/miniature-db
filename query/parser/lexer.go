package parser

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type lexer struct {
	input        string
	inputLen     int
	position     int
	readPosition int
	r            rune
}

func newLexer(input string) *lexer {
	l := &lexer{
		input:        strings.ToLower(input),
		inputLen:     len(input),
		position:     0,
		readPosition: 0,
		r:            0,
	}
	l.readChar()
	return l
}

func (l *lexer) NextToken() token {
	if l.r == utf8.RuneError {
		return token{
			Lit:       "",
			TokenType: UNK,
		}
	}

	l.eatSpace()

	if isDelim(l.r) {
		defer l.readChar()
		return token{
			Lit:       string(l.r),
			TokenType: DELIM,
		}
	}

	hasNonAscii := false

	b := strings.Builder{}
	for r := l.r; !isDelim(r) && l.r != 0; r = l.r {
		if r > unicode.MaxASCII {
			hasNonAscii = true
		}
		b.WriteRune(r)
		l.readChar()
	}

	if b.Len() != 0 {
		s := b.String()
		if hasNonAscii {
			return token{
				Lit:       s,
				TokenType: UNK,
			}
		}

		if isKw(s) {
			return token{
				Lit:       b.String(),
				TokenType: KW,
			}
		}
		return token{
			Lit:       b.String(),
			TokenType: ID,
		}
	}
	return token{
		TokenType: END,
	}
}

func (l *lexer) readChar() {
	if l.readPosition >= l.inputLen {
		l.r = 0
		l.position = l.readPosition
		return
	}

	r, size := utf8.DecodeRuneInString(l.input[l.readPosition:])
	l.position = l.readPosition
	l.readPosition += size
	l.r = r
}

// func (l *Lexer) peekChar() rune {
// 	r, _ := utf8.DecodeLastRuneInString(l.input[l.readPosition:])
// 	return r
// }

func (l *lexer) eatSpace() {
	for l.r == ' ' {
		l.readChar()
	}
}
