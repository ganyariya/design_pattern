package composite

type Directory struct {
	Name    string
	Entries []Entry
}

func NewDirectory(name string, entires []Entry) *Directory {
	return &Directory{Name: name, Entries: entires}
}

func (d *Directory) GetSize() int {
	sz := 0
	for _, entry := range d.Entries {
		sz += entry.GetSize()
	}
	return sz
}
func (d *Directory) GetName() string {
	return d.Name
}
func (d *Directory) GetList(prefix string) string {
	ret := ""
	cur := prefix + "/" + d.GetName()
	for i, entry := range d.Entries {
		ret += entry.GetList(cur)
		if i != len(d.Entries)-1 {
			ret += "\n"
		}
	}
	return ret
}

func (d *Directory) Add(e Entry) {
	d.Entries = append(d.Entries, e)
}
