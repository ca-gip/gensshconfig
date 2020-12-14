package services

import (
	"bytes"
	"fmt"
	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ca-gip/gensshconfig/internal/templates"
	"github.com/ca-gip/gensshconfig/internal/utils"
	"io/ioutil"
	"os"
	"text/template"
)

type SSHConfig struct {
	BasePah    string
	IgnoreDirs []string
	Bastion    Host
	Username   string
	Name       string
	Clusters   []*Cluster
	Template   string
}

func NewSSHConfig(basePath string, ignoreDirs []string, username string, configName string, bastion *Host) *SSHConfig {
	return &SSHConfig{
		BasePah:    basePath,
		IgnoreDirs: ignoreDirs,
		Name:       configName,
		Username:   username,
		Bastion:    *bastion,
	}
}

func (s *SSHConfig) FindCluster() (err error) {
	childs, err := s.readChildDir()

	if err != nil {
		return
	}

	clusters := utils.Difference(childs, s.IgnoreDirs)

	if len(clusters) == 0 {
		err = fmt.Errorf("could not find any child directory")
		return
	}

	for _, cluster := range clusters {
		s.Clusters = append(s.Clusters, s.newCluster(cluster))
	}

	return
}

func (s *SSHConfig) BuildClusterInventory() (err error) {
	for _, cluster := range s.Clusters {
		_ = cluster.readHostInCluster()
	}

	if hosts := s.hostTotal(); hosts == 0 {
		err = fmt.Errorf("could not find any hosts")
		return
	}

	return
}

func (s *SSHConfig) Render() {

	tpl, _ := template.New("config").Parse(templates.SshConfig)

	err := tpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}

}

// TODO : Filter out duplicates machine (like native-deploy) and move to a dedicated groups
func (s *SSHConfig) FilterDuplicate() {

}

func (s *SSHConfig) hostTotal() (result int64) {
	underscore.
		Chain(s.Clusters).
		Map(func(cluster Cluster, _ int) int64 { return int64(len(cluster.Hosts)) }).
		Aggregate(int64(0), func(acc int64, cur int64, _ int) int64 { return acc + cur }).
		Value(&result)
	return
}

func (s *SSHConfig) buildHostPath(dir string) string {
	var buffer bytes.Buffer

	buffer.WriteString(s.BasePah)
	buffer.WriteString("/")
	buffer.WriteString(dir)
	buffer.WriteString("/")
	buffer.WriteString(utils.HostFile)

	return buffer.String()
}

func (s *SSHConfig) newCluster(name string) *Cluster {
	return &Cluster{
		Name:             name,
		HostAbsolutePath: s.buildHostPath(name),
	}
}

func (s *SSHConfig) loadTemplate() {
	rawTemplate, _ := ioutil.ReadFile("./templates/ssh_config")
	s.Template = string(rawTemplate)
}

func (s *SSHConfig) readChildDir() (childs []string, err error) {
	files, err := ioutil.ReadDir(s.BasePah)

	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			childs = append(childs, file.Name())
		}
	}

	if len(childs) == 0 {
		err = fmt.Errorf("%s does not contains any diectories", s.BasePah)
		return
	}

	return
}
