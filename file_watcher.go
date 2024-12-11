package filesage

import (
	"os"
	"path/filepath"
)

type Watchlist struct {
	Dirs  []string `json:"dirs"`
	Files []string `json:"files"`
}

func (w *Watchlist) SaveToJson() {
}

func (w *Watchlist) LoadFromJson() {

}

func (w *Watchlist) AddWatchlistFile(filename string) {
	w.Files = append(w.Files, filename)

}

func (w *Watchlist) AddWatchlistDir(doRecursion bool, dirs ...string) error {
	for _, dir := range dirs {
		w.Dirs = append(w.Dirs, dir)

		if doRecursion {
			err := w.handleRecursiveDir(dir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *Watchlist) handleRecursiveDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			absFilePath := filepath.Join(dir, file.Name())
			if err != nil {
				return err
			}

			w.Dirs = append(w.Dirs, absFilePath)
			w.handleRecursiveDir(absFilePath)
		}
	}

	return nil
}
