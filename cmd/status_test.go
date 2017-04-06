package cmd



import (

	"github.com/mitchellh/cli"

	"testing"

	

)



func TestStatusCommand_implements(t *testing.T) {

	var _ cli.Command = &StatusCommand{}

}
