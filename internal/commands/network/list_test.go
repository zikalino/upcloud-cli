package network

import (
	"github.com/UpCloudLtd/cli/internal/commands"
	"github.com/UpCloudLtd/cli/internal/config"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListCommand(t *testing.T) {

	var Network1 = upcloud.Network{
		Name: "network-1",
		UUID: "28e15cf5-8817-42ab-b017-970666be96ec",
		Type: upcloud.NetworkTypeUtility,
		Zone: "fi-hel1",
	}

	var Network2 = upcloud.Network{
		Name: "network-2",
		UUID: "f9f5ad16-a63a-4670-8449-c01d1e97281e",
		Type: upcloud.NetworkTypePrivate,
		Zone: "fi-hel1",
	}

	var Network3 = upcloud.Network{
		Name: "network-3",
		UUID: "e157ce0a-eeb0-49fc-9f2c-a05c3ac57066",
		Type: upcloud.NetworkTypeUtility,
		Zone: "uk-lon1",
	}

	var Network4 = upcloud.Network{
		Name: Network1.Name,
		UUID: "b3e49768-f13a-42c3-bea7-4e2471657f2f",
		Type: upcloud.NetworkTypePublic,
		Zone: "uk-lon1",
	}

	var networks = &upcloud.Networks{Networks: []upcloud.Network{Network1, Network2, Network3, Network4}}

	for _, test := range []struct {
		name     string
		flags    []string
		expected []upcloud.Network
	}{
		{
			name:     "get all",
			flags:    []string{},
			expected: []upcloud.Network{Network1, Network2, Network3, Network4},
		},
		{
			name:     "filter where type is utility",
			flags:    []string{"--type", "utility"},
			expected: []upcloud.Network{Network1, Network3},
		},
		{
			name:     "filter where zone is uk-lon1",
			flags:    []string{"--zone", "uk-lon1"},
			expected: []upcloud.Network{Network3, Network4},
		},
		{
			name:     "filter where zone is uk-lon1 and type is utility",
			flags:    []string{"--zone", "uk-lon1", "--type", "utility"},
			expected: []upcloud.Network{Network3},
		},
	} {
		cachedNetworks = nil
		mns := MockNetworkService{}
		mns.On("GetNetworks").Return(networks, nil)
		mns.On("GetNetworksInZone", &request.GetNetworksInZoneRequest{Zone: "fi-hel1"}).Return(&upcloud.Networks{Networks: []upcloud.Network{Network1, Network2}}, nil)
		mns.On("GetNetworksInZone", &request.GetNetworksInZoneRequest{Zone: "uk-lon1"}).Return(&upcloud.Networks{Networks: []upcloud.Network{Network3, Network4}}, nil)

		c := commands.BuildCommand(ListCommand(&mns), nil, config.New(viper.New()))
		c.SetFlags(test.flags)

		res, err := c.MakeExecuteCommand()([]string{})

		assert.Nil(t, err)
		assert.Equal(t, &upcloud.Networks{Networks: test.expected}, res)
	}
}