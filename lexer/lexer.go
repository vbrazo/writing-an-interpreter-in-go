package lexer

type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	readpos int  // current read position in input (after current char)
	chr     byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readchar()
	return l
}

func (l *Lexer) readchar() {
	if l.readpos >= len(l.input) {
		l.chr = 0
	} else {
		l.chr = l.input[l.readpos]
		l.pos = l.readpos
		l.readpos++
	}
}
