/*

Creates a funKy project file and other boilerplate

Code reused and derived from the excellent Apex project:

	https://apex.run
	https://github.com/apex/apex/tree/master/boot

*/

package boot

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

type Bootstrapper struct {
	name string
}

func NewBootstrapper(name string) *Bootstrapper {
	return &Bootstrapper{
		name: name,
	}
}

func (b *Bootstrapper) Boot() error {
	fmt.Println(logo)
	return nil
}
