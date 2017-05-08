/*

Creates a funky project file and other boilerplate files.

Partially derived from the excellent Apex project...
* https://github.com/apex/apex/blob/master/boot/boot.go

*/

package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tj/go-prompt"
)

const logo = `

      _/_/_/_/  _/    _/  _/      _/  _/    _/  _/      _/   
     _/        _/    _/  _/_/    _/  _/  _/      _/  _/      
    _/_/_/    _/    _/  _/  _/  _/  _/_/          _/         
   _/        _/    _/  _/    _/_/  _/  _/        _/          
  _/          _/_/    _/      _/  _/    _/      _/           

`

const projectConfig = `---
project: %s
description: %s
memory: 128
timeout: 5
env: []
`

type Project struct {
	Name        string
	Description string
	Prefix      string
}

func NewProject(name string) *Project {
	return &Project{
		Name: name,
	}
}

func (p *Project) Boot() error {
	// dump header
	fmt.Println(logo)

	// if a project config file already exists
	if p.isConfig() {
		help("The './funky.yaml' project config already exists")
		return nil
	}

	// ask for the project name
	help("Enter the project name. This name will be used\nas a prefix to the cloud function names.\n")
	p.Name = prompt.StringRequired("  Project name: ")

	// create the function prefix based on the project name
	help(indent("Enter an optional project description"))
	p.Description = prompt.String(indent("  Project description: "))
	fmt.Println()

	// initialize the project files
	if err := p.initFiles(); err != nil {
		return err
	}

	return nil
}

func (p *Project) initFiles() error {
	fmt.Printf("creating ./funky.yaml")

	// create the project file
	project := fmt.Sprintf(projectConfig, p.Name, p.Description)
	if err := ioutil.WriteFile("funky.yaml", []byte(project), 644); err != nil {
		return err
	}

	// create the functions directory
	fpath := "functions"
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		os.Mkdir(fpath, 0755)
	}

	return nil
}

func (p *Project) isConfig() bool {
	_, err := os.Stat("funky.yaml")
	return err == nil
}

func help(s string) {
	os.Stdout.WriteString("\n")
	os.Stdout.WriteString(indent(s))
	os.Stdout.WriteString("\n")
}

// indent multiline strings
func indent(s string) (out string) {
	for _, l := range strings.SplitAfter(s, "\n") {
		out += fmt.Sprintf("  %s", l)
	}

	return
}
