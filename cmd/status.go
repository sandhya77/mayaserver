package cmd



import (

	

	"fmt"

	

	"github.com/mitchellh/cli"

	

)



type StatusCommand struct {

	Ui   cli.Ui

}



func (c *StatusCommand) Help() string {

	helpText := `

Usage: m-apiserver status



	Queries and Prints the current status ,

	port details and -bind ip address of  m-apiserver

`

	return (helpText)

}



func (c *StatusCommand) Run(_ []string) int {

// Initialise hard code stuff

	Status 		:= "running"

	NodeName	:= "m-apiserver"

	BindAddr	:= "172.28.128.6"

	Ports		:= "5656"





	//format the output

	out := fmt.Sprintf("Name          IP            Ports       Status       Data     Log \n%-14s%-14s%d\t%-14s",

		NodeName,

		BindAddr,

		Ports,

		Status)

	c.Ui.Output(out)



	return 0

}



func (c *StatusCommand) Synopsis() string {

	return "Prints the m-apiserver Status"

}
