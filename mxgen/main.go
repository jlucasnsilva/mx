// Command mxgen is a CLI tool to convert HTML files to mx Go functions.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	input := flag.String("in", "", "HTML input file")
	name := flag.String("name", "Component", "Component name for Go function")
	flag.Parse()

	if *input == "" || *name == "" {
		log.Fatal("Usage: mxgen -in=input.html -name=ComponentName")
	}

	f, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	node, err := html.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("func %s(n *mx.Node) {\n", *name)
	renderNode(os.Stdout, node.FirstChild, "\t")
	fmt.Println("}")
}

func renderNode(w io.Writer, n *html.Node, indent string) {
	for n != nil {
		switch n.Type {
		case html.ElementNode:
			tag := strings.Title(n.Data)
			attrStr := buildAttr(n.Attr)
			fmt.Fprintf(w, "%sn.%s(%s, func(n *mx.Node) {\n", indent, tag, attrStr)
			renderNode(w, n.FirstChild, indent+"\t")
			fmt.Fprintf(w, "%s})\n", indent)
		case html.TextNode:
			text := strings.TrimSpace(n.Data)
			if text != "" {
				fmt.Fprintf(w, "%sn.Text(\"%s\")\n", indent, escapeString(text))
			}
		}
		n = n.NextSibling
	}
}

func buildAttr(attrs []html.Attribute) string {
	if len(attrs) == 0 {
		return "nil"
	}
	out := make([]string, 0, len(attrs))
	for _, a := range attrs {
		if a.Key != "" {
			out = append(out, fmt.Sprintf("\"%s\": \"%s\"", a.Key, escapeString(a.Val)))
		}
	}
	sort.Strings(out)
	return "mx.M{" + strings.Join(out, ", ") + "}"
}

func escapeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}
