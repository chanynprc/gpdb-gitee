package config

import (
	"github.com/greenplum-db/gpdb/gp/internal/enums"
)

type InfraConfig interface {
	GetRequestPort() int
	GetPublishPort() int
	GetCoordinator() HostConfig
	GetStandby() HostConfig
	GetSegmentHost() SegmentHostsConfig
}
type Infra struct {
	RequestPort  int           `json:"requestPort"`
	PublishPort  int           `json:"publishPort"`
	Coordinator  *Host         `json:"coordinatorHost" mapstructure:"coordinatorHost"`
	Standby      *Host         `json:"standbyHost" mapstructure:"standbyHost"`
	SegmentHosts *SegmentHosts `json:"segmentHost" mapstructure:"segmentHost"`
}

func (i Infra) GetRequestPort() int {
	return i.RequestPort
}
func (i Infra) GetPublishPort() int {
	return i.PublishPort
}
func (i Infra) GetCoordinator() HostConfig {
	return i.Coordinator
}
func (i Infra) GetStandby() HostConfig {
	return i.Standby
}
func (i Infra) GetSegmentHost() SegmentHostsConfig {
	return i.SegmentHosts
}

type AuthenticationConfig interface {
	GetType() enums.AuthType
	GetPassword() string
}

type Authentication struct {
	Type     enums.AuthType `json:"type"`
	Password string         `json:"password"`
}

func (a Authentication) GetType() enums.AuthType {
	return a.Type
}
func (a Authentication) GetPassword() string {
	return a.Password
}

type HostConfig interface {
	GetIp() string
	GetHostname() string
	GetDomainName() string
	GetAuth() AuthenticationConfig
}

type Host struct {
	Hostname   string          `json:"hostname"`
	DomainName string          `json:"domainName"`
	Auth       *Authentication `json:"authentication"  mapstructure:"authentication"`
	Ip         string          `json:"ip"`
}

func (h Host) GetIp() string {
	return h.Ip
}
func (h Host) GetHostname() string {
	return h.Hostname
}
func (h Host) GetDomainName() string {
	return h.DomainName
}
func (h Host) GetAuth() AuthenticationConfig {
	return h.Auth
}

type SegmentHostsConfig interface {
	GetSegmentHostsCount() int
	GetNetwork() SegmentHostsNetworkConfig
	GetAuthentication() AuthenticationConfig
	GetHostnamePrefix() string
	GetDomainName() string
}
type SegmentHosts struct {
	SegmentHostsCount int                  `json:"segmentHostsCount"`
	Network           *SegmentHostsNetwork `json:"network"`
	Authentication    *Authentication      `json:"authentication"`
	HostnamePrefix    string               `json:"hostnamePrefix"`
	DomainName        string               `json:"domainName"`
}

func (s SegmentHosts) GetSegmentHostsCount() int {
	return s.SegmentHostsCount
}
func (s SegmentHosts) GetNetwork() SegmentHostsNetworkConfig {
	return s.Network
}
func (s SegmentHosts) GetAuthentication() AuthenticationConfig {
	return s.Authentication
}
func (s SegmentHosts) GetHostnamePrefix() string {
	return s.HostnamePrefix
}
func (s SegmentHosts) GetDomainName() string {
	return s.DomainName
}

type SegmentHostsNetworkConfig interface {
	GetInternalCidr() string
	GetIpRange() IpRangeConfig
	GetIpList() []string
}

type SegmentHostsNetwork struct {
	InternalCidr string   `json:"internalCidr"`
	IpRange      *IpRange `json:"ipRange"`
	IpList       []string `json:"ipList"`
}

func (sn SegmentHostsNetwork) GetInternalCidr() string {
	return sn.InternalCidr
}
func (sn SegmentHostsNetwork) GetIpRange() IpRangeConfig {
	return sn.IpRange
}
func (sn SegmentHostsNetwork) GetIpList() []string {
	return sn.IpList
}

type IpRangeConfig interface {
	GetFirstIp() string
	GetLastIp() string
}
type IpRange struct {
	FirstIp string `json:"first" mapstructure:"first"`
	LastIp  string `json:"last" mapstructure:"last"`
}

func (ip IpRange) GetFirstIp() string {
	return ip.FirstIp
}
func (ip IpRange) GetLastIp() string {
	return ip.LastIp
}
