package main

import (
	"encoding/json"
)

type ProjectStatus int

const (
	ActiveProject ProjectStatus = iota
	TrashedProject
)

func (s ProjectStatus) String() string {
	return [...]string{"active", "trashed"}[s]
}

func (s *ProjectStatus) FromString(status string) ProjectStatus {
	return map[string]ProjectStatus{
		"active":  ActiveProject,
		"trashed": TrashedProject,
	}[status]
}

func (s ProjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (p *ProjectStatus) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*p = p.FromString(s)
	return nil
}

type Project struct {
	Id          int
	Status      ProjectStatus
	Name        string
	Description string
	Dock        []Dock
}

type Dock struct {
	Title   string
	Enabled bool
	Id      int
}
