package main

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {

	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
