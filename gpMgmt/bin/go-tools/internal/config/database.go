package config

import (
	"github.com/greenplum-db/gpdb/gp/internal/enums"
)

type DatabaseConfig interface {
	GetDeploymentType() enums.DeploymentType
	GetAdmin() UserConfig
	GetSegmentsPerSegmentHost() int
}
type Database struct {
	DeploymentType         enums.DeploymentType `json:"deploymentType"`
	Admin                  *User                `json:"admin"`
	SegmentsPerSegmentHost int                  `json:"segmentsPerSegmentHost"`
}

func (d Database) GetDeploymentType() enums.DeploymentType {
	return d.DeploymentType
}
func (d Database) GetAdmin() UserConfig {
	return d.Admin
}
func (d Database) GetSegmentsPerSegmentHost() int {
	return d.SegmentsPerSegmentHost
}

type UserConfig interface {
	GetName() string
	GetPassword() string
}
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u User) GetName() string {
	return u.Name
}
func (u User) GetPassword() string {
	return u.Password
}

type ArtifactConfig interface {
	GetGreenplum() string
	GetDependencyList() []string
}
type Artifact struct {
	Greenplum      string   `json:"greenplum"`
	DependencyList []string `json:"dependencyList"`
}

func (a Artifact) GetGreenplum() string {
	return a.Greenplum
}
func (a Artifact) GetDependencyList() []string {
	return a.DependencyList
}
