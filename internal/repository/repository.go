package repository

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Repo struct {
	Charts []Chart `json:"charts"`
}
type Chart struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Repository struct {
	path string
}

func NewRepository(path string) *Repository {
	return &Repository{
		path: path,
	}
}

func init() {

}

func (r *Repository) List(filterString string) (ret []Chart, err error) {
	jsonFile, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var repo Repo
	json.Unmarshal(byteValue, &repo)

	for _, c := range repo.Charts {
		if strings.Contains(strings.ToLower(c.Name), strings.ToLower(filterString)) {
			ret = append(ret, c)
		}
	}

	return ret, err
}
