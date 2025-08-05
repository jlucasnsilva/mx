package mx

import (
	"html"
	"strings"
)

type (
	// Attr defines the interface for HTML attributes.
	Attr interface {
		Attributes() string
	}

	// S is a raw string attribute.
	S string

	// M represents key-value HTML attributes.
	M map[string]string

	// N represents conditional attributes like classes.
	N map[string]map[string]bool

	// Slice allows combining multiple Attrs.
	Slice []Attr
)

// Support custom void attributes
var customVoidAttrs = map[string]bool{}

// Attributes implementation for each Attr type
func (s S) Attributes() string {
	return string(s)
}

func (m M) Attributes() string {
	var b strings.Builder
	i := 0
	for k, v := range m {
		if k == "" {
			continue
		}
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(k)
		if !isVoidAttr(k) || v != "" {
			b.WriteString(`="` + html.EscapeString(v) + `"`)
		}
		i++
	}
	return b.String()
}

func (n N) Attributes() string {
	var b strings.Builder
	i := 0
	for k, conds := range n {
		if k == "" {
			continue
		}
		if i > 0 {
			b.WriteByte(' ')
		}
		if isVoidAttr(k) {
			i++
			ok := false
			for _, flag := range conds {
				ok = ok || flag
			}
			if ok {
				b.WriteString(k)
			}
			continue
		}

		b.WriteString(k)
		var vals []string
		for val, ok := range conds {
			if ok {
				vals = append(vals, html.EscapeString(val))
			}
		}
		b.WriteString(`="` + strings.Join(vals, " ") + `"`)
		i++
	}
	return b.String()
}

func (s Slice) Attributes() string {
	var b strings.Builder
	for i, a := range s {
		if a == nil {
			continue
		}
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(a.Attributes())
	}
	return b.String()
}

// Class creates a class attribute from a list of strings.
func Class(classes ...string) Attr {
	return M{"class": strings.Join(classes, " ")}
}

// RegisterVoidAttrs registers attributes that are treated as void.
func RegisterVoidAttrs(attrs ...string) {
	for _, a := range attrs {
		customVoidAttrs[a] = true
	}
}

// isVoidAttr checks if an attribute is a void attribute.
func isVoidAttr(attr string) bool {
	return attr == "disabled" || attr == "defer" || attr == "open" || customVoidAttrs[attr]
}
