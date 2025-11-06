package mx

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	nodeTestCase struct {
		component func(*Node)
		expected  string
		indent    bool
	}
)

func TestNode(t *testing.T) {
	testCases := []nodeTestCase{
		{
			indent:   false,
			expected: "foo",
			component: func(n *Node) {
				Text("foo")(n)
			},
		},
		{
			indent:   false,
			expected: "foo bar",
			component: func(n *Node) {
				Textf("foo %v", "bar")(n)
			},
		},
		{
			indent: true,
			expected: `<form>
  <input type="number" />
</form>
`,
			component: func(n *Node) {
				n.Form(nil, func(n *Node) {
					n.Input(S(`type="number"`))
				})
			},
		},
		{
			indent:   false,
			expected: "<!DOCTYPE html>",
			component: func(n *Node) {
				Raw("<!DOCTYPE html>")(n)
			},
		},
		{
			indent:   false,
			expected: `<main class="container"><p>Hello, world!</p></main>`,
			component: func(n *Node) {
				n.Main(S(`class="container"`), func(n *Node) {
					n.P(nil, Text("Hello, world!"))
				})
			},
		},
		{
			indent: true,
			expected: `<div class="grid">
  <div class="grid-item">
    Hello, world!
  </div>
  <div class="grid-item">
    Hello, world!
  </div>
  <div class="grid-item">
    Hello, world!
  </div>
</div>
`,
			component: func(n *Node) {
				n.Div(S(`class="grid"`), func(n *Node) {
					n.Div(S(`class="grid-item"`), Text("Hello, world!"))
					n.Div(S(`class="grid-item"`), Textf("Hello, world!"))
					n.Div(S(`class="grid-item"`), Raw("Hello, world!"))
				})
			},
		},
		{
			indent:   false,
			expected: `<div class="grid"><div class="grid-item"><p>Hello, world!</p></div><div class="grid-item"><p>Hello, world!</p></div></div>`,
			component: func(n *Node) {
				n.Div(S(`class="grid"`), func(n *Node) {
					WrapEach(
						n,
						func(n *Node, content func(*Node)) {
							n.Div(S(`class="grid-item"`), content)
						},
						func(n *Node) {
							n.P(nil, Text("Hello, world!"))
							n.P(nil, Text("Hello, world!"))
						},
					)
				})
			},
		},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("expects: %v", tc.expected)
		t.Run(name, func(t *testing.T) {
			b := &strings.Builder{}
			tc.component(&Node{Writer: b, DevMode: tc.indent})
			assert.Equal(t, tc.expected, b.String())
		})
	}
}
