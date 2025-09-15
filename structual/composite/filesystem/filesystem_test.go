package filesystem

import (
	"testing"
)

func TestFileSystem(t *testing.T) {
	home := NewDirectory("home")
	ubuntu := NewDirectory("ubuntu")
	_ = ubuntu.AddChild(NewFile("a.txt", 2048))
	_ = home.AddChild(ubuntu)
	bin := NewDirectory("bin")
	_ = bin.AddChild(NewFile("a.sh", 1024))
	_ = bin.AddChild(NewFile("b.sh", 1024))
	opt := NewDirectory("opt")
	root := NewDirectory("/", home, bin, opt)
	totalSize := root.GetSize()
	if totalSize != 4096 {
		t.Errorf("got %d, want 4096", totalSize)
	}
}
