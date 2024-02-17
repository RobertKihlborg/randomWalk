package pkg

import (
	"fmt"
	"os"
	"strconv"
)

var builders = map[string]WalkerBuilder{
	"*": generalBuilder,
}

func ValidBuilders() []string {
	var keys []string
	for k, _ := range builders {
		keys = append(keys, k)
	}
	return keys
}

func BuildWalker(args []string) (Walker, string, int, error) {
	builder, ok := builders[args[0]]
	if !ok {
		return nil, "", 0, fmt.Errorf("invalid builder name. choose one in list:  %v", ValidBuilders())
	}

	walker, name, usedArgs, err := builder(args[1:])
	return walker, name, 1 + usedArgs, err
}

func generalBuilder(args []string) (Walker, string, int, error) {
	if len(args) < 1 {
		return nil, "", 0, fmt.Errorf("no walk type specified")
	}
	switch args[0] {
	case "space":
		w, name, usedArgs, err := generalSpaceBuilder(args[1:])
		return w, name, usedArgs + 1, err
	case "grid":
		w, name, usedArgs, err := generalGridBuilder(args[1:])
		return w, name, usedArgs + 1, err
	default:
		return nil, "", 0, fmt.Errorf("invalid walk type")
	}
}

func generalGridBuilder(args []string) (Walker, string, int, error) {
	if len(args) < 1 {
		return nil, "", 0, fmt.Errorf("no dimension specified")
	}
	if len(args) < 2 {
		return nil, "", 0, fmt.Errorf("no intersection policy specified")
	}

	var dim int
	var err error
	if dim, err = strconv.Atoi(args[0]); err != nil {
		return nil, "", 0, err
	}

	switch args[1] {
	case "i":
		return func(n int) Walk {
			return basicIGridWalk(dim, n)
		}, fmt.Sprintf("grid%vDi", dim), 2, nil
	case "a":
		return func(n int) Walk {
			return nil

		}, "", 2, nil
	default:
		return nil, "", 0, fmt.Errorf("intersection policy needs to be 'i' or 'a'")
	}

}

func generalSpaceBuilder(args []string) (Walker, string, int, error) {
	switch len(args) {
	case 0:
		return nil, "", 0, fmt.Errorf("no dimension specified")

	case 1:
		return nil, "", 0, fmt.Errorf("no intersection policy specified")
	}

	var dim int
	var err error
	if dim, err = strconv.Atoi(os.Args[0]); err != nil {
		return nil, "", 0, err
	}
	switch args[1] {
	case "i":
		return func(n int) Walk {
			return nil
		}, "", 2, nil
	case "a":
		if len(args) < 3 {
			return nil, "", 0, fmt.Errorf("no d specified for self avoiding space walk")
		}
		var d float64
		if d, err = strconv.ParseFloat(os.Args[0], 64); err != nil {
			return nil, "", 0, err
		}
		print(d) //TODO

		return func(n int) Walk {
			return nil

		}, "", 2, nil
	default:
		print(dim) //TODO
		return nil, "", 0, fmt.Errorf("intersection policy needs to be 'i' or 'a'")

	}

}
