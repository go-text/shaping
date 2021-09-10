package shaping

import (
	"github.com/go-text/di"
	"github.com/go-text/font"
	"golang.org/x/image/math/fixed"
)

type Input interface {
	// Text returns the characters to be shaped.
	//
	// `itemOffset` is the position of the first character from `text` that will be shaped, and
	// `itemLength` is the number of character to shape (-1 means the end of the slice).
	//
	// When shaping part of a larger text (e.g. a run of text from a paragraph),
	// instead of passing just the substring corresponding to the run, it is preferable to pass the whole
	// paragraph and specify the run start and length as `itemOffset` and
	// `itemLength`, respectively, to give the shaper the full context to be able,
	// for example, to do cross-run Arabic shaping or properly handle combining marks at start of run.
	Text() (text []rune, itemOffset, itemLength int)
	// Direction returns the directionality of the text.
	Direction() di.Direction
	// Face returns the font face to render the text in.
	Face() font.Face
	// Size returns the size of the font, eg. 14.
	// TODO is this a scaled value, exact pixels, or display dependand?
	Size() fixed.Int26_6
}
