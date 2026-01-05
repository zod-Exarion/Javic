package external

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/zod-Exarion/javic/emitter"
	"github.com/zod-Exarion/javic/lexer"
	"github.com/zod-Exarion/javic/parser"
)

type customTheme struct {
	fyne.Theme
}

func (c customTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 18
	}
	// Return default for all other sizes
	return c.Theme.Size(name)
}

func newCustomTheme() fyne.Theme {
	return customTheme{Theme: theme.DefaultTheme()}
}

func NewGUI() {
	a := app.New()
	a.Settings().SetTheme(newCustomTheme())
	w := a.NewWindow("QBASIC → Java Converter")

	// Left: QBASIC input
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapWord
	input.SetPlaceHolder("Type QBASIC code here...")

	// Right: output
	output := widget.NewMultiLineEntry()
	// output.SetReadOnly(true)

	// Buttons
	convertBtn := widget.NewButton("Convert", func() {
		src := input.Text

		tokens := lexer.Lex(src)
		ast := parser.Parse(tokens)

		java := emitter.Emit(ast)

		output.SetText(java)
	})

	tokensBtn := widget.NewButton("Tokens", func() {
		src := input.Text
		tokens := lexer.Lex(src)

		out := lexer.DisplayTokens(tokens)
		output.SetText(out)
	})

	astBtn := widget.NewButton("AST", func() {
		src := input.Text
		tokens := lexer.Lex(src)
		stmts := parser.Parse(tokens)

		out := parser.DisplayStatements(stmts)
		output.SetText(out)
	})

	header := widget.NewLabel("QBASIC TO JAVA TRANSLATOR")
	centeredContent := container.New(layout.NewCenterLayout(), header)
	buttons := container.NewHBox(convertBtn, tokensBtn, astBtn)
	buttonbar := container.New(layout.NewCenterLayout(), buttons)
	topbar := container.NewHBox(centeredContent, buttonbar)

	split := container.NewHSplit(input, output)
	split.SetOffset(0.5)

	w.SetContent(
		container.NewBorder(
			topbar, nil, nil, nil, split,
		),
	)

	w.Resize(fyne.NewSize(900, 600))
	w.ShowAndRun()
}
