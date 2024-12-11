package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/krbreyn/filesage"
)

func main() {
	watchCmd := flag.NewFlagSet("watch", flag.ExitOnError)

	watchRecurse := watchCmd.Bool("r", false, "recursive dirs")

	if len(os.Args) < 2 {
		fmt.Println("ERROR: expected command")
		return
	}

	switch os.Args[1] {
	case "watch":
		watchCmd.Parse(os.Args[2:])
		files := watchCmd.Args()

		addToWatchlist(files, *watchRecurse)

	default:
		fmt.Println("ERROR: invalid command")
	}
}

func addToWatchlist(files []string, watchRecurse bool) {
	watchlist := filesage.Watchlist{}

	if watchRecurse {
		fmt.Println("recurse")
	}

	for _, file := range files {
		stat, err := os.Stat(file)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("does not exist:", file)
				continue
			} else {
				fmt.Println("err", err)
				return
			}
		}

		absFilePath, err := filepath.Abs(file)
		if err != nil {
			fmt.Println("error getting absolute path", err)
		}

		if stat.IsDir() {
			watchlist.AddWatchlistDir(watchRecurse, absFilePath)
		} else {
			watchlist.AddWatchlistFile(absFilePath)
		}
	}
	fmt.Println(watchlist)
}
