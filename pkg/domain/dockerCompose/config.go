package dockerCompose

import (
	"fmt"
	"strings"
)

const (
	environmentDelimiter  string = "="
	labelDelimiter        string = "="
	volumeDelimiter       string = ":"
	portDelimiter         string = ":"
	portProtocolDelimiter string = "/"
)

type Config struct {
	Networks map[string]*Network `json:"networks,omitempty" yaml:"networks,omitempty"`
	Secrets  map[string]*Secret  `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	Services map[string]*Service `json:"services,omitempty" yaml:"services,omitempty"`
	Version  string              `json:"version,omitempty" yaml:"version,omitempty"`
	Volumes  map[string]*Volume  `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (c *Config) Equal(equalable Equalable) bool {
	config, ok := equalable.(*Config)
	if !ok {
		return false
	}

	switch {
	case c == nil && config == nil:
		return true
	case c != nil && config == nil:
		fallthrough
	case c == nil && config != nil:
		return false
	default:
		return EqualStringMap(c.Networks, config.Networks) &&
			EqualStringMap(c.Secrets, config.Secrets) &&
			EqualStringMap(c.Services, config.Services) &&
			c.Version == config.Version &&
			EqualStringMap(c.Volumes, config.Volumes)
	}
}

// ExistsNetwork returns true if a network with the passed named exists.
func (c *Config) ExistsNetwork(name string) bool {
	return ExistsInMap(c.Networks, name)
}

// ExistsSecret returns true if a secret with the passed named exists.
func (c *Config) ExistsSecret(name string) bool {
	return ExistsInMap(c.Secrets, name)
}

// ExistsService returns true if a service with the passed named exists.
func (c *Config) ExistsService(name string) bool {
	return ExistsInMap(c.Services, name)
}

// ExistsVolumes returns true if a volume with the passed named exists.
func (c *Config) ExistsVolume(name string) bool {
	return ExistsInMap(c.Volumes, name)
}

// Merge adds only a missing network, secret, service and volume.
func (c *Config) Merge(config *Config) {
	for name, network := range c.Networks {
		if !c.ExistsNetwork(name) {
			c.Networks[name] = network
		}
	}

	for name, secret := range c.Secrets {
		if !c.ExistsSecret(name) {
			c.Secrets[name] = secret
		}
	}

	for name, service := range c.Services {
		if !c.ExistsService(name) {
			c.Services[name] = service
		}
	}

	for name, volume := range c.Volumes {
		if !c.ExistsVolume(name) {
			c.Volumes[name] = volume
		}
	}
}

// MergeLastWin merges a config and overwrite already existing properties
func (c *Config) MergeLastWin(config *Config) {
	switch {
	case c == nil && config == nil:
		fallthrough
	case c != nil && config == nil:
		return

	// WARN: It's not possible to change the memory pointer c *Config
	// to a new initialized config without returning the Config
	// it self.
	//
	// case c == nil && config != nil:
	// 	c = NewConfig()
	// 	fallthrough

	default:
		c.mergeLastWinNetworks(config.Networks)
		c.mergeLastWinSecrets(config.Secrets)
		c.mergeLastWinServices(config.Services)
		c.mergeLastWinVersion(config.Version)
		c.mergeLastWinVolumes(config.Volumes)
	}
}

func (c *Config) mergeLastWinVersion(version string) {
	if c.Version != version {
		c.Version = version
	}
}

func (c *Config) mergeLastWinNetworks(networks map[string]*Network) {
	for networkName, network := range networks {
		if network == nil {
			continue
		}

		if c.ExistsNetwork(networkName) {
			c.Networks[networkName].MergeLastWin(network)
		} else {
			c.Networks[networkName] = network
		}
	}
}

func (c *Config) mergeLastWinSecrets(secrets map[string]*Secret) {
	for secretName, secret := range secrets {
		if secret == nil {
			continue
		}

		if c.ExistsNetwork(secretName) {
			c.Secrets[secretName].MergeLastWin(secret)
		} else {
			c.Secrets[secretName] = secret
		}
	}
}

func (c *Config) mergeLastWinServices(services map[string]*Service) {
	for serviceName, service := range services {
		if service == nil {
			continue
		}

		if c.ExistsService(serviceName) {
			c.Services[serviceName].MergeLastWin(service)
		} else {
			c.Services[serviceName] = service
		}
	}
}

func (c *Config) mergeLastWinVolumes(volumes map[string]*Volume) {
	for volumeName, volume := range volumes {
		if volume == nil {
			continue
		}

		if c.ExistsNetwork(volumeName) {
			c.Volumes[volumeName].MergeLastWin(volume)
		} else {
			c.Volumes[volumeName] = volume
		}
	}
}

func NewConfig() *Config {
	return &Config{
		Services: make(map[string]*Service),
		Networks: make(map[string]*Network),
		Secrets:  make(map[string]*Secret),
		Volumes:  make(map[string]*Volume),
	}
}

type Network struct {
	External bool         `json:"external,omitempty" yaml:"external,omitempty"`
	Driver   string       `json:"driver,omitempty" yaml:"driver,omitempty"`
	IPAM     *NetworkIPAM `json:"ipam,omitempty" yaml:"ipam,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (n *Network) Equal(equalable Equalable) bool {
	network, ok := equalable.(*Network)
	if !ok {
		return false
	}

	switch {
	case n == nil && network == nil:
		return true
	case n != nil && network == nil:
		fallthrough
	case n == nil && network != nil:
		return false
	default:
		return n.External == network.External &&
			n.Driver == network.Driver &&
			n.IPAM.Equal(network.IPAM)
	}
}

func (n *Network) MergeLastWin(network *Network) {
	switch {
	case n == nil && network == nil:
		fallthrough
	case n != nil && network == nil:
		return

	// WARN: It's not possible to change the memory pointer n *Network
	// to a new initialized network without returning the Network
	// it self.
	//
	// case n == nil && network != nil:
	// 	c = NewCNetwork()
	// 	fallthrough

	default:
		n.mergeLastWinIPAM(network.IPAM)
	}
}

func (n *Network) mergeLastWinIPAM(networkIPAM *NetworkIPAM) {
	if !n.IPAM.Equal(networkIPAM) {
		n.IPAM.MergeLastWin(networkIPAM)
	}
}

func NewNetwork() *Network {
	return &Network{
		External: false,
		IPAM:     new(NetworkIPAM),
	}
}

type NetworkIPAM struct {
	Configs []*NetworkIPAMConfig `json:"config,omitempty" yaml:"config,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (nIPAM *NetworkIPAM) Equal(equalable Equalable) bool {
	networkIPAM, ok := equalable.(*NetworkIPAM)
	if !ok {
		return false
	}

	switch {
	case nIPAM == nil && networkIPAM == nil:
		return true
	case nIPAM != nil && networkIPAM == nil:
		fallthrough
	case nIPAM == nil && networkIPAM != nil:
		return false
	default:
		return Equal(nIPAM.Configs, networkIPAM.Configs)
	}
}

func (nIPAM *NetworkIPAM) MergeLastWin(networkIPAM *NetworkIPAM) {
	switch {
	case nIPAM == nil && networkIPAM == nil:
		fallthrough
	case nIPAM != nil && networkIPAM == nil:
		return

	// WARN: It's not possible to change the memory pointer n *NetworkIPAM
	// to a new initialized networkIPAM without returning the NetworkIPAM
	// it self.
	//
	// case nIPAM == nil && networkIPAM != nil:
	// 	c = NewNetworkIPAM()
	// 	fallthrough

	default:
		nIPAM.mergeLastWinConfig(networkIPAM.Configs)
	}
}

func (nIPAM *NetworkIPAM) mergeLastWinConfig(networkIPAMConfigs []*NetworkIPAMConfig) {
	for _, networkIPAMConfig := range networkIPAMConfigs {
		if !existsInSlice(nIPAM.Configs, networkIPAMConfig) {
			nIPAM.Configs = append(nIPAM.Configs, networkIPAMConfig)
		}
	}
}

func NewNetworkIPAM() *NetworkIPAM {
	return &NetworkIPAM{
		Configs: make([]*NetworkIPAMConfig, 0),
	}
}

type NetworkIPAMConfig struct {
	Subnet string `json:"subnet,omitempty" yaml:"subnet,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (nIPAMConfig *NetworkIPAMConfig) Equal(equalable Equalable) bool {
	networkIPAMConfig, ok := equalable.(*NetworkIPAMConfig)
	if !ok {
		return false
	}

	switch {
	case nIPAMConfig == nil && networkIPAMConfig == nil:
		return true
	case nIPAMConfig != nil && networkIPAMConfig == nil:
		fallthrough
	case nIPAMConfig == nil && networkIPAMConfig != nil:
		return false
	default:
		return nIPAMConfig.Subnet == networkIPAMConfig.Subnet
	}
}

func NewNetworkIPAMConfig() *NetworkIPAMConfig {
	return &NetworkIPAMConfig{}
}

type Secret struct {
	File string `json:"file,omitempty" yaml:"file,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (s *Secret) Equal(equalable Equalable) bool {
	secret, ok := equalable.(*Secret)
	if !ok {
		return false
	}

	switch {
	case s == nil && secret == nil:
		return true
	case s != nil && secret == nil:
		fallthrough
	case s == nil && secret != nil:
		return false
	default:
		return s.File == secret.File
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed secret
// with the existing one.
func (s *Secret) MergeLastWin(secret *Secret) {
	if !s.Equal(secret) {
		s.File = secret.File
	}
}

func NewSecret() *Secret {
	return &Secret{}
}

type Service struct {
	CapabilitiesAdd  []string                   `json:"cap_add,omitempty" yaml:"cap_add,omitempty"`
	CapabilitiesDrop []string                   `json:"cap_drop,omitempty" yaml:"cap_drop,omitempty"`
	Deploy           *ServiceDeploy             `json:"deploy,omitempty" yaml:"deploy,omitempty"`
	Environments     []string                   `json:"environment,omitempty" yaml:"environment,omitempty"`
	ExtraHosts       []string                   `json:"extra_hosts,omitempty" yaml:"extra_hosts,omitempty"`
	Image            string                     `json:"image,omitempty" yaml:"image,omitempty"`
	Labels           []string                   `json:"labels,omitempty" yaml:"labels,omitempty"`
	Networks         map[string]*ServiceNetwork `json:"networks,omitempty" yaml:"networks,omitempty"`
	Ports            []string                   `json:"ports,omitempty" yaml:"ports,omitempty"`
	Secrets          []string                   `json:"secrets,omitempty" yaml:"secrets,omitempty"`
	ULimits          *ServiceULimits            `json:"ulimits,omitempty" yaml:"ulimits,omitempty"`
	Volumes          []string                   `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}

// ExistsEnvironment returns true if the passed name of environment variable is
// already present.
func (s *Service) ExistsEnvironment(name string) bool {
	for _, environment := range s.Environments {
		key, _ := splitStringInKeyValue(environment, environmentDelimiter)
		if key == name {
			return true
		}
	}

	return false
}

// ExistsLabel returns true if the passed label name is already present.
func (s *Service) ExistsLabel(name string) bool {
	for _, label := range s.Labels {
		key, _ := splitStringInKeyValue(label, labelDelimiter)
		if key == name {
			return true
		}
	}

	return false
}

// ExistsPort returns true if the port definition is already present.
func (s *Service) ExistsPort(src string, dest string, protocol string) bool {
	for _, port := range s.Ports {
		s, d, p := splitStringInPort(port)
		if s == src && d == dest && p == protocol {
			return true
		}
	}

	return false
}

// ExistsDestinationPort returns true if the destination port is already used.
func (s *Service) ExistsDestinationPort(dest string) bool {
	for _, port := range s.Ports {
		_, d, _ := splitStringInPort(port)
		if d == dest {
			return true
		}
	}

	return false
}

// ExistsSourcePort returns true if the source port is already used.
func (s *Service) ExistsSourcePort(src string) bool {
	for _, port := range s.Ports {
		s, _, _ := splitStringInPort(port)
		if s == src {
			return true
		}
	}

	return false
}

// ExistsVolume returns true if the volume definition is already present.
func (s *Service) ExistsVolume(src string, dest string, perm string) bool {
	for _, volume := range s.Volumes {
		s, d, p := splitStringInVolume(volume)
		if s == src && d == dest && p == perm {
			return true
		}
	}

	return false
}

// ExistsDestinationVolume returns true if the volume definition is already present.
func (s *Service) ExistsDestinationVolume(dest string) bool {
	for _, volume := range s.Volumes {
		_, d, _ := splitStringInVolume(volume)
		if d == dest {
			return true
		}
	}

	return false
}

// ExistsSourceVolume returns true if the volume definition is already present.
func (s *Service) ExistsSourceVolume(src string) bool {
	for _, volume := range s.Volumes {
		s, _, _ := splitStringInVolume(volume)
		if s == src {
			return true
		}
	}

	return false
}

// Equal returns true if the passed equalable is equal
func (s *Service) Equal(equalable Equalable) bool {
	service, ok := equalable.(*Service)
	if !ok {
		return false
	}

	switch {
	case s == nil && service == nil:
		return true
	case s != nil && service == nil:
		fallthrough
	case s == nil && service != nil:
		return false
	default:
		return equalSlice(s.CapabilitiesAdd, service.CapabilitiesAdd) &&
			equalSlice(s.CapabilitiesDrop, service.CapabilitiesDrop) &&
			s.Deploy.Equal(service.Deploy) &&
			equalSlice(s.Environments, service.Environments) &&
			equalSlice(s.ExtraHosts, service.ExtraHosts) &&
			s.Image == service.Image &&
			equalSlice(s.Labels, service.Labels) &&
			EqualStringMap(s.Networks, service.Networks) &&
			equalSlice(s.Ports, service.Ports) &&
			equalSlice(s.Secrets, service.Secrets) &&
			s.ULimits.Equal(service.ULimits) &&
			equalSlice(s.Volumes, service.Volumes)
	}
}

func (s *Service) MergeFirstWin(service *Service) {
	switch {
	case s == nil && service == nil:
		fallthrough
	case s != nil && service == nil:
		return

	// WARN: It's not possible to change the memory pointer s *Service
	// to a new initialized service without returning the Service
	// it self.
	//
	// case s == nil && service != nil:
	// 	s = NewService()
	// 	fallthrough

	default:
		s.mergeFirstWinCapabilitiesAdd(service.CapabilitiesAdd)
		s.mergeFirstWinCapabilitiesDrop(service.CapabilitiesDrop)
		s.mergeFirstWinDeploy(service.Deploy)
		s.mergeFirstWinEnvironments(service.Environments)
		s.mergeFirstWinExtraHosts(service.ExtraHosts)
		s.mergeFirstWinImage(service.Image)
		s.mergeFirstWinLabels(service.Labels)
		s.mergeFirstWinNetworks(service.Networks)
		s.mergeFirstWinPorts(service.Ports)
		s.mergeFirstWinSecrets(service.Secrets)
		s.mergeFirstWinULimits(service.ULimits)
		s.mergeFirstWinVolumes(service.Volumes)
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed secret
// with the existing one.
func (s *Service) MergeLastWin(service *Service) {
	switch {
	case s == nil && service == nil:
		fallthrough
	case s != nil && service == nil:
		return

	// WARN: It's not possible to change the memory pointer s *Service
	// to a new initialized service without returning the Service
	// it self.
	//
	// case s == nil && service != nil:
	// 	s = NewService()
	// 	fallthrough

	default:
		s.mergeLastWinCapabilitiesAdd(service.CapabilitiesAdd)
		s.mergeLastWinCapabilitiesDrop(service.CapabilitiesDrop)
		s.mergeLastWinDeploy(service.Deploy)
		s.mergeLastWinEnvironments(service.Environments)
		s.mergeLastWinExtraHosts(service.ExtraHosts)
		s.mergeLastWinImage(service.Image)
		s.mergeLastWinLabels(service.Labels)
		s.mergeLastWinNetworks(service.Networks)
		s.mergeLastWinPorts(service.Ports)
		s.mergeLastWinSecrets(service.Secrets)
		s.mergeLastWinULimits(service.ULimits)
		s.mergeLastWinVolumes(service.Volumes)
	}
}

func (s *Service) mergeFirstWinCapabilitiesAdd(capabilitiesAdd []string) {
	for _, capabilityAdd := range capabilitiesAdd {
		if !existsInSlice(s.CapabilitiesAdd, capabilityAdd) && len(capabilityAdd) > 0 {
			s.CapabilitiesAdd = append(s.CapabilitiesAdd, capabilityAdd)
		}
	}
}

func (s *Service) mergeFirstWinCapabilitiesDrop(capabilitiesDrop []string) {
	for _, capabilityDrop := range capabilitiesDrop {
		if !existsInSlice(s.CapabilitiesAdd, capabilityDrop) && len(capabilityDrop) > 0 {
			s.CapabilitiesDrop = append(s.CapabilitiesDrop, capabilityDrop)
		}
	}
}

func (s *Service) mergeFirstWinDeploy(deploy *ServiceDeploy) {
	switch {
	case s.Deploy == nil && deploy != nil:
		s.Deploy = deploy
	case s.Deploy != nil && deploy == nil:
		fallthrough
	case s.Deploy == nil && deploy == nil:
		return
	default:
		s.Deploy.MergeFirstWin(deploy)
	}
}

func (s *Service) mergeFirstWinEnvironments(environments []string) {
	switch {
	case s.Environments == nil && environments != nil:
		s.Environments = environments
	case s.Environments != nil && environments == nil:
		fallthrough
	case s.Environments == nil && environments == nil:
		return
	default:
		for _, environment := range environments {
			if len(environment) <= 0 {
				continue
			}

			key, value := splitStringInKeyValue(environment, environmentDelimiter)
			if !s.ExistsEnvironment(key) {
				s.SetEnvironment(key, value)
			}
		}
	}
}

func (s *Service) mergeFirstWinImage(image string) {
	switch {
	case len(s.Image) == 0 && len(image) != 0:
		s.Image = image
	case len(s.Image) != 0 && len(image) == 0:
		fallthrough
	case len(s.Image) == 0 && len(image) == 0:
		fallthrough
	default:
		return
	}
}

func (s *Service) mergeFirstWinExtraHosts(extraHosts []string) {
	for _, extraHost := range extraHosts {
		if !existsInSlice(s.ExtraHosts, extraHost) && len(extraHost) > 0 {
			s.ExtraHosts = append(s.ExtraHosts, extraHost)
		}
	}
}

func (s *Service) mergeFirstWinLabels(labels []string) {
	switch {
	case s.Labels == nil && labels != nil:
		s.Labels = labels
	case s.Labels != nil && labels == nil:
		fallthrough
	case s.Labels == nil && labels == nil:
		return
	default:
		for _, label := range labels {
			if len(label) <= 0 {
				continue
			}

			key, value := splitStringInKeyValue(label, labelDelimiter)
			if !s.ExistsLabel(key) {
				s.SetLabel(key, value)
			}
		}
	}
}

func (s *Service) mergeFirstWinNetworks(networks map[string]*ServiceNetwork) {
	switch {
	case s.Networks == nil && networks != nil:
		s.Networks = networks
	case s.Networks != nil && networks == nil:
		fallthrough
	case s.Networks == nil && networks == nil:
		return
	default:
		for name, network := range networks {
			if _, exists := s.Networks[name]; exists {
				s.Networks[name].MergeFirstWin(network)
			} else {
				s.Networks[name] = network
			}
		}
	}
}

func (s *Service) mergeFirstWinPorts(ports []string) {
	switch {
	case s.Ports == nil && ports != nil:
		s.Ports = ports
	case s.Ports != nil && ports == nil:
		fallthrough
	case s.Ports == nil && ports == nil:
		return
	default:
		for _, port := range ports {
			src, dest, protocol := splitStringInPort(port)
			if !s.ExistsDestinationPort(dest) {
				s.SetPort(src, dest, protocol)
			}
		}
	}
}

func (s *Service) mergeFirstWinSecrets(secrets []string) {
	for _, secret := range secrets {
		if !existsInSlice(s.Secrets, secret) && len(secret) > 0 {
			s.Secrets = append(s.Secrets, secret)
		}
	}
}

func (s *Service) mergeFirstWinULimits(uLimits *ServiceULimits) {
	switch {
	case s.ULimits == nil && uLimits != nil:
		s.ULimits = uLimits
	case s.ULimits != nil && uLimits == nil:
		fallthrough
	case s.ULimits == nil && uLimits == nil:
		return
	default:
		s.ULimits.MergeFirstWin(uLimits)
	}
}

func (s *Service) mergeFirstWinVolumes(volumes []string) {
	switch {
	case s.Volumes == nil && volumes != nil:
		s.Volumes = volumes
	case s.Volumes != nil && volumes == nil:
		fallthrough
	case s.Volumes == nil && volumes == nil:
		return
	default:
		for _, volume := range volumes {
			src, dest, perm := splitStringInVolume(volume)
			if !s.ExistsDestinationVolume(dest) {
				s.SetVolume(src, dest, perm)
			}
		}
	}
}

func (s *Service) mergeLastWinCapabilitiesAdd(capabilitiesAdd []string) {
	for _, capabilityAdd := range capabilitiesAdd {
		if !existsInSlice(s.CapabilitiesAdd, capabilityAdd) {
			s.CapabilitiesAdd = append(s.CapabilitiesAdd, capabilityAdd)
		}
	}
}

func (s *Service) mergeLastWinCapabilitiesDrop(capabilitiesDrop []string) {
	for _, capabilityDrop := range capabilitiesDrop {
		if !existsInSlice(s.CapabilitiesAdd, capabilityDrop) {
			s.CapabilitiesDrop = append(s.CapabilitiesDrop, capabilityDrop)
		}
	}
}

func (s *Service) mergeLastWinDeploy(deploy *ServiceDeploy) {
	switch {
	case s.Deploy == nil && deploy != nil:
		s.Deploy = deploy
	case s.Deploy != nil && deploy == nil:
		fallthrough
	case s.Deploy == nil && deploy == nil:
		return
	default:
		s.Deploy.MergeLastWin(deploy)
	}
}

func (s *Service) mergeLastWinEnvironments(environments []string) {
	switch {
	case s.Environments == nil && environments != nil:
		s.Environments = environments
	case s.Environments != nil && environments == nil:
		fallthrough
	case s.Environments == nil && environments == nil:
		return
	default:
		for _, environment := range environments {
			key, value := splitStringInKeyValue(environment, environmentDelimiter)
			s.SetEnvironment(key, value)
		}
	}
}

func (s *Service) mergeLastWinImage(image string) {
	switch {
	case len(s.Image) == 0 && len(image) != 0:
		s.Image = image
	case len(s.Image) != 0 && len(image) == 0:
		fallthrough
	case len(s.Image) == 0 && len(image) == 0:
		return
	default:
		if s.Image != image {
			s.Image = image
		}
	}
}

func (s *Service) mergeLastWinExtraHosts(extraHosts []string) {
	for _, extraHost := range extraHosts {
		if !existsInSlice(s.ExtraHosts, extraHost) {
			s.ExtraHosts = append(s.ExtraHosts, extraHost)
		}
	}
}

func (s *Service) mergeLastWinLabels(labels []string) {
	switch {
	case s.Labels == nil && labels != nil:
		s.Labels = labels
	case s.Labels != nil && labels == nil:
		fallthrough
	case s.Labels == nil && labels == nil:
		return
	default:
		for _, label := range labels {
			key, value := splitStringInKeyValue(label, labelDelimiter)
			s.SetLabel(key, value)
		}
	}
}

func (s *Service) mergeLastWinNetworks(networks map[string]*ServiceNetwork) {
	switch {
	case s.Networks == nil && networks != nil:
		s.Networks = networks
	case s.Networks != nil && networks == nil:
		fallthrough
	case s.Networks == nil && networks == nil:
		return
	default:
		for name, network := range networks {
			if _, exists := s.Networks[name]; exists {
				s.Networks[name].MergeLastWin(network)
			} else {
				s.Networks[name] = network
			}
		}
	}
}

func (s *Service) mergeLastWinPorts(ports []string) {
	switch {
	case s.Ports == nil && ports != nil:
		s.Ports = ports
	case s.Ports != nil && ports == nil:
		fallthrough
	case s.Ports == nil && ports == nil:
		return
	default:
		for _, port := range ports {
			src, dest, protocol := splitStringInPort(port)
			s.SetPort(src, dest, protocol)
		}
	}
}

func (s *Service) mergeLastWinSecrets(secrets []string) {
	for _, secret := range secrets {
		if !existsInSlice(s.Secrets, secret) {
			s.Secrets = append(s.Secrets, secret)
		}
	}
}

func (s *Service) mergeLastWinULimits(uLimits *ServiceULimits) {
	switch {
	case s.ULimits == nil && uLimits != nil:
		s.ULimits = uLimits
	case s.ULimits != nil && uLimits == nil:
		fallthrough
	case s.ULimits == nil && uLimits == nil:
		return
	default:
		s.ULimits.MergeLastWin(uLimits)
	}
}

func (s *Service) mergeLastWinVolumes(volumes []string) {
	switch {
	case s.Volumes == nil && volumes != nil:
		s.Volumes = volumes
	case s.Volumes != nil && volumes == nil:
		fallthrough
	case s.Volumes == nil && volumes == nil:
		return
	default:
		for _, volume := range volumes {
			src, dest, perm := splitStringInVolume(volume)
			s.SetVolume(src, dest, perm)
		}
	}
}

// RemoveEnvironment remove all found environment variable from the internal
// slice matching by the passed name.
func (s *Service) RemoveEnvironment(name string) {
	environments := make([]string, 0)
	for _, environment := range s.Environments {
		key, value := splitStringInKeyValue(environment, environmentDelimiter)
		if key != name {
			environments = append(environments, fmt.Sprintf("%s%s%s", key, environmentDelimiter, value))
		}
	}
	s.Environments = environments
}

// RemoveLabel remove all found labels from the internal slice matching by the
// passed name.
func (s *Service) RemoveLabel(name string) {
	labels := make([]string, 0)
	for _, label := range s.Labels {
		key, value := splitStringInKeyValue(label, labelDelimiter)
		if key != name {
			labels = append(labels, fmt.Sprintf("%s%s%s", key, labelDelimiter, value))
		}
	}
	s.Labels = labels
}

// RemovePort remove all found ports from the internal slice matching by the
// passed dest port.
func (s *Service) RemovePort(dest string) {
	ports := make([]string, 0)
	for _, port := range s.Ports {
		srcPort, destPort, protocol := splitStringInPort(port)

		switch {
		case destPort == dest && len(protocol) <= 0:
			s.Ports = append(s.Ports, fmt.Sprintf("%s%s%s", srcPort, portDelimiter, destPort))
		case destPort == dest && len(protocol) > 0:
			s.Ports = append(s.Ports, fmt.Sprintf("%s%s%s%s%s", srcPort, portDelimiter, destPort, portProtocolDelimiter, protocol))
		}
	}
	s.Ports = ports
}

// RemoveVolume remove all found volumes from the internal slice matching by the
// dest path.
func (s *Service) RemoveVolume(dest string) {
	volumes := make([]string, 0)
	for _, volume := range s.Volumes {
		srcPath, destPath, perm := splitStringInVolume(volume)

		switch {
		case destPath == dest && len(perm) <= 0:
			s.Volumes = append(s.Volumes, fmt.Sprintf("%s%s%s", srcPath, volumeDelimiter, destPath))
		case destPath == dest && len(perm) > 0:
			s.Volumes = append(s.Volumes, fmt.Sprintf("%s%s%s%s%s", srcPath, volumeDelimiter, destPath, volumeDelimiter, perm))
		}
	}
	s.Volumes = volumes
}

// SetEnvironment add or overwrite an existing environment variable.
func (s *Service) SetEnvironment(name string, value string) {
	s.RemoveEnvironment(name)
	s.Environments = append(s.Environments, fmt.Sprintf("%s%s%s", name, environmentDelimiter, value))
}

// SetLabel add or overwrite an existing label.
func (s *Service) SetLabel(name string, value string) {
	s.RemoveLabel(name)
	s.Labels = append(s.Labels, fmt.Sprintf("%s%s%s", name, labelDelimiter, value))
}

// SetPort add or overwrite an existing port.
func (s *Service) SetPort(src string, dest string, protocol string) {
	s.RemovePort(dest)
	if len(protocol) <= 0 {
		s.Ports = append(s.Ports, fmt.Sprintf("%s%s%s", src, volumeDelimiter, dest))
	} else {
		s.Ports = append(s.Ports, fmt.Sprintf("%s%s%s%s%s", src, portDelimiter, dest, portProtocolDelimiter, protocol))
	}
}

// SetVolume add or overwrite an existing volume.
func (s *Service) SetVolume(src string, dest string, perm string) {
	s.RemoveVolume(dest)
	if len(perm) <= 0 {
		s.Volumes = append(s.Volumes, fmt.Sprintf("%s%s%s", src, volumeDelimiter, dest))
	} else {
		s.Volumes = append(s.Volumes, fmt.Sprintf("%s%s%s%s%s", src, volumeDelimiter, dest, volumeDelimiter, perm))
	}
}

// NewService returns an empty initialized Service.
func NewService() *Service {
	return &Service{
		CapabilitiesAdd:  make([]string, 0),
		CapabilitiesDrop: make([]string, 0),
		Deploy:           new(ServiceDeploy),
		Environments:     make([]string, 0),
		ExtraHosts:       make([]string, 0),
		Labels:           make([]string, 0),
		Networks:         make(map[string]*ServiceNetwork),
		Ports:            make([]string, 0),
		Secrets:          make([]string, 0),
		ULimits:          new(ServiceULimits),
		Volumes:          make([]string, 0),
	}
}

type ServiceDeploy struct {
	Resources *ServiceDeployResources `json:"resources" yaml:"resources"`
}

// Equal returns true if the passed equalable is equal
func (sd *ServiceDeploy) Equal(equalable Equalable) bool {
	serviceDeploy, ok := equalable.(*ServiceDeploy)
	if !ok {
		return false
	}

	switch {
	case sd == nil && serviceDeploy == nil:
		return true
	case sd != nil && serviceDeploy == nil:
		fallthrough
	case sd == nil && serviceDeploy != nil:
		return false
	default:
		return sd.Resources.Equal(serviceDeploy.Resources)
	}
}

// MergeFirstWin merges adds or overwrite the attributes of the passed
// serviceDeploy with the existing one.
func (sd *ServiceDeploy) MergeFirstWin(serviceDeploy *ServiceDeploy) {
	switch {
	case sd == nil && serviceDeploy == nil:
		fallthrough
	case sd != nil && serviceDeploy == nil:
		return

	// WARN: It's not possible to change the memory pointer sd *ServiceDeploy
	// to a new initialized serviceDeploy without returning the ServiceDeploy
	// it self.
	//
	// case sd == nil && serviceDeploy != nil:
	// 	sd = NewServiceDeploy()
	// 	fallthrough

	default:
		sd.mergeFirstWinDeployResources(serviceDeploy.Resources)
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed
// serviceDeploy with the existing one.
func (sd *ServiceDeploy) MergeLastWin(serviceDeploy *ServiceDeploy) {
	switch {
	case sd == nil && serviceDeploy == nil:
		fallthrough
	case sd != nil && serviceDeploy == nil:
		return

	// WARN: It's not possible to change the memory pointer sd *ServiceDeploy
	// to a new initialized serviceDeploy without returning the ServiceDeploy
	// it self.
	//
	// case sd == nil && serviceDeploy != nil:
	// 	sd = NewServiceDeploy()
	// 	fallthrough

	default:
		sd.mergeLastWinDeployResources(serviceDeploy.Resources)
	}
}

func (sd *ServiceDeploy) mergeFirstWinDeployResources(resources *ServiceDeployResources) {
	switch {
	case sd.Resources == nil && resources != nil:
		sd.Resources = resources
	case sd.Resources != nil && resources == nil:
		fallthrough
	case sd.Resources == nil && resources == nil:
		return
	default:
		sd.Resources.MergeFirstWin(resources)
	}
}

func (sd *ServiceDeploy) mergeLastWinDeployResources(resources *ServiceDeployResources) {
	switch {
	case sd.Resources == nil && resources != nil:
		sd.Resources = resources
	case sd.Resources != nil && resources == nil:
		fallthrough
	case sd.Resources == nil && resources == nil:
		return
	default:
		sd.Resources.MergeLastWin(resources)
	}
}

func NewServiceDeploy() *ServiceDeploy {
	return &ServiceDeploy{
		Resources: new(ServiceDeployResources),
	}
}

type ServiceDeployResources struct {
	Limits       *ServiceDeployResourcesLimits `json:"limits,omitempty" yaml:"limits,omitempty"`
	Reservations *ServiceDeployResourcesLimits `json:"reservations,omitempty" yaml:"reservations,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (sdr *ServiceDeployResources) Equal(equalable Equalable) bool {
	serviceDeployResources, ok := equalable.(*ServiceDeployResources)
	if !ok {
		return false
	}

	switch {
	case sdr == nil && serviceDeployResources == nil:
		return true
	case sdr != nil && serviceDeployResources == nil:
		fallthrough
	case sdr == nil && serviceDeployResources != nil:
		return false
	default:
		return sdr.Limits.Equal(serviceDeployResources.Limits) &&
			sdr.Reservations.Equal(serviceDeployResources.Reservations)
	}
}

// MergeFirstWin adds only attributes of the passed serviceDeployResources if
// they are not already exists.
func (sdr *ServiceDeployResources) MergeFirstWin(serviceDeployResources *ServiceDeployResources) {
	switch {
	case sdr == nil && serviceDeployResources == nil:
		fallthrough
	case sdr != nil && serviceDeployResources == nil:
		return

	// WARN: It's not possible to change the memory pointer sdr *ServiceDeployResources
	// to a new initialized serviceDeployResources without returning the
	// serviceDeployResources it self.
	case sdr == nil && serviceDeployResources != nil:
		sdr = NewServiceDeployResources()
		fallthrough
	default:
		sdr.mergeFirstWinLimits(serviceDeployResources.Limits)
		sdr.mergeFirstWinReservations(serviceDeployResources.Reservations)
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed
// serviceDeployResources with the existing one.
func (sdr *ServiceDeployResources) MergeLastWin(serviceDeployResources *ServiceDeployResources) {
	switch {
	case sdr == nil && serviceDeployResources == nil:
		fallthrough
	case sdr != nil && serviceDeployResources == nil:
		return

	// WARN: It's not possible to change the memory pointer sdr *ServiceDeployResources
	// to a new initialized serviceDeployResources without returning the
	// serviceDeployResources it self.
	case sdr == nil && serviceDeployResources != nil:
		sdr = NewServiceDeployResources()
		fallthrough
	default:
		sdr.mergeLastWinLimits(serviceDeployResources.Limits)
		sdr.mergeLastWinReservations(serviceDeployResources.Reservations)
	}
}

func (sdr *ServiceDeployResources) mergeFirstWinLimits(limits *ServiceDeployResourcesLimits) {
	switch {
	case sdr.Limits == nil && limits != nil:
		sdr.Limits = limits
	case sdr.Limits != nil && limits == nil:
		fallthrough
	case sdr.Limits == nil && limits == nil:
		return
	default:
		sdr.Limits.MergeFirstWin(limits)
	}
}

func (sdr *ServiceDeployResources) mergeFirstWinReservations(reservations *ServiceDeployResourcesLimits) {
	switch {
	case sdr.Reservations == nil && reservations != nil:
		sdr.Reservations = reservations
	case sdr.Reservations != nil && reservations == nil:
		fallthrough
	case sdr.Reservations == nil && reservations == nil:
		return
	default:
		sdr.Reservations.MergeFirstWin(reservations)
	}
}

func (sdr *ServiceDeployResources) mergeLastWinLimits(limits *ServiceDeployResourcesLimits) {
	switch {
	case sdr.Limits == nil && limits != nil:
		sdr.Limits = limits
	case sdr.Limits != nil && limits == nil:
		fallthrough
	case sdr.Limits == nil && limits == nil:
		return
	default:
		sdr.Limits.MergeLastWin(limits)
	}
}

func (sdr *ServiceDeployResources) mergeLastWinReservations(reservations *ServiceDeployResourcesLimits) {
	switch {
	case sdr.Reservations == nil && reservations != nil:
		sdr.Reservations = reservations
	case sdr.Reservations != nil && reservations == nil:
		fallthrough
	case sdr.Reservations == nil && reservations == nil:
		return
	default:
		sdr.Reservations.MergeLastWin(reservations)
	}
}

func NewServiceDeployResources() *ServiceDeployResources {
	return &ServiceDeployResources{
		Limits:       new(ServiceDeployResourcesLimits),
		Reservations: new(ServiceDeployResourcesLimits),
	}
}

type ServiceDeployResourcesLimits struct {
	CPUs   string `json:"cpus,omitempty" yaml:"cpus,omitempty"`
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (sdrl *ServiceDeployResourcesLimits) Equal(equalable Equalable) bool {
	serviceDeployResourcesLimits, ok := equalable.(*ServiceDeployResourcesLimits)
	if !ok {
		return false
	}

	switch {
	case sdrl == nil && serviceDeployResourcesLimits == nil:
		return true
	case sdrl != nil && serviceDeployResourcesLimits == nil:
		fallthrough
	case sdrl == nil && serviceDeployResourcesLimits != nil:
		return false
	default:
		return sdrl.CPUs == serviceDeployResourcesLimits.CPUs &&
			sdrl.Memory == serviceDeployResourcesLimits.Memory
	}
}

// MergeFirstWin adds only attributes of the passed serviceDeployResourcesLimits
// if they are not already exists.
func (sdrl *ServiceDeployResourcesLimits) MergeFirstWin(serviceDeployResourcesLimits *ServiceDeployResourcesLimits) {
	switch {
	case sdrl == nil && serviceDeployResourcesLimits == nil:
		fallthrough
	case sdrl != nil && serviceDeployResourcesLimits == nil:
		return

	// WARN: It's not possible to change the memory pointer sdrl *ServiceDeployResourcesLimits
	// to a new initialized serviceDeployResourcesLimits without returning the
	// serviceDeployResourcesLimits it self.
	//
	// case sdrl == nil && serviceDeployResourcesLimits != nil:
	// 	sdrl = NewServiceDeployResourcesLimits()
	// 	fallthrough

	default:
		sdrl.mergeFirstWinCPUs(serviceDeployResourcesLimits.CPUs)
		sdrl.mergeFirstWinMemory(serviceDeployResourcesLimits.Memory)
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed
// serviceDeployResourcesLimits with the existing one.
func (sdrl *ServiceDeployResourcesLimits) MergeLastWin(serviceDeployResourcesLimits *ServiceDeployResourcesLimits) {
	switch {
	case sdrl == nil && serviceDeployResourcesLimits == nil:
		fallthrough
	case sdrl != nil && serviceDeployResourcesLimits == nil:
		return

	// WARN: It's not possible to change the memory pointer sdrl *ServiceDeployResourcesLimits
	// to a new initialized serviceDeployResourcesLimits without returning the
	// serviceDeployResourcesLimits it self.
	//
	// case sdrl == nil && serviceDeployResourcesLimits != nil:
	// 	sdrl = NewServiceDeployResourcesLimits()
	// 	fallthrough

	default:
		sdrl.mergeLastWinCPUs(serviceDeployResourcesLimits.CPUs)
		sdrl.mergeLastWinMemory(serviceDeployResourcesLimits.Memory)
	}
}

func (sdrl *ServiceDeployResourcesLimits) mergeFirstWinCPUs(cpus string) {
	if len(sdrl.CPUs) <= 0 {
		sdrl.CPUs = cpus
	}
}

func (sdrl *ServiceDeployResourcesLimits) mergeFirstWinMemory(memory string) {
	if len(sdrl.Memory) <= 0 {
		sdrl.Memory = memory
	}
}

func (sdrl *ServiceDeployResourcesLimits) mergeLastWinCPUs(cpus string) {
	if sdrl.CPUs != cpus {
		sdrl.CPUs = cpus
	}
}

func (sdrl *ServiceDeployResourcesLimits) mergeLastWinMemory(memory string) {
	if sdrl.Memory != memory {
		sdrl.Memory = memory
	}
}

func NewServiceDeployResourcesLimits() *ServiceDeployResourcesLimits {
	return &ServiceDeployResourcesLimits{}
}

type ServiceNetwork struct {
	Aliases []string `json:"aliases,omitempty" yaml:"aliases,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (sn *ServiceNetwork) Equal(equalable Equalable) bool {
	serviceNetwork, ok := equalable.(*ServiceNetwork)
	if !ok {
		return false
	}

	switch {
	case sn == nil && serviceNetwork == nil:
		return true
	case sn != nil && serviceNetwork == nil:
		fallthrough
	case sn == nil && serviceNetwork != nil:
		return false
	default:
		return equalSlice(sn.Aliases, serviceNetwork.Aliases)
	}
}

// MergeFirstWin adds only attributes of the passed
// serviceNetwork if they are undefined.
func (sn *ServiceNetwork) MergeFirstWin(serviceNetwork *ServiceNetwork) {
	switch {
	case sn == nil && serviceNetwork == nil:
		fallthrough
	case sn != nil && serviceNetwork == nil:
		return

	// WARN: It's not possible to change the memory pointer sn *ServiceNetwork to a new
	// initialized ServiceNetwork without returning the serviceNetwork it self.
	//
	// case l == nil && serviceULimits != nil:
	// 	l = NewServiceULimits()
	// 	fallthrough

	case sn == nil && serviceNetwork != nil:
		sn = NewServiceNetwork()
		fallthrough
	default:
		sn.mergeFirstWinAliases(serviceNetwork.Aliases)
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed
// serviceNetwork with the existing one.
func (sn *ServiceNetwork) MergeLastWin(serviceNetwork *ServiceNetwork) {
	switch {
	case sn == nil && serviceNetwork == nil:
		fallthrough
	case sn != nil && serviceNetwork == nil:
		return

	// WARN: It's not possible to change the memory pointer sn *ServiceNetwork to a new
	// initialized ServiceNetwork without returning the serviceNetwork it self.
	//
	// case l == nil && serviceULimits != nil:
	// 	l = NewServiceULimits()
	// 	fallthrough

	case sn == nil && serviceNetwork != nil:
		sn = NewServiceNetwork()
		fallthrough
	default:
		sn.mergeLastWinAliases(serviceNetwork.Aliases)
	}
}

func (sn *ServiceNetwork) mergeFirstWinAliases(aliases []string) {
	for _, alias := range aliases {
		if !existsInSlice(sn.Aliases, alias) && len(alias) > 0 {
			sn.Aliases = append(sn.Aliases, alias)
		}
	}
}

func (sn *ServiceNetwork) mergeLastWinAliases(aliases []string) {
	for _, alias := range aliases {
		if !existsInSlice(sn.Aliases, alias) && len(alias) > 0 {
			sn.Aliases = append(sn.Aliases, alias)
		}
	}
}

func NewServiceNetwork() *ServiceNetwork {
	return &ServiceNetwork{
		Aliases: make([]string, 0),
	}
}

type ServiceULimits struct {
	NProc  uint                  `json:"nproc,omitempty" yaml:"nproc,omitempty"`
	NoFile *ServiceULimitsNoFile `json:"nofile,omitempty" yaml:"nofile,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (l *ServiceULimits) Equal(equalable Equalable) bool {
	serviceULimits, ok := equalable.(*ServiceULimits)
	if !ok {
		return false
	}

	switch {
	case l == nil && serviceULimits == nil:
		return true
	case l != nil && serviceULimits == nil:
		fallthrough
	case l == nil && serviceULimits != nil:
		return false
	default:
		return l.NProc == serviceULimits.NProc &&
			l.NoFile.Equal(serviceULimits.NoFile)
	}
}

// MergeFirstWin adds only the attributes of the passed ServiceULimits they are
// undefined.
func (l *ServiceULimits) MergeFirstWin(serviceULimits *ServiceULimits) {
	switch {
	case l == nil && serviceULimits == nil:
		fallthrough
	case l != nil && serviceULimits == nil:
		return

	// WARN: It's not possible to change the memory pointer l *ServiceULimits to a new
	// initialized ServiceULimits without returning the serviceULimits it self.
	//
	// case l == nil && serviceULimits != nil:
	// 	l = NewServiceULimits()
	// 	fallthrough

	default:
		l.mergeFirstWinNProc(serviceULimits.NProc)
		l.mergeFirstWinNoFile(serviceULimits.NoFile)
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed
// ServiceULimits with the existing one.
func (l *ServiceULimits) MergeLastWin(serviceULimits *ServiceULimits) {
	switch {
	case l == nil && serviceULimits == nil:
		fallthrough
	case l != nil && serviceULimits == nil:
		return

	// WARN: It's not possible to change the memory pointer l *ServiceULimits to a new
	// initialized ServiceULimits without returning the serviceULimits it self.
	//
	// case l == nil && serviceULimits != nil:
	// 	l = NewServiceULimits()
	// 	fallthrough

	default:
		l.mergeLastWinNProc(serviceULimits.NProc)
		l.mergeLastWinNoFile(serviceULimits.NoFile)
	}
}

func (l *ServiceULimits) mergeFirstWinNProc(nproc uint) {
	if l.NProc != nproc {
		return
	}
	l.NProc = nproc
}

func (l *ServiceULimits) mergeFirstWinNoFile(noFile *ServiceULimitsNoFile) {
	if !l.NoFile.Equal(noFile) {
		l.NoFile.MergeFirstWin(noFile)
	}
}

func (l *ServiceULimits) mergeLastWinNProc(nproc uint) {
	if l.NProc != nproc {
		l.NProc = nproc
	}
}

func (l *ServiceULimits) mergeLastWinNoFile(noFile *ServiceULimitsNoFile) {
	if !l.NoFile.Equal(noFile) {
		l.NoFile.MergeLastWin(noFile)
	}
}

func NewServiceULimits() *ServiceULimits {
	return &ServiceULimits{
		NoFile: new(ServiceULimitsNoFile),
	}
}

type ServiceULimitsNoFile struct {
	Hard uint `json:"hard" yaml:"hard"`
	Soft uint `json:"soft" yaml:"soft"`
}

// Equal returns true if the passed equalable is equal
func (nf *ServiceULimitsNoFile) Equal(equalable Equalable) bool {
	serviceULimitsNoFile, ok := equalable.(*ServiceULimitsNoFile)
	if !ok {
		return false
	}

	switch {
	case nf == nil && serviceULimitsNoFile == nil:
		return true
	case nf != nil && serviceULimitsNoFile == nil:
		fallthrough
	case nf == nil && serviceULimitsNoFile != nil:
		return false
	default:
		return nf.Hard == serviceULimitsNoFile.Hard &&
			nf.Soft == serviceULimitsNoFile.Soft
	}
}

// MergeFirstWin adds only the attributes of the passed ServiceULimits they are
// undefined.
func (nf *ServiceULimitsNoFile) MergeFirstWin(serviceULimitsNoFile *ServiceULimitsNoFile) {
	switch {
	case nf == nil && serviceULimitsNoFile == nil:
		fallthrough
	case nf != nil && serviceULimitsNoFile == nil:
		return

	// WARN: It's not possible to change the memory pointer nf *ServiceULimitsNoFile
	// to a new initialized ServiceULimitsNoFile without returning the serviceULimitsNoFile
	// it self.
	//
	// case nf == nil && serviceULimitsNoFile != nil:
	// 	nf = NewServiceULimitsNoFile()
	// 	fallthrough

	default:
		nf.mergeFirstWinHard(serviceULimitsNoFile.Hard)
		nf.mergeFirstWinSoft(serviceULimitsNoFile.Soft)
	}
}

// MergeLastWin merges adds or overwrite the attributes of the passed
// ServiceULimitsNoFile with the existing one.
func (nf *ServiceULimitsNoFile) MergeLastWin(serviceULimitsNoFile *ServiceULimitsNoFile) {
	switch {
	case nf == nil && serviceULimitsNoFile == nil:
		fallthrough
	case nf != nil && serviceULimitsNoFile == nil:
		return

	// WARN: It's not possible to change the memory pointer nf *ServiceULimitsNoFile
	// to a new initialized ServiceULimitsNoFile without returning the serviceULimitsNoFile
	// it self.
	//
	// case nf == nil && serviceULimitsNoFile != nil:
	// 	nf = NewServiceULimitsNoFile()
	// 	fallthrough

	default:
		nf.mergeLastWinHard(serviceULimitsNoFile.Hard)
		nf.mergeLastWinSoft(serviceULimitsNoFile.Soft)
	}
}

func (nf *ServiceULimitsNoFile) mergeFirstWinHard(hard uint) {
	if nf.Hard != hard {
		return
	}
	nf.Hard = hard
}

func (nf *ServiceULimitsNoFile) mergeFirstWinSoft(soft uint) {
	if nf.Soft != soft {
		return
	}
	nf.Soft = soft
}

func (nf *ServiceULimitsNoFile) mergeLastWinHard(hard uint) {
	if nf.Hard != hard {
		nf.Hard = hard
	}
}

func (nf *ServiceULimitsNoFile) mergeLastWinSoft(soft uint) {
	if nf.Soft != soft {
		nf.Soft = soft
	}
}

func NewServiceULimitsNoFile() *ServiceULimitsNoFile {
	return &ServiceULimitsNoFile{}
}

type Volume struct {
	External bool `json:"external,omitempty" yaml:"external,omitempty"`
}

// Equal returns true if the passed equalable is equal
func (v *Volume) Equal(equalable Equalable) bool {
	volume, ok := equalable.(*Volume)
	if !ok {
		return false
	}

	switch {
	case v == nil && volume == nil:
		return true
	case v != nil && volume == nil:
		fallthrough
	case v == nil && volume != nil:
		return false
	default:
		return v.External == volume.External
	}
}

// MergeFirstWin adds only the attributes of the passed Volume they are
// undefined.
func (v *Volume) MergeFirstWin(volume *Volume) {
	switch {
	case v == nil && volume == nil:
		fallthrough
	case v != nil && volume == nil:
		return

	// WARN: It's not possible to change the memory pointer v *Volume to a new
	// initialized Volume without returning the volume it self.
	//
	// case v == nil && volume != nil:
	// 	v = NewVolume()
	// 	fallthrough

	default:
		v.mergeFirstWinExternal(volume.External)
	}
}

func (v *Volume) MergeLastWin(volume *Volume) {
	switch {
	case v == nil && volume == nil:
		fallthrough
	case v != nil && volume == nil:
		return

	// WARN: It's not possible to change the memory pointer v *Volume to a new
	// initialized Volume without returning the volume it self.
	//
	// case v == nil && volume != nil:
	// 	v = NewVolume()
	// 	fallthrough

	default:
		v.mergeLastWinExternal(volume.External)
	}
}

func (v *Volume) mergeFirstWinExternal(external bool) {
	if v.External {
		return
	}
	v.External = true
}

func (v *Volume) mergeLastWinExternal(external bool) {
	if v.External != external {
		v.External = external
	}
}

func NewVolume() *Volume {
	return &Volume{
		External: false,
	}
}

// existsInSlice returns true when the passed comparable K exists in slice of
// comparables []K.
func existsInSlice[K comparable](comparables []K, k K) bool {
	for _, c := range comparables {
		if c == k {
			return true
		}
	}
	return false
}

func equalSlice[K comparable](sliceA []K, sliceB []K) bool {
	equalFunc := func(sliceA []K, sliceB []K) bool {
	LOOP:
		for i := range sliceA {
			for j := range sliceB {
				if sliceA[i] == sliceB[j] {
					continue LOOP
				}
			}
			return false
		}
		return true
	}

	return equalFunc(sliceA, sliceB) && equalFunc(sliceB, sliceA)
}

func splitStringInKeyValue(s, sep string) (string, string) {
	key := strings.Split(s, sep)[0]
	value := strings.TrimPrefix(s, fmt.Sprintf("%s%s", key, sep))
	return key, value
}

func splitStringInPort(s string) (string, string, string) {
	parts := strings.Split(s, portDelimiter)
	src := parts[0]
	rest := parts[1]

	parts = strings.Split(rest, portProtocolDelimiter)
	if len(parts) == 2 {
		return src, parts[0], parts[1]
	}

	return src, parts[0], ""
}

func splitStringInVolume(s string) (string, string, string) {
	parts := strings.Split(s, volumeDelimiter)
	src := parts[0]
	dest := parts[1]
	if len(parts) == 3 && len(parts[2]) > 0 {
		perm := parts[2]
		return src, dest, perm
	}
	return src, dest, ""
}
