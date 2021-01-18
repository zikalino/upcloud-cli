package server

import (
	"github.com/UpCloudLtd/cli/internal/commands"
	"github.com/UpCloudLtd/cli/internal/config"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestModifyCommand(t *testing.T) {
	methodName := "ModifyServer"

	var Server1 = upcloud.Server{
		CoreNumber:   1,
		Hostname:     "server-1-hostname",
		License:      0,
		MemoryAmount: 1024,
		Plan:         "server-1-plan",
		Progress:     0,
		State:        "started",
		Tags:         nil,
		Title:        "server-1-title",
		UUID:         "1fdfda29-ead1-4855-b71f-1e33eb2ca9de",
		Zone:         "fi-hel1",
	}

	var Server2 = upcloud.Server{
		CoreNumber:   1,
		Hostname:     "server-2-hostname",
		License:      0,
		MemoryAmount: 1024,
		Plan:         "server-2-plan",
		Progress:     0,
		State:        "started",
		Tags:         nil,
		Title:        "server-2-title",
		UUID:         "f77a5b25-84af-4f52-bc40-581930091fad",
		Zone:         "fi-hel1",
	}

	var servers = &upcloud.Servers{
		Servers: []upcloud.Server{
			Server1,
			Server2,
		},
	}

	details := upcloud.ServerDetails{
		Server: Server1,
	}

	for _, test := range []struct {
		name       string
		args       []string
		server     upcloud.Server
		modifyCall request.ModifyServerRequest
	}{
		{
			name: "Backend called, flags mapped to the correct field",
			args: []string{
				"--hostname", "example.com",
				"--title", "test-server",
				"--boot-order", "cdrom,network",
				"--cores", "12",
				"--memory", "4096",
				"--plan", "custom",
				"--simple-backup", "00,monthlies",
				"--time-zone", "EET",
				"--video-model", "VM",
				"--firewall", "false",
				"--metadata", "true",
				"--remote-access-enabled", "true",
				"--remote-access-type", upcloud.RemoteAccessTypeVNC,
				"--remote-access-password", "secret",
			},
			server: Server1,
			modifyCall: request.ModifyServerRequest{
				UUID:                 Server1.UUID,
				Hostname:             "example.com",
				Title:                "test-server",
				BootOrder:            "cdrom,network",
				CoreNumber:           12,
				MemoryAmount:         4096,
				Plan:                 "custom",
				SimpleBackup:         "00,monthlies",
				TimeZone:             "EET",
				VideoModel:           "VM",
				Firewall:             "off",
				Metadata:             upcloud.FromBool(true),
				RemoteAccessEnabled:  upcloud.FromBool(true),
				RemoteAccessType:     upcloud.RemoteAccessTypeVNC,
				RemoteAccessPassword: "secret",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			CachedServers = nil
			mss := MockServerService{}
			mss.On(methodName, &test.modifyCall).Return(&details, nil)
			mss.On("GetServers", mock.Anything).Return(servers, nil)
			mc := commands.BuildCommand(ModifyCommand(&mss), nil, config.New(viper.New()))
			mc.SetFlags(test.args)

			mc.MakeExecuteCommand()([]string{test.server.UUID})

			mss.AssertNumberOfCalls(t, methodName, 1)
		})
	}
}