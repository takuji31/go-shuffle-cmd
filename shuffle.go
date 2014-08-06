package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/xtgo/sortutil"
)

func main() {
	var first = flag.String("first", "", "force set first element")
	var last = flag.String("last", "", "force set last element")
	flag.Parse()
	var args = sort.StringSlice(flag.Args())
	rand.Seed(time.Now().Unix())
	sortutil.Shuffle(args)
	if *first != "" {
		fmt.Println(*first)
	}
	for i := 0; i < args.Len(); i++ {
		fmt.Println(args[i])
	}
	if *last != "" {
		fmt.Println(*last)
	}
}
