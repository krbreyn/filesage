package filesage

import "os"

type Watchlist struct {
	Dirs  []string `json:"dirs"`
	Files []string `json:"files"`
}

func (w *Watchlist) SaveToJson() {
}

func (w *Watchlist) LoadFromJson() {

}

func (w *Watchlist) AddWatchlistFile(filenames ...string) {
	w.Files = append(w.Files, filenames...)

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
			w.Dirs = append(w.Dirs, file.Name())
			w.handleRecursiveDir(file.Name())
		}
	}

	return nil
}
