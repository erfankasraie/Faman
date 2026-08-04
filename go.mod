module github.com/faman-project/faman

go 1.24.0

toolchain go1.24.4

require (
	github.com/charmbracelet/glamour v0.7.0
	github.com/charmbracelet/lipgloss v0.12.1
	github.com/mattn/go-isatty v0.0.20
	github.com/mattn/go-runewidth v0.0.15
	github.com/spf13/cobra v1.8.1
	github.com/yuin/goldmark v1.7.4
	golang.org/x/term v0.40.0
)

require (
	github.com/alecthomas/chroma/v2 v2.8.0 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/charmbracelet/x/ansi v0.1.4 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.25 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/yuin/goldmark-emoji v1.0.2 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
)

replace golang.org/x/net => github.com/golang/net v0.50.0

replace golang.org/x/sys => github.com/golang/sys v0.22.0

replace golang.org/x/term => github.com/golang/term v0.22.0

replace gopkg.in/yaml.v3 => github.com/go-yaml/yaml v0.0.0-20220527083530-f6f7691b1fde // v3.0.1 pinned by commit due to sandbox network allowlist

replace gopkg.in/check.v1 => github.com/go-check/check v0.0.0-20161208181325-20d25e280405
