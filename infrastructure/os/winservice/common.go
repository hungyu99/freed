package winservice

import "github.com/hungyu99/freed/infrastructure/config"

// ServiceDescription contains information about a service, needed to administer it
type ServiceDescription struct {
	Name        string
	DisplayName string
	Description string
}

// MainFunc specifies the signature of an application's main function to be able to run as a windows service
type MainFunc func(startedChan chan<- struct{}) error

// WinServiceMain is only invoked on Windows. It detects when freed is running
// as a service and reacts accordingly.
var WinServiceMain = func(MainFunc, *ServiceDescription, *config.Config) (bool, error) { return false, nil }
