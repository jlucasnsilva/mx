package mx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	voidAttrTestCase struct {
		attr     string
		expected bool
	}

	simpleAttrTestCase struct {
		attr     Attr
		expected string
	}

	attrTestCase struct {
		attr  Attr
		check func(string) bool
	}
)

func TestVoidAttributes(t *testing.T) {
	RegisterVoidAttrs("it-is", "it-also-is")
	testCases := []voidAttrTestCase{
		{attr: "disabled", expected: true},
		{attr: "defer", expected: true},
		{attr: "open", expected: true},
		{attr: "it-is", expected: true},
		{attr: "it-also-is", expected: true},
		{attr: "it-isn't", expected: false},
		{attr: "it-also-isn't", expected: false},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("checks if '%v' is an void attribute", tc.attr)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, isVoidAttr(tc.attr))
		})
	}
}

func TestAttr(t *testing.T) {
	t.Run("Test simple attributes", func(t *testing.T) {
		testCases := []simpleAttrTestCase{
			{
				attr:     S(`type="text" class="input"`),
				expected: `type="text" class="input"`,
			},
			{
				attr:     M{},
				expected: "",
			},
			{
				attr:     M{"": "grid"},
				expected: "",
			},
			{
				attr:     M{"class": "grid"},
				expected: `class="grid"`,
			},
			{
				attr:     N{},
				expected: "",
			},
			{
				attr:     N{"": {}},
				expected: "",
			},
			{
				attr:     Slice{S(`class="input" type="number"`), N{"placeholder": {"Hello, world!": false}}},
				expected: `class="input" type="number" placeholder=""`,
			},
			{
				attr:     Slice{S(`class="input" type="number"`), nil, N{"placeholder": {"Hello, world!": true}}},
				expected: `class="input" type="number" placeholder="Hello, world!"`,
			},
			{
				attr:     Class("container", "responsive", "red"),
				expected: `class="container responsive red"`,
			},
		}

		for _, tc := range testCases {
			name := fmt.Sprintf("checks if '%T' renders attributes correctly", tc.attr)
			t.Run(name, func(t *testing.T) {
				assert.Equal(t, tc.expected, tc.attr.Attributes())
			})
		}
	})

	t.Run("Test map-based attributes", func(t *testing.T) {
		testCases := []attrTestCase{
			{
				attr: M{"class": "input", "type": "text"},
				check: func(s string) bool {
					return s == `class="input" type="text"` ||
						s == `type="text" class="input"`
				},
			},
			{
				attr: N{
					"type":     {"text": true},
					"disabled": {"": true},
				},
				check: func(s string) bool {
					return s == `type="text" disabled` ||
						s == `disabled type="text"`
				},
			},
		}

		for _, tc := range testCases {
			name := fmt.Sprintf("checks if '%T' renders attributes correctly", tc.attr)
			t.Run(name, func(t *testing.T) {
				attr := tc.attr.Attributes()
				assert.True(t, tc.check(attr))
			})
		}
	})
}
