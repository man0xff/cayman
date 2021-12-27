package core_test

import (
	// "fmt"
	"testing"

	"cayman/internal/core"
	// "fmt"
	// "context"
	"github.com/davecgh/go-spew/spew"
	// "github.com/stretchr/testify/require"
)

func TestCore(t *testing.T) {
	c, err := core.New(&core.Params{
		ConfigFile: "/home/okovalev/my/cayman_game.conf.yaml",
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(c.ListCategories())
}
