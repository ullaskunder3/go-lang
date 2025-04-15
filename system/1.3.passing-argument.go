package system

import (
	"flag"
	"fmt"
	"os"
)

func PassingArgument() {
	// steup flags
	name := flag.String("name", "Guest", "Your name")
	rank := flag.Int("age", 0, "Your rank")

	// parse flags
	flag.Parse()

	// remaining non-flag args
	remainingArgs := flag.Args()

	// print logs
	fmt.Println("---flag base input----")
	fmt.Printf("Name: %s\n", *name)
	fmt.Printf("Rank: %d\n", *rank)

	fmt.Println("\n ---positionla ragument----")
	for i, arg := range remainingArgs {
		fmt.Printf("Arg[%d]: %s\n", i, arg)
	}
	fmt.Println("---args info----")
	fmt.Println("argsWithProg:", os.Args)
	fmt.Println("argsWithoutProg:", os.Args[1:])

	if len(os.Args) > 3 {
		fmt.Println("\nos.Args[3]:", os.Args[3])
	} else {
		fmt.Println("\nos.Args[3] not available (length =", len(os.Args), ")")
	}

}
