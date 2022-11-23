package composite

import (
	"fmt"
	"strings"
	"testing"
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

func TestList(t *testing.T) {
	root := NewDirectory("root", []Entry{})

	bin := NewDirectory("bin", []Entry{})
	viFile := NewFile("vi", 100)
	emacsFile := NewFile("emacs", 200)
	bin.Add(viFile)
	bin.Add(emacsFile)

	python := NewDirectory("python", []Entry{})
	py3File := NewFile("python3", 500)
	python.Add(py3File)
	bin.Add(python)

	root.Add(bin)

	root.Add(NewFile(".vimrc", 300))

	if strings.Trim(root.GetList(""), " \n") != strings.Trim(`
/root/bin/vi
/root/bin/emacs
/root/bin/python/python3
/root/.vimrc
`, " \n") {
		fmt.Println(root.GetList(""))
		t.Fatal()
	}

}
