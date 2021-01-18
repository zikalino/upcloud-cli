package storage

import (
	"github.com/UpCloudLtd/cli/internal/commands"
	"github.com/UpCloudLtd/cli/internal/ui"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/service"
)

type restoreBackupCommand struct {
	*commands.BaseCommand
	service service.Storage
	params  restoreBackupParams
}

type restoreBackupParams struct {
	request.RestoreBackupRequest
}

func RestoreBackupCommand(service service.Storage) commands.Command {
	return &restoreBackupCommand{
		BaseCommand: commands.New("restore", "Restore backup of a storage"),
		service:     service,
	}
}

func (s *restoreBackupCommand) InitCommand() {
	s.SetPositionalArgHelp(positionalArgHelp)
	s.ArgCompletion(GetArgCompFn(s.service))
}

func (s *restoreBackupCommand) MakeExecuteCommand() func(args []string) (interface{}, error) {
	return func(args []string) (interface{}, error) {
		return Request{
			BuildRequest: func(uuid string) (interface{}, error) {
				req := s.params.RestoreBackupRequest
				req.UUID = uuid
				return &req, nil
			},
			Service: s.service,
			Handler: ui.HandleContext{
				RequestID:  func(in interface{}) string { return in.(*request.RestoreBackupRequest).UUID },
				MaxActions: maxStorageActions,
				ActionMsg:  "Restoring backup of storage",
				Action: func(req interface{}) (interface{}, error) {
					return nil, s.service.RestoreBackup(req.(*request.RestoreBackupRequest))
				},
			},
		}.Send(args)
	}
}