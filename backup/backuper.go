package backup

import (
	"github.com/pivotalservices/cfops/system"
)

type Backuper struct {
}

func New(factory system.CommandFactory, runner system.CommandRunner) Backuper {

	factory.Register("backup", BackupCommand{
		CommandRunner: runner,
	}).Register("restore", RestoreCommand{
		CommandRunner: runner,
	})

	return Backuper{}
}
