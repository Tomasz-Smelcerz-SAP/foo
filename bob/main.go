package main
import (
	"fmt"
	atypes "github.com/Tomasz-Smelcerz-SAP/foo/alice/pkg/types"
)

func main() {
	fmt.Println("Hello, World!")
	Tree := atypes.NewNode(42)
	Tree.SetLeftChild(atypes.NewNode(21))
	Tree.SetRightChild(atypes.NewNode(84))
	//TODO: Uncomment the next line to print the tree structure
	//fmt.Println("Tree structure:", Tree.ToString())
	fmt.Println("We can't print the tree structure yet :(")
}
