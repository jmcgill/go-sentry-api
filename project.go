package sentry

import (
	"fmt"
	"time"
)

// Asset is used from a plugin. Things like js/html
type Asset struct {
	URL string `json:"url,omitempty"`
}

// Plugin is a type of project plugin
type Plugin struct {
	Assets     []Asset                `json:"assets,omitempty"`
	IsTestable bool                   `json:"isTestable,omitempty"`
	Enabled    bool                   `json:"enabled,omitempty"`
	Name       string                 `json:"name,omitempty"`
	CanDisable bool                   `json:"canDisable,omitempty"`
	Type       string                 `json:"type,omitempty"`
	ID         string                 `json:"id,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

// Project is your project in sentry
type Project struct {
	Status             string                  `json:"status,omitempty"`
	Slug               *string                 `json:"slug,omitempty"`
	DefaultEnvironment *string                 `json:"defaultEnvironment,omitempty"`
	Color              *string                 `json:"color,omitempty"`
	IsPublic           bool                    `json:"isPublic,omitempty"`
	DateCreated        *time.Time              `json:"dateCreated,omitempty"`
	CallSign           string                  `json:"callSign,omitempty"`
	FirstEvent         *time.Time              `json:"firstEvent,omitempty"`
	IsBookmarked       bool                    `json:"isBookmarked,omitempty"`
	CallSignReviewed   bool                    `json:"callSignReviewed,omitempty"`
	ID                 string                  `json:"id,omitempty"`
	Name               string                  `json:"name"`
	Platforms          *[]string               `json:"platforms,omitempty"`
	Options            *map[string]interface{} `json:"options,omitempty"`
	Plugins            *[]Plugin               `json:"plugins,omitempty"`
	Team               *Team                   `json:"team,omitempty"`
	Organization       *Organization           `json:"organization,omitempty"`
	DigestMinDelay     *int                    `json:"digestMinDelay,omitempty"`
	DigestMaxDelay     *int                    `json:"digestMaxDelay,omitempty"`
}

type ProjectTag struct {
	Count     int       `json:"count"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	LastSeen  time.Time `json:"lastSeen"`
	Key       string    `json:"key"`
	FirstSeen time.Time `json:"firstSeen"`
}

// CreateProject will create a new project in your org and team
func (c *Client) CreateProject(o Organization, t Team, name string, slug *string) (Project, error) {
	var proj Project
	projreq := &Project{
		Name: name,
		Slug: slug,
	}

	err := c.do("POST", fmt.Sprintf("teams/%s/%s/projects", *o.Slug, *t.Slug), &proj, projreq)
	return proj, err
}

// GetProject takes a project slug and returns back the project
func (c *Client) GetProject(o Organization, projslug string) (Project, error) {
	var proj Project

	err := c.do("GET", fmt.Sprintf("projects/%s/%s", *o.Slug, projslug), &proj, nil)
	return proj, err
}

// UpdateProject takes a organization and project then updates it on the server side
func (c *Client) UpdateProject(o Organization, p Project) error {
	return c.do("PUT", fmt.Sprintf("projects/%s/%s", *o.Slug, *p.Slug), &p, &p)
}

// GetProjects fetchs all projects in a sentry instance
func (c *Client) GetProjects() ([]Project, error) {
	var proj []Project
	err := c.do("GET", "projects", &proj, nil)
	return proj, err
}

// GetTagValues fetchs all values for a tag for a given project
func (c *Client) GetProjectTagValues(o Organization, p Project, tag string) ([]ProjectTag, *Link, error) {
	var tags []ProjectTag
	link, err := c.doWithPagination("GET", fmt.Sprintf("projects/%s/%s/tags/%s/values", o.Slug, p.Slug, tag), &tags, nil)
	return tags, link, err
}

// DeleteProject will take your org, team, and proj and delete it from sentry.
func (c *Client) DeleteProject(o Organization, p Project) error {
	return c.do("DELETE", fmt.Sprintf("projects/%s/%s", *o.Slug, *p.Slug), nil, nil)
}
