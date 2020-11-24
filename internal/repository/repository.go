package repository

import (
	"encoding/json"
	"errors"
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

//NewRepository returns a new instance of chart repository
func NewRepository(path string) *Repository {
	return &Repository{
		path: path,
	}
}

//List returns a list of charts from repository
func (r *Repository) List(filterString string) (ret []Chart, err error) {
	repo, err := r.load()
	if err != nil {
		return nil, err
	}

	for _, c := range repo.Charts {
		if strings.Contains(strings.ToLower(c.Name), strings.ToLower(filterString)) {
			ret = append(ret, c)
		}
	}

	return ret, err
}

//Get returns a charts entity from repository, if exists
func (r *Repository) Get(name string) (ret *Chart, err error) {
	repo, err := r.load()
	if err != nil {
		return nil, err
	}

	for _, c := range repo.Charts {
		if strings.ToLower(c.Name) == strings.ToLower(name) {
			return &c, nil
		}
	}

	return nil, errors.New("no chart with name " + name + " found in repository")
}

func (r *Repository) load() (*Repo, error) {
	jsonFile, err := os.Open(r.path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var repo Repo
	json.Unmarshal(byteValue, &repo)
	return &repo, nil
}
