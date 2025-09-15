package filesystem

import "errors"

type FileSystem interface {
	GetName() string
	GetSize() int64
	AddChild(child FileSystem) error
	RemoveChild(child FileSystem) error
}

type File struct {
	name string
	size int64
}

func (f *File) GetName() string {
	return f.name
}
func (f *File) GetSize() int64 {
	return f.size
}
func (f *File) AddChild(child FileSystem) error {
	return errors.New("not supported")
}
func (f *File) RemoveChild(child FileSystem) error {
	return errors.New("not supported")
}

func NewFile(name string, size int64) FileSystem {
	return &File{name: name, size: size}
}

type Directory struct {
	name     string
	children []FileSystem
}

func (d *Directory) GetName() string {
	return d.name
}
func (d *Directory) GetSize() int64 {
	total := int64(0)
	for _, child := range d.children {
		total += child.GetSize()
	}
	return total
}
func (d *Directory) AddChild(child FileSystem) error {
	d.children = append(d.children, child)
	return nil
}
func (d *Directory) RemoveChild(child FileSystem) error {
	for i, c := range d.children {
		if c == child {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
	return errors.New("not found")
}

func NewDirectory(name string, system ...FileSystem) FileSystem {
	return &Directory{name: name, children: system}
}
