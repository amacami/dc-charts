package pkg

import (
	"fmt"
	"io"
	"os"

	"github.com/amacami/dc-charts/internal/repository"
	"github.com/olekukonko/tablewriter"
	"github.com/prometheus/common/log"
)

//RepoClient
type RepoClient struct {
	repository *repository.Repository
}

//NewRepoClient returns new instance of the repository client
func NewRepoClient() *RepoClient {
	return &RepoClient{
		repository: repository.NewRepository("testdata/repo.json"),
	}
}

//ListRepoEntries prints a list of all charts from repository
func (c *RepoClient) ListRepoEntries(out io.Writer, filter string) error {
	list, err := c.repository.List(filter)
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(out)
	table.SetHeader([]string{"Name", "Version", "Description"})

	for _, v := range list {
		table.Append([]string{v.Name, v.Version, v.Description})
	}
	table.Render()
	return err
}

//Get downloads the DC files for the provided chart name
func (c *RepoClient) Get(out io.Writer, name string) error {
	chart, err := c.repository.Get(name)
	if err != nil {
		log.Error(err)
		return err
	}

	fileName := "testdata/" + chart.Url
	file, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return fmt.Errorf("cannot download chart from url %s", fileName)
	}

	fmt.Println("File to copy " + file.Name())
	return nil
}
