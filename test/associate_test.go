package test

import (
	"fmt"
	"testing"

	"github.com/benny502/gorm-helper/associate"
	"github.com/stretchr/testify/assert"
)

func TestAssociate(t *testing.T) {
	ass := associate.NewAssociate(&User{}, "TeamUser.Team")
	joins := ass.GetJoinsString()
	assert.NotEmpty(t, joins)
	preload := ass.GetPreload()
	assert.NotEmpty(t, preload)
	fmt.Printf("%s\n%v", preload, joins)
	//panic("doom")
}

func TestAssociateSlice(t *testing.T) {
	ass := associate.NewAssociate(&Team{}, "TeamUser")
	joins := ass.GetJoinsString()
	assert.NotEmpty(t, joins)
	preload := ass.GetPreload()
	assert.NotEmpty(t, preload)
	fmt.Printf("%s\n%v", preload, joins)
	//panic("ok")
}
