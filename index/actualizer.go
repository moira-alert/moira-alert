package index

import "time"

func (index *Index) actualizeIndex() error {
	ticker := time.NewTicker(time.Second)
	//actualizationTime := time.Now()
	for {
		select {
		case <-index.tomb.Dying():
			return nil
		case <-ticker.C:
			// do something
		}
	}
}