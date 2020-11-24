package pkg

import (
	"io"

	"github.com/amacami/dc-charts/internal/repository"
	"github.com/olekukonko/tablewriter"
	"github.com/prometheus/common/log"
)

type RepoClient struct {
	repository *repository.Repository
}

func NewRepoClient() *RepoClient {
	return &RepoClient{
		repository: repository.NewRepository("repository/repo.json"),
	}
}
func (c *RepoClient) ListRepoEntries(out io.Writer, filter string) error {
	list, err := c.repository.List(filter)
	if err != nil {
		log.Error(err)
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
