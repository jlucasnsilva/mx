// Package mx provides a type-safe, composable, and streamable HTML rendering engine
// for Go applications. Inspired by JSX and hyperscript, it uses Go functions to build
// the HTML tree in a declarative style.
package mx

import (
	"fmt"
	"html"
	"io"
	"strings"
)

// Node represents an HTML node being rendered.
type Node struct {
	Writer  io.Writer         // where HTML output is written to (usually http.ResponseWriter)
	err     error             // stores the first write error encountered during rendering
	indent  int               // used for pretty printing indentation in dev mode
	DevMode bool              // enables pretty printing and dev features like data-node
	writeFn func(func(*Node)) // optional hook to intercept element rendering (used by WrapEach)
}

// Text writes escaped text.
func Text(text string) func(*Node) {
	return func(n *Node) {
		n.write(html.EscapeString(text))
	}
}

// Textf writes formatted escaped text.
func Textf(format string, args ...any) func(*Node) {
	return func(n *Node) {
		t := fmt.Sprintf(format, args...)
		n.write(html.EscapeString(t))
	}
}

// Raw writes unescaped HTML. Use with caution.
func Raw(raw string) func(*Node) {
	return func(n *Node) {
		n.write(raw)
	}
}

// WrapEach intercepts a function that renders multiple sibling elements and wraps each one
// with a provided wrapper (e.g., a div with a class).
func WrapEach(n *Node, wrapper func(*Node, func(*Node)), children func(*Node)) {
	proxy := &Node{
		Writer:  n.Writer,
		DevMode: n.DevMode,
		indent:  n.indent,
		err:     n.err,
		writeFn: func(inner func(*Node)) {
			wrapper(n, inner)
		},
	}
	children(proxy)
}

// el renders an HTML element with tag, attributes, and children.
func (n *Node) el(tag string, attr Attr, children ...func(*Node)) {
	if n.writeFn != nil {
		n.writeFn(func(child *Node) {
			child.el(tag, attr, children...)
		})
		return
	}

	if n.DevMode {
		n.writeIndent()
	}

	n.write("<" + tag)
	if attr != nil {
		attrs := attr.Attributes()
		if attrs != "" {
			n.write(" " + attrs)
		}
	}
	if n.DevMode {
		n.write(fmt.Sprintf(` data-node="%s"`, tag))
	}

	if isVoidTag(tag) {
		n.write(" />")
		if n.DevMode {
			n.write("\n")
		}
		return
	}

	n.write(">")
	if n.DevMode {
		n.write("\n")
		n.indent++
	}
	for _, child := range children {
		if child != nil {
			child(n)
		}
	}
	if n.DevMode {
		n.indent--
		n.writeIndent()
	}
	n.write("</" + tag + ">")
	if n.DevMode {
		n.write("\n")
	}
}

// write safely writes to the writer and sets error if occurred.
func (n *Node) write(s string) {
	if n.err != nil {
		return
	}
	_, n.err = io.WriteString(n.Writer, s)
}

// writeIndent writes the indentation spaces.
func (n *Node) writeIndent() {
	if n.DevMode {
		n.write(strings.Repeat("  ", n.indent))
	}
}

// Err returns the write error if any occurred during rendering.
func (n *Node) Err() error {
	return n.err
}

// HTML5 void tags
var voidTags = map[string]struct{}{
	"area":   {},
	"base":   {},
	"br":     {},
	"col":    {},
	"embed":  {},
	"hr":     {},
	"img":    {},
	"input":  {},
	"link":   {},
	"meta":   {},
	"source": {},
	"track":  {},
	"wbr":    {},
}

// isVoidTag checks if a tag is self-closing.
func isVoidTag(tag string) bool {
	_, ok := voidTags[tag]
	return ok
}
