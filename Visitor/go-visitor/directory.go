package visitor

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
func (d *Directory) Accept(v VisitorInterface) any {
	return v.Visit(d)
}

func (d *Directory) Add(e Entry) {
	d.Entries = append(d.Entries, e)
}
