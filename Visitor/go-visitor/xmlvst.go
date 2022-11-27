package visitor

import "fmt"

type XMLVisitor struct {
}

func (v *XMLVisitor) Visit(e Entry) any {
	switch e := e.(type) {
	case *File:
		return v.VisitFile(e)
	case *Directory:
		return v.VisitDirectory(e)
	}
	return ""
}

func (v *XMLVisitor) VisitFile(e *File) string {
	return fmt.Sprintf("<file>%s</file>", e.GetName())
}

func (v *XMLVisitor) VisitDirectory(e *Directory) string {
	tpl := `<directory><name>%s</name><contents>%s</contents></directory>`
	contents := ""
	for _, ee := range e.Entries {
		ret := ee.Accept(v)
		contents += ret.(string)
	}
	return fmt.Sprintf(tpl, e.GetName(), contents)
}
