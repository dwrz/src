package line

import (
	"testing"
)

type GlyphIndexTest struct {
	data   string
	glyph  int
	index  int
	target int
}

func TestGlyphIndex(t *testing.T) {
	var tests = []GlyphIndexTest{
		{
			data:   "",
			index:  0,
			glyph:  0,
			target: 0,
		},
		{
			data:   "\n",
			index:  0,
			glyph:  0,
			target: 0,
		},
		{
			data:   "\t🤔",
			index:  0,
			glyph:  0,
			target: 0,
		},
		{
			data:   "\t🤔",
			index:  0,
			glyph:  0,
			target: 1,
		},
		{
			data:   "\t🤔",
			index:  1,
			glyph:  8,
			target: 8,
		},
		{
			data:   "\t🤔**",
			index:  5,  // \t = 0, 🤔 = 1-4, * = 5
			glyph:  10, // \t = 0, 🤔 = 8, * = 10.
			target: 10, // *
		},
		{
			data:   "\t🤔*",
			index:  6,  // \t = 0, 🤔 = 1-4, * = 5
			glyph:  11, // \t = 0, 🤔 = 8, * = 10.
			target: 11, // *
		},

		{
			data:   "你是谁？",
			index:  3, // 你 size is 3 bytes.
			glyph:  2, // 你 width is 2.
			target: 2, // 是
		},
		{
			data:   "磨杵成针",
			index:  12, // 针
			glyph:  8,  // 针
			target: 8,  // Doesn't exist; return end of last rune.
		},
	}

	for n, test := range tests {
		i, g := New(test.data).FindGlyphIndex(test.target)
		if i != test.index {
			t.Errorf(
				"#%d: expected index %d, but got %d",
				n, test.index, i,
			)
		}
		if g != test.glyph {
			t.Errorf(
				"#%d: expected glyph %d, but got %d",
				n, test.glyph, g,
			)
		}
	}
}

type WidthTest struct {
	data     string
	expected int
}

func TestWidth(t *testing.T) {
	var tests = []WidthTest{
		{
			data:     "",
			expected: 0,
		},
		{
			data:     "\n",
			expected: 0,
		},
		{
			data:     "\t",
			expected: 8,
		},
		{
			data:     "Hello, World!",
			expected: 13,
		},
		{
			data:     "Hello, 世界!",
			expected: 12,
		},
		{
			data:     "💻💩",
			expected: 4,
		},
	}

	for _, test := range tests {
		if w := New(test.data).Width(); w != test.expected {
			t.Errorf(
				"expected width %d, but got %d",
				test.expected, w,
			)
		}
	}
}
