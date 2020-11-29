package url

import (
	"fmt"
	"net/url"
	"os"
)

type Download struct {
}

func NewDownloader() *Download {
	return &Download{}
}

//Get fetches the files from the provided url and copies locally
func (f *Download) Get(path string) error {
	u, _ := url.ParseRequestURI(path)

	if u.Scheme == "file" {
		return f.getFromLocalFS(u)
	}

	return nil
}

func (f *Download) getFromLocalFS(path *url.URL) error {
	fmt.Println("a" + path.Path + "a")
	file, err := os.Stat("./testdata/charts" + path.Path)
	if err != nil {
		return err
	}
	fmt.Println("File to copy " + file.Name())
	return nil
}
