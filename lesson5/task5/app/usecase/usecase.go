package usecase

import "app/models"

type Usecase interface {
	CreateProxy(proxy *models.Nodes) error
	FindProxyByHostname(proxy *models.Nodes, hostname string) error
	Get(addr, path string, target interface{}) error
	RebootModem(addr, path string) error
	FindNodes(nodes *[]models.Nodes) error
}
