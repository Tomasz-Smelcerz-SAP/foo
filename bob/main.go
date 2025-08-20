package main
import (
	"fmt"
	atypes "github.com/Tomasz-Smelcerz-SAP/foo/alice/v2/pkg/types"
)

func main() {
	fmt.Println("Hello, World!")
	Tree := atypes.T(atypes.L(21), 42, atypes.L(84))
	fmt.Println("Tree structure:", Tree.String())
}
