/*

Creates a funKy project file and other boilerplate

*/

package project

import (
	"fmt"
)

const logo = `

      _/_/_/_/  _/    _/  _/      _/  _/    _/  _/      _/   
     _/        _/    _/  _/_/    _/  _/  _/      _/  _/      
    _/_/_/    _/    _/  _/  _/  _/  _/_/          _/         
   _/        _/    _/  _/    _/_/  _/  _/        _/          
  _/          _/_/    _/      _/  _/    _/      _/           

`

type Project struct {
	Name   string
	Prefix string
}

func NewProject(name string) *Project {
	return &Project{
		Name: name,
	}
}

func (p *Project) Boot() error {
	// dump header
	fmt.Println(logo)

	// if the project name is given

	// ask for the project name

	// create the function prefix based on the project name

	return nil
}
