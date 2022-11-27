package visitor

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	root := NewDirectory("root", []Entry{})

	bin := NewDirectory("bin", []Entry{})
	viFile := NewFile("vi", 100)
	emacsFile := NewFile("emacs", 200)
	bin.Add(viFile)
	bin.Add(emacsFile)

	root.Add(bin)

	root.Add(NewFile(".vimrc", 300))

	if root.GetSize() != 600 {
		t.Fatal()
	}
}

func TestXml(t *testing.T) {
	root := NewDirectory("root", []Entry{})

	bin := NewDirectory("bin", []Entry{})
	viFile := NewFile("vi", 100)
	emacsFile := NewFile("emacs", 200)
	bin.Add(viFile)
	bin.Add(emacsFile)

	root.Add(bin)

	xmlVisitor := &XMLVisitor{}
	ret := root.Accept(xmlVisitor)
	rets := ret.(string)

	erase := func(s string) string {
		for _, t := range []string{" ", "\n", "\t"} {
			s = strings.ReplaceAll(s, t, "")
		}
		return s
	}

	expected := `
<directory>
   <name>root</name>
   <contents>
      <directory>
         <name>bin</name>
         <contents>
            <file>vi</file>
            <file>emacs</file>
         </contents>
      </directory>
   </contents>
</directory>	
`
	assert.Equal(t, erase(expected), rets)
}
