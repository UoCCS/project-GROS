package lexer

import (
	"unicode"
	"unicode/utf8"
)

// Token represents a parsed token.
type Token struct {
	Kind    TokenKind
	Len     int
	Literal string
}

// TokenKind represents common lexeme types.
type TokenKind int

const (
	LineComment TokenKind = iota
	BlockComment
	Whitespace
	Ident
	InvalidIdent
	RawIdent
	UnknownPrefix
	UnknownPrefixLifetime
	RawLifetime
	GuardedStrPrefix
	Literal
	Lifetime
	Semi
	Comma
	Dot
	OpenParen
	CloseParen
	OpenBrace
	CloseBrace
	OpenBracket
	CloseBracket
	At
	Pound
	Tilde
	Question
	Colon
	Dollar
	Eq
	Bang
	Lt
	Gt
	Minus
	And
	Or
	Plus
	Star
	Slash
	Caret
	Percent
	Unknown
	Eof
)

// DocStyle represents the style of documentation comments.
type DocStyle int

const (
	Outer DocStyle = iota
	Inner
)

// LiteralKind represents the literal types supported by the lexer.
type LiteralKind int

const (
	Int LiteralKind = iota
	Float
	Char
	Byte
	Str
	ByteStr
	CStr
	RawStr
	RawByteStr
	RawCStr
)

// Base represents the base of numeric literal encoding according to its prefix.
type Base int

const (
	Binary      Base = 2
	Octal       Base = 8
	Decimal     Base = 10
	Hexadecimal Base = 16
)

// Cursor represents the cursor for parsing tokens.
type Cursor struct {
	input string
	pos   int
}

// NewCursor creates a new cursor for the given input string.
func NewCursor(input string) *Cursor {
	return &Cursor{input: input}
}

// Bump advances the cursor and returns the next character.
func (c *Cursor) Bump() (rune, bool) {
	if c.pos >= len(c.input) {
		return 0, false
	}
	r, size := utf8.DecodeRuneInString(c.input[c.pos:])
	c.pos += size
	return r, true
}

// First returns the first character without advancing the cursor.
func (c *Cursor) First() (rune, bool) {
	if c.pos >= len(c.input) {
		return 0, false
	}
	r, _ := utf8.DecodeRuneInString(c.input[c.pos:])
	return r, true
}

// AdvanceToken parses a token from the input string.
func (c *Cursor) AdvanceToken() Token {
	startPos := c.pos
	firstChar, ok := c.Bump()
	if !ok {
		return Token{Kind: Eof, Len: 0}
	}

	var tokenKind TokenKind
	switch firstChar {
	case '/':
		if nextChar, ok := c.First(); ok && nextChar == '/' {
			tokenKind = c.LineComment()
		} else if nextChar == '*' {
			tokenKind = c.BlockComment()
		} else {
			tokenKind = Slash
		}
	case ' ', '\t', '\n', '\r':
		tokenKind = c.Whitespace()
	case 'r':
		if nextChar, ok := c.First(); ok && nextChar == '#' {
			tokenKind = c.RawIdent()
		} else {
			tokenKind = c.IdentOrUnknownPrefix()
		}
	case 'b':
		tokenKind = c.COrByteString()
	case 'c':
		tokenKind = c.COrByteString()
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		tokenKind = c.Number(firstChar)
	case ';':
		tokenKind = Semi
	case ',':
		tokenKind = Comma
	case '.':
		tokenKind = Dot
	case '(':
		tokenKind = OpenParen
	case ')':
		tokenKind = CloseParen
	case '{':
		tokenKind = OpenBrace
	case '}':
		tokenKind = CloseBrace
	case '[':
		tokenKind = OpenBracket
	case ']':
		tokenKind = CloseBracket
	case '@':
		tokenKind = At
	case '#':
		tokenKind = Pound
	case '~':
		tokenKind = Tilde
	case '?':
		tokenKind = Question
	case ':':
		tokenKind = Colon
	case '$':
		tokenKind = Dollar
	case '=':
		tokenKind = Eq
	case '!':
		tokenKind = Bang
	case '<':
		tokenKind = Lt
	case '>':
		tokenKind = Gt
	case '-':
		tokenKind = Minus
	case '&':
		tokenKind = And
	case '|':
		tokenKind = Or
	case '+':
		tokenKind = Plus
	case '*':
		tokenKind = Star
	case '^':
		tokenKind = Caret
	case '%':
		tokenKind = Percent
	case '\'':
		tokenKind = c.LifetimeOrChar()
	case '"':
		tokenKind = c.DoubleQuotedString()
	default:
		if unicode.IsLetter(firstChar) || firstChar == '_' {
			tokenKind = c.IdentOrUnknownPrefix()
		} else {
			tokenKind = Unknown
		}
	}

	return Token{Kind: tokenKind, Len: c.pos - startPos, Literal: c.input[startPos:c.pos]}
}

// LineComment parses a line comment token.
func (c *Cursor) LineComment() TokenKind {
	for {
		if ch, ok := c.Bump(); !ok || ch == '\n' {
			break
		}
	}
	return LineComment
}

// BlockComment parses a block comment token.
func (c *Cursor) BlockComment() TokenKind {
	depth := 1
	for depth > 0 {
		ch, ok := c.Bump()
		if !ok {
			break
		}
		if ch == '*' {
			if nextChar, ok := c.First(); ok && nextChar == '/' {
				c.Bump()
				depth--
			}
		} else if ch == '/' {
			if nextChar, ok := c.First(); ok && nextChar == '*' {
				c.Bump()
				depth++
			}
		}
	}
	return BlockComment
}

// Whitespace parses a whitespace token.
func (c *Cursor) Whitespace() TokenKind {
	for {
		if ch, ok := c.First(); !ok || !unicode.IsSpace(ch) {
			break
		}
		c.Bump()
	}
	return Whitespace
}

// RawIdent parses a raw identifier token.
func (c *Cursor) RawIdent() TokenKind {
	c.Bump() // Consume '#'
	for {
		if ch, ok := c.First(); !ok || !unicode.IsLetter(ch) {
			break
		}
		c.Bump()
	}
	return RawIdent
}

// IdentOrUnknownPrefix parses an identifier or unknown prefix token.
func (c *Cursor) IdentOrUnknownPrefix() TokenKind {
	for {
		if ch, ok := c.First(); !ok || !unicode.IsLetter(ch) {
			break
		}
		c.Bump()
	}
	return Ident
}

// COrByteString parses a C or byte string token.
func (c *Cursor) COrByteString() TokenKind {
	// Implementation omitted for brevity
	return Unknown
}

// Number parses a numeric literal token.
func (c *Cursor) Number(firstChar rune) TokenKind {
	// Implementation omitted for brevity
	return Literal
}

// LifetimeOrChar parses a lifetime or character literal token.
func (c *Cursor) LifetimeOrChar() TokenKind {
	// Implementation omitted for brevity
	return Literal
}

// DoubleQuotedString parses a double-quoted string token.
func (c *Cursor) DoubleQuotedString() TokenKind {
	for {
		ch, ok := c.Bump()
		if !ok || ch == '"' {
			break
		}
		if ch == '\\' {
			c.Bump() // Consume escaped character
		}
	}
	return Literal
}

// Tokenize creates an iterator that produces tokens from the input string.
func Tokenize(input string) []Token {
	cursor := NewCursor(input)
	var tokens []Token
	for {
		token := cursor.AdvanceToken()
		if token.Kind == Eof {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens
}
