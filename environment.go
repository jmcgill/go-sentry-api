package sentry

import (
	"fmt"
)

// Project is your project in sentry
type Environment struct {
	Name             string                  `json:"name"`
	Id               *string                 `json:"id"`
	IsHidden		   bool					`json:"isHidden"`
}

// GetProjects fetchs all projects in a sentry instance
func (c *Client) GetEnvironments(o Organization, projslug string) ([]Environment, error) {
	var e []Environment
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/environments", *o.Slug, projslug), &e, nil)
	return e, err
}