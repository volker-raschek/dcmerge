package dockerCompose_test

import (
	"testing"

	"git.cryptic.systems/volker.raschek/dcmerge/pkg/domain/dockerCompose"
	"github.com/stretchr/testify/require"
)

func TestNetwork_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA: &dockerCompose.Network{
				External: true,
			},
			equalableB:     &dockerCompose.NetworkIPAM{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Network{
				External: true,
			},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Network{
				External: false,
				Driver:   "bridge",
				IPAM:     nil,
			},
			equalableB: &dockerCompose.Network{
				External: false,
				Driver:   "bridge",
				IPAM:     nil,
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Network{
				External: false,
				Driver:   "host",
				IPAM:     nil,
			},
			equalableB: &dockerCompose.Network{
				External: false,
				Driver:   "bride",
				IPAM:     nil,
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Network{
				External: true,
				Driver:   "bridge",
				IPAM:     nil,
			},
			equalableB: &dockerCompose.Network{
				External: false,
				Driver:   "bridge",
				IPAM:     nil,
			},
			expectedResult: false,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestNetworkIPAM_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.NetworkIPAM{},
			equalableB:     &dockerCompose.Service{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.NetworkIPAM{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.NetworkIPAM{
				Configs: make([]*dockerCompose.NetworkIPAMConfig, 0),
			},
			equalableB:     &dockerCompose.NetworkIPAM{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.NetworkIPAM{
				Configs: make([]*dockerCompose.NetworkIPAMConfig, 0),
			},
			equalableB: &dockerCompose.NetworkIPAM{
				Configs: make([]*dockerCompose.NetworkIPAMConfig, 0),
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestNetworkIPAMConfig_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.NetworkIPAMConfig{},
			equalableB:     &dockerCompose.Service{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.NetworkIPAMConfig{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.NetworkIPAMConfig{
				Subnet: "10.12.13.14/15",
			},
			equalableB:     &dockerCompose.NetworkIPAMConfig{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.NetworkIPAMConfig{
				Subnet: "10.12.13.14/15",
			},
			equalableB: &dockerCompose.NetworkIPAMConfig{
				Subnet: "10.12.13.14/15",
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestSecret_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.Secret{},
			equalableB:     &dockerCompose.Service{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.Secret{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Secret{
				File: "/var/run/docker/app/secret",
			},
			equalableB:     &dockerCompose.Secret{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Secret{
				File: "/var/run/docker/app/secret",
			},
			equalableB: &dockerCompose.Secret{
				File: "/var/run/docker/app/secret",
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestService_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.Service{},
			equalableB:     &dockerCompose.Secret{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.Service{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.Service{},
			equalableB:     &dockerCompose.Service{},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				CapabilitiesAdd:  []string{},
				CapabilitiesDrop: []string{},
				Deploy:           nil,
				Environments:     []string{},
				ExtraHosts:       []string{},
				Image:            "",
				Labels:           []string{},
				Networks:         map[string]*dockerCompose.ServiceNetwork{},
				Ports:            []string{},
				Secrets:          []string{},
				ULimits:          nil,
				Volumes:          []string{},
			},
			equalableB: &dockerCompose.Service{
				CapabilitiesAdd:  []string{},
				CapabilitiesDrop: []string{},
				Deploy:           nil,
				Environments:     []string{},
				ExtraHosts:       []string{},
				Image:            "",
				Labels:           []string{},
				Networks:         map[string]*dockerCompose.ServiceNetwork{},
				Ports:            []string{},
				Secrets:          []string{},
				ULimits:          nil,
				Volumes:          []string{},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_ADMIN"},
			},
			equalableB: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_ADMIN"},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_ADMIN"},
			},
			equalableB: &dockerCompose.Service{
				CapabilitiesAdd: []string{},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_ADMIN"},
			},
			equalableB: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_ADMIN"},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_ADMIN"},
			},
			equalableB: &dockerCompose.Service{
				CapabilitiesDrop: []string{},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Deploy: &dockerCompose.ServiceDeploy{},
			},
			equalableB: &dockerCompose.Service{
				Deploy: nil,
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=localhost.localdomain"},
			},
			equalableB: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=localhost.localdomain"},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=localhost.localdomain"},
			},
			equalableB: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=localhost"},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=localhost.localdomain"},
			},
			equalableB: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=localdomain.localhost"},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				ExtraHosts: []string{"my-app.u.orbis-healthcare.com"},
			},
			equalableB: &dockerCompose.Service{
				ExtraHosts: []string{"my-app.u.orbis-healthcare.com"},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				ExtraHosts: []string{"my-app.u.orbis-healthcare.com"},
			},
			equalableB: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Image: "registry.example.local/my/app:latest",
			},
			equalableB: &dockerCompose.Service{
				Image: "registry.example.local/my/app:latest",
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				Image: "registry.example.local/my/app:latest",
			},
			equalableB: &dockerCompose.Service{
				Image: "",
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Labels: []string{"keyA=valueA"},
			},
			equalableB: &dockerCompose.Service{
				Labels: []string{"keyA=valueA"},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				Labels: []string{"keyA=valueA", "keyA=valueB"},
			},
			equalableB: &dockerCompose.Service{
				Labels: []string{"keyA=valueA"},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
			equalableB: &dockerCompose.Service{
				Networks: nil,
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
			equalableB: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				Ports: []string{"80:80/tcp"},
			},
			equalableB: &dockerCompose.Service{
				Ports: []string{"80:80/tcp"},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				Ports: []string{"80:80/tcp"},
			},
			equalableB: &dockerCompose.Service{
				Ports: []string{"80:80/udp"},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Secrets: make([]string, 0),
			},
			equalableB: &dockerCompose.Service{
				Secrets: make([]string, 0),
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
			equalableB:     &dockerCompose.Service{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
			equalableB: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Service{
				Volumes: []string{"/var/run/docker/volume/mountA"},
			},
			equalableB: &dockerCompose.Service{
				Volumes: []string{"/var/run/docker/volume/mountB"},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Service{
				Volumes: []string{"/var/run/docker/volume/mountA"},
			},
			equalableB: &dockerCompose.Service{
				Volumes: []string{"/var/run/docker/volume/mountA"},
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestService_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		serviceDeploymentA *dockerCompose.Service
		serviceDeploymentB *dockerCompose.Service
		expectedService    *dockerCompose.Service
	}{
		{
			serviceDeploymentA: nil,
			serviceDeploymentB: nil,
			expectedService:    nil,
		},
		{
			serviceDeploymentA: &dockerCompose.Service{},
			serviceDeploymentB: &dockerCompose.Service{},
			expectedService:    &dockerCompose.Service{},
		},

		// CapabilitiesAdd
		{
			serviceDeploymentA: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_RAW"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				CapabilitiesAdd: []string{},
			},
			expectedService: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_RAW"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				CapabilitiesAdd: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_RAW"},
			},
			expectedService: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_RAW"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_RAW"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_RAW"},
			},
			expectedService: &dockerCompose.Service{
				CapabilitiesAdd: []string{"NET_RAW"},
			},
		},

		// CapabilitiesDrop
		{
			serviceDeploymentA: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_RAW"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				CapabilitiesDrop: []string{},
			},
			expectedService: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_RAW"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				CapabilitiesDrop: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_RAW"},
			},
			expectedService: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_RAW"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_RAW"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_RAW"},
			},
			expectedService: &dockerCompose.Service{
				CapabilitiesDrop: []string{"NET_RAW"},
			},
		},

		// Deploy
		{
			serviceDeploymentA: &dockerCompose.Service{
				Deploy: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Deploy: nil,
			},
			expectedService: &dockerCompose.Service{
				Deploy: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Deploy: dockerCompose.NewServiceDeploy(),
			},
			serviceDeploymentB: &dockerCompose.Service{
				Deploy: nil,
			},
			expectedService: &dockerCompose.Service{
				Deploy: dockerCompose.NewServiceDeploy(),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Deploy: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Deploy: dockerCompose.NewServiceDeploy(),
			},
			expectedService: &dockerCompose.Service{
				Deploy: dockerCompose.NewServiceDeploy(),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Deploy: dockerCompose.NewServiceDeploy(),
			},
			serviceDeploymentB: &dockerCompose.Service{
				Deploy: dockerCompose.NewServiceDeploy(),
			},
			expectedService: &dockerCompose.Service{
				Deploy: dockerCompose.NewServiceDeploy(),
			},
		},

		// Environments
		{
			serviceDeploymentA: &dockerCompose.Service{
				Environments: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Environments: nil,
			},
			expectedService: &dockerCompose.Service{
				Environments: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Environments: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Environments: nil,
			},
			expectedService: &dockerCompose.Service{
				Environments: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Environments: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Environments: []string{},
			},
			expectedService: &dockerCompose.Service{
				Environments: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Environments: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Environments: []string{},
			},
			expectedService: &dockerCompose.Service{
				Environments: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=u.example.com"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=u.example.local"},
			},
			expectedService: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=u.example.local"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=u.example.com", "PROXY_HOST=u.example.de"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=u.example.local"},
			},
			expectedService: &dockerCompose.Service{
				Environments: []string{"PROXY_HOST=u.example.local"},
			},
		},

		// ExtraHosts
		{
			serviceDeploymentA: &dockerCompose.Service{
				ExtraHosts: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				ExtraHosts: nil,
			},
			expectedService: &dockerCompose.Service{
				ExtraHosts: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				ExtraHosts: nil,
			},
			expectedService: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ExtraHosts: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
			expectedService: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
			expectedService: &dockerCompose.Service{
				ExtraHosts: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ExtraHosts: []string{"extra.host.local"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				ExtraHosts: []string{"extra.host.local"},
			},
			expectedService: &dockerCompose.Service{
				ExtraHosts: []string{"extra.host.local"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ExtraHosts: []string{"extra.host.local"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				ExtraHosts: []string{"extra.host.com"},
			},
			expectedService: &dockerCompose.Service{
				ExtraHosts: []string{"extra.host.com", "extra.host.local"},
			},
		},

		// Labels
		{
			serviceDeploymentA: &dockerCompose.Service{
				Labels: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Labels: nil,
			},
			expectedService: &dockerCompose.Service{
				Labels: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Labels: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Labels: nil,
			},
			expectedService: &dockerCompose.Service{
				Labels: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Labels: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Labels: []string{},
			},
			expectedService: &dockerCompose.Service{
				Labels: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Labels: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Labels: []string{},
			},
			expectedService: &dockerCompose.Service{
				Labels: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Labels: []string{"prometheus.io/scrape=true"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Labels: []string{"prometheus.io/scrape=true"},
			},
			expectedService: &dockerCompose.Service{
				Labels: []string{"prometheus.io/scrape=true"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Labels: []string{"prometheus.io/scrape=true", "prometheus.io/scrape=false"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Labels: []string{"prometheus.io/scrape=true"},
			},
			expectedService: &dockerCompose.Service{
				Labels: []string{"prometheus.io/scrape=true"},
			},
		},

		// Networks
		{
			serviceDeploymentA: &dockerCompose.Service{
				Networks: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Networks: nil,
			},
			expectedService: &dockerCompose.Service{
				Networks: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
			serviceDeploymentB: &dockerCompose.Service{
				Networks: nil,
			},
			expectedService: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Networks: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
			expectedService: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
			serviceDeploymentB: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
			expectedService: &dockerCompose.Service{
				Networks: make(map[string]*dockerCompose.ServiceNetwork),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Networks: map[string]*dockerCompose.ServiceNetwork{
					"proxy": {Aliases: []string{"app.proxy.network"}},
				},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Networks: nil,
			},
			expectedService: &dockerCompose.Service{
				Networks: map[string]*dockerCompose.ServiceNetwork{
					"proxy": {Aliases: []string{"app.proxy.network"}},
				},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Networks: map[string]*dockerCompose.ServiceNetwork{
					"proxy": {Aliases: []string{"app.proxy.network"}},
				},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Networks: map[string]*dockerCompose.ServiceNetwork{
					"db": {Aliases: []string{"app.db.network"}},
				},
			},
			expectedService: &dockerCompose.Service{
				Networks: map[string]*dockerCompose.ServiceNetwork{
					"db":    {Aliases: []string{"app.db.network"}},
					"proxy": {Aliases: []string{"app.proxy.network"}},
				},
			},
		},

		// Ports
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: nil,
			},
			expectedService: &dockerCompose.Service{
				Ports: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: nil,
			},
			expectedService: &dockerCompose.Service{
				Ports: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: []string{},
			},
			expectedService: &dockerCompose.Service{
				Ports: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: []string{},
			},
			expectedService: &dockerCompose.Service{
				Ports: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: []string{"80:80"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: []string{"80:80"},
			},
			expectedService: &dockerCompose.Service{
				Ports: []string{"80:80"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: []string{"80:80"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: []string{"10080:80"},
			},
			expectedService: &dockerCompose.Service{
				Ports: []string{"10080:80"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: []string{"80:80/tcp"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: []string{"80:80"},
			},
			expectedService: &dockerCompose.Service{
				Ports: []string{"80:80"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Ports: []string{"80:80"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Ports: []string{"10080:80/udp"},
			},
			expectedService: &dockerCompose.Service{
				Ports: []string{"10080:80/udp"},
			},
		},

		// Secrets
		{
			serviceDeploymentA: &dockerCompose.Service{
				Secrets: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Secrets: nil,
			},
			expectedService: &dockerCompose.Service{
				Secrets: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Secrets: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Secrets: nil,
			},
			expectedService: &dockerCompose.Service{
				Secrets: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Secrets: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Secrets: []string{},
			},
			expectedService: &dockerCompose.Service{
				Secrets: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Secrets: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Secrets: []string{},
			},
			expectedService: &dockerCompose.Service{
				Secrets: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Secrets: []string{"db_pass_credentials"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Secrets: []string{"db_pass_credentials"},
			},
			expectedService: &dockerCompose.Service{
				Secrets: []string{"db_pass_credentials"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Secrets: []string{"db_pass_credentials"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Secrets: []string{"oauth2_pass_credentials"},
			},
			expectedService: &dockerCompose.Service{
				Secrets: []string{"db_pass_credentials", "oauth2_pass_credentials"},
			},
		},

		// ULimits
		{
			serviceDeploymentA: &dockerCompose.Service{
				ULimits: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				ULimits: nil,
			},
			expectedService: &dockerCompose.Service{
				ULimits: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
			serviceDeploymentB: &dockerCompose.Service{
				ULimits: nil,
			},
			expectedService: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ULimits: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
			expectedService: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
			serviceDeploymentB: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
			expectedService: &dockerCompose.Service{
				ULimits: dockerCompose.NewServiceULimits(),
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ULimits: &dockerCompose.ServiceULimits{
					NProc: 10,
					NoFile: &dockerCompose.ServiceULimitsNoFile{
						Hard: 10,
						Soft: 10,
					},
				},
			},
			serviceDeploymentB: &dockerCompose.Service{
				ULimits: &dockerCompose.ServiceULimits{
					NProc: 10,
					NoFile: &dockerCompose.ServiceULimitsNoFile{
						Hard: 10,
						Soft: 10,
					},
				},
			},
			expectedService: &dockerCompose.Service{
				ULimits: &dockerCompose.ServiceULimits{
					NProc: 10,
					NoFile: &dockerCompose.ServiceULimitsNoFile{
						Hard: 10,
						Soft: 10,
					},
				},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				ULimits: &dockerCompose.ServiceULimits{
					NProc: 10,
					NoFile: &dockerCompose.ServiceULimitsNoFile{
						Hard: 10,
						Soft: 10,
					},
				},
			},
			serviceDeploymentB: &dockerCompose.Service{
				ULimits: &dockerCompose.ServiceULimits{
					NProc: 15,
					NoFile: &dockerCompose.ServiceULimitsNoFile{
						Hard: 25,
						Soft: 20,
					},
				},
			},
			expectedService: &dockerCompose.Service{
				ULimits: &dockerCompose.ServiceULimits{
					NProc: 15,
					NoFile: &dockerCompose.ServiceULimitsNoFile{
						Hard: 25,
						Soft: 20,
					},
				},
			},
		},

		// Volumes
		{
			serviceDeploymentA: &dockerCompose.Service{
				Volumes: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Volumes: nil,
			},
			expectedService: &dockerCompose.Service{
				Volumes: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Volumes: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Volumes: nil,
			},
			expectedService: &dockerCompose.Service{
				Volumes: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Volumes: nil,
			},
			serviceDeploymentB: &dockerCompose.Service{
				Volumes: []string{},
			},
			expectedService: &dockerCompose.Service{
				Volumes: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Volumes: []string{},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Volumes: []string{},
			},
			expectedService: &dockerCompose.Service{
				Volumes: []string{},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Volumes: []string{"/etc/localtime:/etc/localtime"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Volumes: []string{"/etc/localtime:/etc/localtime"},
			},
			expectedService: &dockerCompose.Service{
				Volumes: []string{"/etc/localtime:/etc/localtime"},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.Service{
				Volumes: []string{"/etc/localtime:/etc/localtime"},
			},
			serviceDeploymentB: &dockerCompose.Service{
				Volumes: []string{"/usr/share/zoneinfo/Europe/Berlin:/etc/localtime"},
			},
			expectedService: &dockerCompose.Service{
				Volumes: []string{"/usr/share/zoneinfo/Europe/Berlin:/etc/localtime"},
			},
		},
	}

	for i, testCase := range testCases {
		testCase.serviceDeploymentA.MergeLastWin(testCase.serviceDeploymentB)
		require.True(testCase.expectedService.Equal(testCase.serviceDeploymentA), "Failed test case %v", i)
	}
}

func TestSecretDeploy_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.ServiceDeploy{},
			equalableB:     &dockerCompose.Service{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.ServiceDeploy{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeploy{
				Resources: dockerCompose.NewServiceDeployResources(),
			},
			equalableB:     &dockerCompose.ServiceDeploy{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeploy{
				Resources: dockerCompose.NewServiceDeployResources(),
			},
			equalableB: &dockerCompose.ServiceDeploy{
				Resources: dockerCompose.NewServiceDeployResources(),
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestServiceDeploy_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		serviceDeploymentA        *dockerCompose.ServiceDeploy
		serviceDeploymentB        *dockerCompose.ServiceDeploy
		expectedServiceDeployment *dockerCompose.ServiceDeploy
	}{
		{
			serviceDeploymentA:        nil,
			serviceDeploymentB:        nil,
			expectedServiceDeployment: nil,
		},
		{
			serviceDeploymentA: &dockerCompose.ServiceDeploy{
				Resources: nil,
			},
			serviceDeploymentB: &dockerCompose.ServiceDeploy{
				Resources: nil,
			},
			expectedServiceDeployment: &dockerCompose.ServiceDeploy{
				Resources: nil,
			},
		},
		{
			serviceDeploymentA: &dockerCompose.ServiceDeploy{
				Resources: &dockerCompose.ServiceDeployResources{
					Limits: nil,
				},
			},
			serviceDeploymentB: &dockerCompose.ServiceDeploy{
				Resources: nil,
			},
			expectedServiceDeployment: &dockerCompose.ServiceDeploy{
				Resources: &dockerCompose.ServiceDeployResources{
					Limits: nil,
				},
			},
		},
		{
			serviceDeploymentA: &dockerCompose.ServiceDeploy{
				Resources: nil,
			},
			serviceDeploymentB: &dockerCompose.ServiceDeploy{
				Resources: &dockerCompose.ServiceDeployResources{
					Limits: nil,
				},
			},
			expectedServiceDeployment: &dockerCompose.ServiceDeploy{
				Resources: &dockerCompose.ServiceDeployResources{
					Limits: nil,
				},
			},
		},
	}

	for i, testCase := range testCases {
		testCase.serviceDeploymentA.MergeLastWin(testCase.serviceDeploymentB)
		require.True(testCase.expectedServiceDeployment.Equal(testCase.serviceDeploymentA), "Failed test case %v", i)
	}
}

func TestSecretDeployResources_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.ServiceDeployResources{},
			equalableB:     &dockerCompose.Service{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.ServiceDeployResources{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResources{
				Limits: dockerCompose.NewServiceDeployResourcesLimits(),
			},
			equalableB:     &dockerCompose.ServiceDeployResources{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResources{
				Limits: dockerCompose.NewServiceDeployResourcesLimits(),
			},
			equalableB: &dockerCompose.ServiceDeployResources{
				Limits: dockerCompose.NewServiceDeployResourcesLimits(),
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResources{
				Reservations: dockerCompose.NewServiceDeployResourcesLimits(),
			},
			equalableB:     &dockerCompose.ServiceDeployResources{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResources{
				Reservations: dockerCompose.NewServiceDeployResourcesLimits(),
			},
			equalableB: &dockerCompose.ServiceDeployResources{
				Reservations: dockerCompose.NewServiceDeployResourcesLimits(),
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestServiceDeployResources_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		serviceDeploymentResourcesA        *dockerCompose.ServiceDeployResources
		serviceDeploymentResourcesB        *dockerCompose.ServiceDeployResources
		expectedServiceDeploymentResources *dockerCompose.ServiceDeployResources
	}{
		{
			serviceDeploymentResourcesA:        nil,
			serviceDeploymentResourcesB:        nil,
			expectedServiceDeploymentResources: nil,
		},
		{
			serviceDeploymentResourcesA: &dockerCompose.ServiceDeployResources{
				Limits: nil,
			},
			serviceDeploymentResourcesB: &dockerCompose.ServiceDeployResources{
				Limits: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
			expectedServiceDeploymentResources: &dockerCompose.ServiceDeployResources{
				Limits: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
		},
		{
			serviceDeploymentResourcesA: &dockerCompose.ServiceDeployResources{
				Limits: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
			serviceDeploymentResourcesB: &dockerCompose.ServiceDeployResources{
				Limits: nil,
			},
			expectedServiceDeploymentResources: &dockerCompose.ServiceDeployResources{
				Limits: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
		},
		{
			serviceDeploymentResourcesA: &dockerCompose.ServiceDeployResources{
				Limits: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
			serviceDeploymentResourcesB: &dockerCompose.ServiceDeployResources{
				Limits: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "2",
					Memory: "1000",
				},
			},
			expectedServiceDeploymentResources: &dockerCompose.ServiceDeployResources{
				Limits: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "2",
					Memory: "1000",
				},
			},
		},
		{
			serviceDeploymentResourcesA: &dockerCompose.ServiceDeployResources{
				Reservations: nil,
			},
			serviceDeploymentResourcesB: &dockerCompose.ServiceDeployResources{
				Reservations: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
			expectedServiceDeploymentResources: &dockerCompose.ServiceDeployResources{
				Reservations: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
		},
		{
			serviceDeploymentResourcesA: &dockerCompose.ServiceDeployResources{
				Reservations: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
			serviceDeploymentResourcesB: &dockerCompose.ServiceDeployResources{
				Reservations: nil,
			},
			expectedServiceDeploymentResources: &dockerCompose.ServiceDeployResources{
				Reservations: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
		},
		{
			serviceDeploymentResourcesA: &dockerCompose.ServiceDeployResources{
				Reservations: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "1",
					Memory: "500",
				},
			},
			serviceDeploymentResourcesB: &dockerCompose.ServiceDeployResources{
				Reservations: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "2",
					Memory: "1000",
				},
			},
			expectedServiceDeploymentResources: &dockerCompose.ServiceDeployResources{
				Reservations: &dockerCompose.ServiceDeployResourcesLimits{
					CPUs:   "2",
					Memory: "1000",
				},
			},
		},
	}

	for i, testCase := range testCases {
		testCase.serviceDeploymentResourcesA.MergeLastWin(testCase.serviceDeploymentResourcesB)
		require.True(testCase.expectedServiceDeploymentResources.Equal(testCase.serviceDeploymentResourcesA), "Failed test case %v", i)
	}
}

func TestServiceDeployResourcesLimits_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
			equalableB:     &dockerCompose.NetworkIPAM{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
			equalableB: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs: "1",
			},
			equalableB: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs: "2",
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceDeployResourcesLimits{
				Memory: "500",
			},
			equalableB: &dockerCompose.ServiceDeployResourcesLimits{
				Memory: "1000",
			},
			expectedResult: false,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestServiceDeployResourcesLimits_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		serviceDeploymentResourcesLimitsA        *dockerCompose.ServiceDeployResourcesLimits
		serviceDeploymentResourcesLimitsB        *dockerCompose.ServiceDeployResourcesLimits
		expectedServiceDeploymentResourcesLimits *dockerCompose.ServiceDeployResourcesLimits
	}{
		{
			serviceDeploymentResourcesLimitsA:        nil,
			serviceDeploymentResourcesLimitsB:        nil,
			expectedServiceDeploymentResourcesLimits: nil,
		},
		{
			serviceDeploymentResourcesLimitsA: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
			serviceDeploymentResourcesLimitsB: nil,
			expectedServiceDeploymentResourcesLimits: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
		},
		{
			serviceDeploymentResourcesLimitsA: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
			serviceDeploymentResourcesLimitsB: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
			expectedServiceDeploymentResourcesLimits: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs:   "1",
				Memory: "500",
			},
		},
		{
			serviceDeploymentResourcesLimitsA: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs: "1",
			},
			serviceDeploymentResourcesLimitsB: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs: "2",
			},
			expectedServiceDeploymentResourcesLimits: &dockerCompose.ServiceDeployResourcesLimits{
				CPUs: "2",
			},
		},
		{
			serviceDeploymentResourcesLimitsA: &dockerCompose.ServiceDeployResourcesLimits{
				Memory: "500",
			},
			serviceDeploymentResourcesLimitsB: &dockerCompose.ServiceDeployResourcesLimits{
				Memory: "1000",
			},
			expectedServiceDeploymentResourcesLimits: &dockerCompose.ServiceDeployResourcesLimits{
				Memory: "1000",
			},
		},
	}

	for i, testCase := range testCases {
		testCase.serviceDeploymentResourcesLimitsA.MergeLastWin(testCase.serviceDeploymentResourcesLimitsB)
		require.True(testCase.expectedServiceDeploymentResourcesLimits.Equal(testCase.serviceDeploymentResourcesLimitsA), "Failed test case %v", i)
	}
}

func TestServiceNetwork_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA: &dockerCompose.ServiceNetwork{
				Aliases: []string{},
			},
			equalableB:     &dockerCompose.NetworkIPAM{},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceNetwork{
				Aliases: []string{},
			},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceNetwork{
				Aliases: []string{},
			},
			equalableB: &dockerCompose.ServiceNetwork{
				Aliases: []string{},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.ServiceNetwork{
				Aliases: []string{"HelloWorld"},
			},
			equalableB: &dockerCompose.ServiceNetwork{
				Aliases: []string{"HelloWorld"},
			},
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.ServiceNetwork{
				Aliases: []string{"HelloWorld"},
			},
			equalableB: &dockerCompose.ServiceNetwork{
				Aliases: []string{"FooBar"},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceNetwork{
				Aliases: []string{"Hello", "World"},
			},
			equalableB: &dockerCompose.ServiceNetwork{
				Aliases: []string{"FooBar"},
			},
			expectedResult: false,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestServiceNetwork_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		ServiceNetworkA        *dockerCompose.ServiceNetwork
		ServiceNetworkB        *dockerCompose.ServiceNetwork
		expectedServiceNetwork *dockerCompose.ServiceNetwork
	}{
		{
			ServiceNetworkA:        nil,
			ServiceNetworkB:        nil,
			expectedServiceNetwork: nil,
		},
		{
			ServiceNetworkA:        &dockerCompose.ServiceNetwork{},
			ServiceNetworkB:        nil,
			expectedServiceNetwork: &dockerCompose.ServiceNetwork{},
		},
		{
			ServiceNetworkA: &dockerCompose.ServiceNetwork{
				Aliases: []string{"my-app.example.com"},
			},
			ServiceNetworkB: &dockerCompose.ServiceNetwork{
				Aliases: []string{"my-app.example.com"},
			},
			expectedServiceNetwork: &dockerCompose.ServiceNetwork{
				Aliases: []string{"my-app.example.com"},
			},
		},
		{
			ServiceNetworkA: &dockerCompose.ServiceNetwork{
				Aliases: []string{"my-app.example.com"},
			},
			ServiceNetworkB: &dockerCompose.ServiceNetwork{
				Aliases: []string{"my-app.example.local"},
			},
			expectedServiceNetwork: &dockerCompose.ServiceNetwork{
				Aliases: []string{"my-app.example.com", "my-app.example.local"},
			},
		},
	}

	for i, testCase := range testCases {
		testCase.ServiceNetworkA.MergeLastWin(testCase.ServiceNetworkB)
		require.True(testCase.expectedServiceNetwork.Equal(testCase.ServiceNetworkA), "Failed test case %v", i)
	}
}

func TestServiceULimits_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.ServiceULimits{},
			equalableB:     &dockerCompose.NetworkIPAM{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.ServiceULimits{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceULimits{
				NProc:  0,
				NoFile: dockerCompose.NewServiceULimitsNoFile(),
			},
			equalableB: &dockerCompose.ServiceULimits{
				NProc: 0,
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceULimits{
				NProc: 0,
				NoFile: &dockerCompose.ServiceULimitsNoFile{
					Hard: 10,
				},
			},
			equalableB: &dockerCompose.ServiceULimits{
				NProc: 0,
				NoFile: &dockerCompose.ServiceULimitsNoFile{
					Soft: 10,
				},
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceULimits{
				NProc: 20,
				NoFile: &dockerCompose.ServiceULimitsNoFile{
					Hard: 10,
					Soft: 10,
				},
			},
			equalableB: &dockerCompose.ServiceULimits{
				NProc: 20,
				NoFile: &dockerCompose.ServiceULimitsNoFile{
					Hard: 10,
					Soft: 10,
				},
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestServiceULimits_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		ServiceULimitsA        *dockerCompose.ServiceULimits
		ServiceULimitsB        *dockerCompose.ServiceULimits
		expectedServiceULimits *dockerCompose.ServiceULimits
	}{
		{
			ServiceULimitsA:        nil,
			ServiceULimitsB:        nil,
			expectedServiceULimits: nil,
		},
		{
			ServiceULimitsA:        &dockerCompose.ServiceULimits{},
			ServiceULimitsB:        nil,
			expectedServiceULimits: &dockerCompose.ServiceULimits{},
		},
		{
			ServiceULimitsA: &dockerCompose.ServiceULimits{
				NProc: 10,
			},
			ServiceULimitsB: &dockerCompose.ServiceULimits{
				NProc: 10,
			},
			expectedServiceULimits: &dockerCompose.ServiceULimits{
				NProc: 10,
			},
		},
		{
			ServiceULimitsA: &dockerCompose.ServiceULimits{
				NProc: 10,
			},
			ServiceULimitsB: &dockerCompose.ServiceULimits{
				NProc: 20,
			},
			expectedServiceULimits: &dockerCompose.ServiceULimits{
				NProc: 20,
			},
		},
	}

	for i, testCase := range testCases {
		testCase.ServiceULimitsA.MergeLastWin(testCase.ServiceULimitsB)
		require.True(testCase.expectedServiceULimits.Equal(testCase.ServiceULimitsA), "Failed test case %v", i)
	}
}

func TestServiceULimitsNoFile_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.ServiceULimitsNoFile{},
			equalableB:     &dockerCompose.NetworkIPAM{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.ServiceULimitsNoFile{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA:     dockerCompose.NewServiceULimitsNoFile(),
			equalableB:     dockerCompose.NewServiceULimitsNoFile(),
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
			},
			equalableB: &dockerCompose.ServiceULimitsNoFile{
				Soft: 10,
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
			equalableB: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestServiceULimitsNoFile_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		ServiceULimitsNoFileA        *dockerCompose.ServiceULimitsNoFile
		ServiceULimitsNoFileB        *dockerCompose.ServiceULimitsNoFile
		expectedServiceULimitsNoFile *dockerCompose.ServiceULimitsNoFile
	}{
		{
			ServiceULimitsNoFileA:        nil,
			ServiceULimitsNoFileB:        nil,
			expectedServiceULimitsNoFile: nil,
		},
		{
			ServiceULimitsNoFileA:        &dockerCompose.ServiceULimitsNoFile{},
			ServiceULimitsNoFileB:        nil,
			expectedServiceULimitsNoFile: &dockerCompose.ServiceULimitsNoFile{},
		},
		{
			ServiceULimitsNoFileA: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
			ServiceULimitsNoFileB: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
			expectedServiceULimitsNoFile: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
		},
		{
			ServiceULimitsNoFileA: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
			ServiceULimitsNoFileB: &dockerCompose.ServiceULimitsNoFile{
				Hard: 20,
				Soft: 10,
			},
			expectedServiceULimitsNoFile: &dockerCompose.ServiceULimitsNoFile{
				Hard: 20,
				Soft: 10,
			},
		},
		{
			ServiceULimitsNoFileA: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
			ServiceULimitsNoFileB: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 20,
			},
			expectedServiceULimitsNoFile: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 20,
			},
		},
		{
			ServiceULimitsNoFileA: &dockerCompose.ServiceULimitsNoFile{
				Hard: 10,
				Soft: 10,
			},
			ServiceULimitsNoFileB: &dockerCompose.ServiceULimitsNoFile{
				Hard: 20,
				Soft: 20,
			},
			expectedServiceULimitsNoFile: &dockerCompose.ServiceULimitsNoFile{
				Hard: 20,
				Soft: 20,
			},
		},
	}

	for i, testCase := range testCases {
		testCase.ServiceULimitsNoFileA.MergeLastWin(testCase.ServiceULimitsNoFileB)
		require.True(testCase.expectedServiceULimitsNoFile.Equal(testCase.ServiceULimitsNoFileA), "Failed test case %v", i)
	}
}

func TestVolume_Equal(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		equalableA     dockerCompose.Equalable
		equalableB     dockerCompose.Equalable
		expectedResult bool
	}{
		{
			equalableA:     &dockerCompose.Volume{},
			equalableB:     &dockerCompose.NetworkIPAM{},
			expectedResult: false,
		},
		{
			equalableA:     &dockerCompose.Volume{},
			equalableB:     nil,
			expectedResult: false,
		},
		{
			equalableA:     dockerCompose.NewVolume(),
			equalableB:     dockerCompose.NewVolume(),
			expectedResult: true,
		},
		{
			equalableA: &dockerCompose.Volume{
				External: true,
			},
			equalableB: &dockerCompose.Volume{
				External: false,
			},
			expectedResult: false,
		},
		{
			equalableA: &dockerCompose.Volume{
				External: true,
			},
			equalableB: &dockerCompose.Volume{
				External: true,
			},
			expectedResult: true,
		},
	}

	for i, testCase := range testCases {
		require.Equal(testCase.expectedResult, testCase.equalableA.Equal(testCase.equalableB), "Failed test case %v", i)
	}
}

func TestVolume_MergeLastWin(t *testing.T) {
	require := require.New(t)

	testCases := []struct {
		volumeA        *dockerCompose.Volume
		volumeB        *dockerCompose.Volume
		expectedVolume *dockerCompose.Volume
	}{
		{
			volumeA:        nil,
			volumeB:        nil,
			expectedVolume: nil,
		},
		{
			volumeA:        &dockerCompose.Volume{},
			volumeB:        nil,
			expectedVolume: &dockerCompose.Volume{},
		},
		{
			volumeA: &dockerCompose.Volume{
				External: true,
			},
			volumeB: &dockerCompose.Volume{
				External: true,
			},
			expectedVolume: &dockerCompose.Volume{
				External: true,
			},
		},
		{
			volumeA: &dockerCompose.Volume{
				External: true,
			},
			volumeB: &dockerCompose.Volume{
				External: false,
			},
			expectedVolume: &dockerCompose.Volume{
				External: false,
			},
		},
	}

	for i, testCase := range testCases {
		testCase.volumeA.MergeLastWin(testCase.volumeB)
		require.True(testCase.expectedVolume.Equal(testCase.volumeA), "Failed test case %v", i)
	}
}
