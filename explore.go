package main




import (
	"fmt"
)




func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		fmt.Errorf("invalid area name")
	}
	//explore
	return nil
}
