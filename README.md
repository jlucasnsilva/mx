# mx — Streamable HTML Templates in Go

`mx` is a minimal, type-safe, and composable HTML templating engine for Go. Write HTML directly in Go using functions — no string templating, no runtime parsing, and no magic.

```go
func HomePage(n *mx.Node) {
	n.Div(mx.S(`class="container"`), func(n *mx.Node) {
		n.H1(nil, mx.Text("Welcome!"))
	})
}
```

## 🚀 Features

- ⚡ Streamed rendering (`io.Writer`)
- ✅ Safe HTML escaping by default
- 🧱 Typed HTML elements
- 🧩 Composable components (just Go functions)
- 🧪 Dev mode: pretty output
- 🎯 Zero dependencies

---

## 📦 Installation

```bash
go get github.com/jlucasnsilva/mx
```

---

## 📐 Usage

### Define a Page

```go
func HelloPage(n *mx.Node) {
	n.Div(mx.S(`class="greeting"`), func(n *mx.Node) {
		n.H1(nil, mx.Text("Hello, world!"))
	})
}
```

### Render It

```go
func handler(w http.ResponseWriter, r *http.Request) {
	node := &mx.Node{Writer: w}
	HelloPage(node)
	if err := mx.Error(node); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
```

---

## 💻 Dev Mode (Pretty Output + Profiling)

```go
node := &mx.Node{Writer: os.Stdout, DevMode: true}
HelloPage(node)
```

Outputs:

```html
<div class="greeting">
  <h1>Hello, world!</h1>
</div>
```

---

## 🧩 Create Components

```go
func Card(title string, body func(*mx.Node)) func(*mx.Node) {
	return func(n *mx.Node) {
		n.Div(mx.S(`class="card"`), func(n *mx.Node) {
			n.H2(nil, mx.Text(title))
			body(n)
		})
	}
}
```

---

## 🧪 Testing

Test rendering via `strings.Builder`:

```go
var b strings.Builder
node := &mx.Node{Writer: &b}
HomePage(node)
assert.Contains(t, b.String(), "<div")
```

---

## 🪵 Debugging / Error Handling

```go
if err := mx.Error(node); err != nil {
	log.Println("render error:", err)
}
```

---

## 🧱 `WrapEach` Example

Wrap each child node in a component-defined wrapper:

```go
func Grid(n *mx.Node, children func(*mx.Node)) {
	n.Div(mx.S(`class="grid"`), func(n *mx.Node) {
		mx.WrapEach(n, func(n *mx.Node, content func(*mx.Node)) {
			n.Div(mx.S(`class="grid-item"`), content)
		}, children)
	})
}
```

Use like this:

```go
Grid(n, func(n *mx.Node) {
	n.Text("Item 1")
	n.Text("Item 2")
	n.Text("Item 3")
})
```

Output:

```html
<div class="grid">
  <div class="grid-item">Item 1</div>
  <div class="grid-item">Item 2</div>
  <div class="grid-item">Item 3</div>
</div>
```

---

## 🛠️ mxgen — CLI to convert HTML → mx Go functions

`mxgen` is a simple CLI tool that reads an HTML file and generates a Go function using the `mx` template system.

### ✨ Usage

```bash
go run ./mxgen input.html ComponentName
```

### 📤 Example

Given this `hero.html`:

```html
<div class="hero">
  <h1 id="title">Welcome!</h1>
</div>
```

Running:

```bash
go run ./mxgen hero.html Hero
```

Outputs:

```go
func Hero(n *mx.Node) {
	n.Div(mx.M{"class": "hero"}, func(n *mx.Node) {
		n.H1(mx.M{"id": "title"}, func(n *mx.Node) {
			n.Text("Welcome!")
		})
	})
}
```

- All tags and attributes are supported.
- Output is written to `stdout`.
- Can be used in CI, generators, or quick prototyping.



---

## 🧰 Install mxgen as a CLI tool

You can compile `mxgen` as a standalone binary and use it from the terminal.

### 🔨 Build and Install

```bash
go install ./mxgen
```

This will install the `mxgen` binary into your `$GOBIN` (e.g. `$HOME/go/bin/mxgen`).

### 🚀 Usage

```bash
mxgen -in=input.html -name=ComponentName
```

It will print:

```go
func ComponentName(n *mx.Node) {
  // converted HTML
}
```

Use it in CI pipelines, scaffolding tools, or rapid prototyping.
