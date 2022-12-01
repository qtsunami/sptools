package tree

// TODO：思路，

type Node struct {
	Path    string
	Name    string
	Size    int
	SubNode []*Node
	IsDir   bool
}

func Start(dir string) {

}
