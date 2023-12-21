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
	readRune     int
	r            rune
}

func newLexer(input string) *lexer {
	l := &lexer{
		input:        strings.ToLower(input),
		inputLen:     len(input),
		position:     0,
		readPosition: 0,
		readRune:     0,
		r:            0,
	}
	l.readChar()
	return l
}

func (l *lexer) NextToken() token {
	if l.r == utf8.RuneError {
		return token{
			Lit:       "",
			EndPos:    l.readRune,
			TokenType: UNK,
		}
	}

	l.eatSpace()

	if isDelim(l.r) {
		defer l.readChar()
		return token{
			Lit:       string(l.r),
			EndPos:    l.readRune,
			TokenType: DELIM,
		}
	}

	hasNonAscii := false

	b := strings.Builder{}
	for l.r != 0 {
		if l.r > unicode.MaxASCII {
			hasNonAscii = true
		}
		b.WriteRune(l.r)

		// peek
		if isDelim(l.peekChar()) {
			break
		}

		l.readChar()
	}

	if b.Len() != 0 {
		s := b.String()
		if hasNonAscii {
			defer l.readChar()
			return token{
				Lit:       s,
				EndPos:    l.readRune,
				TokenType: UNK,
			}
		}

		if isKw(s) {
			defer l.readChar()
			return token{
				Lit:       b.String(),
				EndPos:    l.readRune,
				TokenType: KW,
			}
		}
		defer l.readChar()
		return token{
			Lit:       b.String(),
			EndPos:    l.readRune,
			TokenType: ID,
		}
	}
	return token{
		EndPos:    l.readRune,
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
	l.readRune += 1
	l.r = r
}

func (l *lexer) peekChar() rune {
	if l.readPosition >= l.inputLen {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return r
}

func (l *lexer) eatSpace() {
	for l.r == ' ' {
		l.readChar()
	}
}
