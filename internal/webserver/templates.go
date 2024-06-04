package webserver

import (
	"html/template"
	"io"

	"github.com/Laellekoenig/boi/internal/emulator"
	"github.com/Laellekoenig/boi/internal/emulator/cpu"
	"github.com/labstack/echo/v4"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.New("t").Funcs(template.FuncMap{
			"romName": func(e *emulator.Emulator) interface{} {
				return e.RomName()
			},
			"pc": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetPc()
			},
			"sp": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetSp()
			},
			"a": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegA)
			},
			"f": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegF)
			},
			"b": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegB)
			},
			"c": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegC)
			},
			"d": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegD)
			},
			"e": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegE)
			},
			"h": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegH)
			},
			"l": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetRegister(cpu.RegL)
			},
			"memory": func(e *emulator.Emulator) interface{} {
				return e.Memory.GetMemory(e.Cpu.Pc())
			},
			"flags": func(e *emulator.Emulator) interface{} {
				return e.Cpu.GetFlags()
			},
			"pastOps": func(e *emulator.Emulator) interface{} {
				return e.Cpu.PastOps
			},
		}).ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}
