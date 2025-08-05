package mx

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	tagsTestCase struct {
		expected  string
		component func(*Node)
	}
)

func TestTags(t *testing.T) {
	testCases := []tagsTestCase{
		{expected: `<!DOCTYPE html>`, component: func(n *Node) { n.DocType() }},
		{expected: `<a></a>`, component: func(n *Node) { n.A(nil) }},
		{expected: `<abbr></abbr>`, component: func(n *Node) { n.Abbr(nil) }},
		{expected: `<address></address>`, component: func(n *Node) { n.Address(nil) }},
		{expected: `<area />`, component: func(n *Node) { n.Area(nil) }},
		{expected: `<article></article>`, component: func(n *Node) { n.Article(nil) }},
		{expected: `<aside></aside>`, component: func(n *Node) { n.Aside(nil) }},
		{expected: `<audio></audio>`, component: func(n *Node) { n.Audio(nil) }},
		{expected: `<b></b>`, component: func(n *Node) { n.B(nil) }},
		{expected: `<base />`, component: func(n *Node) { n.Base(nil) }},
		{expected: `<bdi></bdi>`, component: func(n *Node) { n.Bdi(nil) }},
		{expected: `<bdo></bdo>`, component: func(n *Node) { n.Bdo(nil) }},
		{expected: `<blockquote></blockquote>`, component: func(n *Node) { n.Blockquote(nil) }},
		{expected: `<body></body>`, component: func(n *Node) { n.Body(nil) }},
		{expected: `<br />`, component: func(n *Node) { n.Br(nil) }},
		{expected: `<button></button>`, component: func(n *Node) { n.Button(nil) }},
		{expected: `<canvas></canvas>`, component: func(n *Node) { n.Canvas(nil) }},
		{expected: `<caption></caption>`, component: func(n *Node) { n.Caption(nil) }},
		{expected: `<cite></cite>`, component: func(n *Node) { n.Cite(nil) }},
		{expected: `<code></code>`, component: func(n *Node) { n.Code(nil) }},
		{expected: `<col />`, component: func(n *Node) { n.Col(nil) }},
		{expected: `<colgroup></colgroup>`, component: func(n *Node) { n.Colgroup(nil) }},
		{expected: `<data></data>`, component: func(n *Node) { n.Data(nil) }},
		{expected: `<datalist></datalist>`, component: func(n *Node) { n.Datalist(nil) }},
		{expected: `<dd></dd>`, component: func(n *Node) { n.Dd(nil) }},
		{expected: `<del></del>`, component: func(n *Node) { n.Del(nil) }},
		{expected: `<details></details>`, component: func(n *Node) { n.Details(nil) }},
		{expected: `<dfn></dfn>`, component: func(n *Node) { n.Dfn(nil) }},
		{expected: `<dialog></dialog>`, component: func(n *Node) { n.Dialog(nil) }},
		{expected: `<div></div>`, component: func(n *Node) { n.Div(nil) }},
		{expected: `<dl></dl>`, component: func(n *Node) { n.Dl(nil) }},
		{expected: `<dt></dt>`, component: func(n *Node) { n.Dt(nil) }},
		{expected: `<em></em>`, component: func(n *Node) { n.Em(nil) }},
		{expected: `<embed />`, component: func(n *Node) { n.Embed(nil) }},
		{expected: `<fieldset></fieldset>`, component: func(n *Node) { n.Fieldset(nil) }},
		{expected: `<figcaption></figcaption>`, component: func(n *Node) { n.Figcaption(nil) }},
		{expected: `<figure></figure>`, component: func(n *Node) { n.Figure(nil) }},
		{expected: `<footer></footer>`, component: func(n *Node) { n.Footer(nil) }},
		{expected: `<form></form>`, component: func(n *Node) { n.Form(nil) }},
		{expected: `<h1></h1>`, component: func(n *Node) { n.H1(nil) }},
		{expected: `<h2></h2>`, component: func(n *Node) { n.H2(nil) }},
		{expected: `<h3></h3>`, component: func(n *Node) { n.H3(nil) }},
		{expected: `<h4></h4>`, component: func(n *Node) { n.H4(nil) }},
		{expected: `<h5></h5>`, component: func(n *Node) { n.H5(nil) }},
		{expected: `<h6></h6>`, component: func(n *Node) { n.H6(nil) }},
		{expected: `<head></head>`, component: func(n *Node) { n.Head(nil) }},
		{expected: `<header></header>`, component: func(n *Node) { n.Header(nil) }},
		{expected: `<hr />`, component: func(n *Node) { n.Hr(nil) }},
		{expected: `<html></html>`, component: func(n *Node) { n.Html(nil) }},
		{expected: `<i></i>`, component: func(n *Node) { n.I(nil) }},
		{expected: `<iframe></iframe>`, component: func(n *Node) { n.Iframe(nil) }},
		{expected: `<img />`, component: func(n *Node) { n.Img(nil) }},
		{expected: `<input />`, component: func(n *Node) { n.Input(nil) }},
		{expected: `<ins></ins>`, component: func(n *Node) { n.Ins(nil) }},
		{expected: `<kbd></kbd>`, component: func(n *Node) { n.Kbd(nil) }},
		{expected: `<label></label>`, component: func(n *Node) { n.Label(nil) }},
		{expected: `<legend></legend>`, component: func(n *Node) { n.Legend(nil) }},
		{expected: `<li></li>`, component: func(n *Node) { n.Li(nil) }},
		{expected: `<link />`, component: func(n *Node) { n.Link(nil) }},
		{expected: `<main></main>`, component: func(n *Node) { n.Main(nil) }},
		{expected: `<map></map>`, component: func(n *Node) { n.Map(nil) }},
		{expected: `<mark></mark>`, component: func(n *Node) { n.Mark(nil) }},
		{expected: `<meta />`, component: func(n *Node) { n.Meta(nil) }},
		{expected: `<meter></meter>`, component: func(n *Node) { n.Meter(nil) }},
		{expected: `<nav></nav>`, component: func(n *Node) { n.Nav(nil) }},
		{expected: `<noscript></noscript>`, component: func(n *Node) { n.Noscript(nil) }},
		{expected: `<object></object>`, component: func(n *Node) { n.Object(nil) }},
		{expected: `<ol></ol>`, component: func(n *Node) { n.Ol(nil) }},
		{expected: `<optgroup></optgroup>`, component: func(n *Node) { n.Optgroup(nil) }},
		{expected: `<option></option>`, component: func(n *Node) { n.Option(nil) }},
		{expected: `<output></output>`, component: func(n *Node) { n.Output(nil) }},
		{expected: `<p></p>`, component: func(n *Node) { n.P(nil) }},
		{expected: `<picture></picture>`, component: func(n *Node) { n.Picture(nil) }},
		{expected: `<pre></pre>`, component: func(n *Node) { n.Pre(nil) }},
		{expected: `<progress></progress>`, component: func(n *Node) { n.Progress(nil) }},
		{expected: `<q></q>`, component: func(n *Node) { n.Q(nil) }},
		{expected: `<rp></rp>`, component: func(n *Node) { n.Rp(nil) }},
		{expected: `<rt></rt>`, component: func(n *Node) { n.Rt(nil) }},
		{expected: `<ruby></ruby>`, component: func(n *Node) { n.Ruby(nil) }},
		{expected: `<s></s>`, component: func(n *Node) { n.S(nil) }},
		{expected: `<samp></samp>`, component: func(n *Node) { n.Samp(nil) }},
		{expected: `<script></script>`, component: func(n *Node) { n.Script(nil) }},
		{expected: `<section></section>`, component: func(n *Node) { n.Section(nil) }},
		{expected: `<select></select>`, component: func(n *Node) { n.Select(nil) }},
		{expected: `<slot></slot>`, component: func(n *Node) { n.Slot(nil) }},
		{expected: `<small></small>`, component: func(n *Node) { n.Small(nil) }},
		{expected: `<source />`, component: func(n *Node) { n.Source(nil) }},
		{expected: `<span></span>`, component: func(n *Node) { n.Span(nil) }},
		{expected: `<strong></strong>`, component: func(n *Node) { n.Strong(nil) }},
		{expected: `<style></style>`, component: func(n *Node) { n.Style(nil) }},
		{expected: `<sub></sub>`, component: func(n *Node) { n.Sub(nil) }},
		{expected: `<summary></summary>`, component: func(n *Node) { n.Summary(nil) }},
		{expected: `<sup></sup>`, component: func(n *Node) { n.Sup(nil) }},
		{expected: `<svg></svg>`, component: func(n *Node) { n.SVG(nil) }},
		{expected: `<table></table>`, component: func(n *Node) { n.Table(nil) }},
		{expected: `<tbody></tbody>`, component: func(n *Node) { n.Tbody(nil) }},
		{expected: `<td></td>`, component: func(n *Node) { n.Td(nil) }},
		{expected: `<template></template>`, component: func(n *Node) { n.Template(nil) }},
		{expected: `<textarea></textarea>`, component: func(n *Node) { n.Textarea(nil) }},
		{expected: `<tfoot></tfoot>`, component: func(n *Node) { n.Tfoot(nil) }},
		{expected: `<th></th>`, component: func(n *Node) { n.Th(nil) }},
		{expected: `<thead></thead>`, component: func(n *Node) { n.Thead(nil) }},
		{expected: `<time></time>`, component: func(n *Node) { n.Time(nil) }},
		{expected: `<title></title>`, component: func(n *Node) { n.Title(nil) }},
		{expected: `<tr></tr>`, component: func(n *Node) { n.Tr(nil) }},
		{expected: `<track />`, component: func(n *Node) { n.Track(nil) }},
		{expected: `<u></u>`, component: func(n *Node) { n.U(nil) }},
		{expected: `<ul></ul>`, component: func(n *Node) { n.Ul(nil) }},
		{expected: `<var></var>`, component: func(n *Node) { n.Var(nil) }},
		{expected: `<video></video>`, component: func(n *Node) { n.Video(nil) }},
		{expected: `<wbr />`, component: func(n *Node) { n.Wbr(nil) }},
	}

	for _, tc := range testCases {
		name := fmt.Sprintf("expects: %v", tc.expected)
		t.Run(name, func(t *testing.T) {
			b := &strings.Builder{}
			tc.component(&Node{Writer: b})
			assert.Equal(t, tc.expected, b.String())
		})
	}
}

// TODO
