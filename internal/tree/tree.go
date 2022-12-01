package tree

// TODO：思路，
// 参考地址：https://juejin.cn/post/7027585316185702414
// 参考实现： du 命令

type Node struct {
	Path    string
	Name    string
	Size    int
	SubNode []*Node
	IsDir   bool
}

func Start(dir string) {

}
