package core

import (
	"errors"
	"sync"
)

type DatasourceFactoryRegistry interface {
	RegisterDatasource(name string, factory func() Datasource)
	GetDatasource(name string) (Datasource, error)
}

type datasourceFactoryRegistry struct {
	factoryRegistry      map[string]func() Datasource
	registryFactoryMutex sync.RWMutex
}

var _ DatasourceFactoryRegistry = &datasourceFactoryRegistry{}

var registryInstance *datasourceFactoryRegistry
var once sync.Once

func GetDatasourceFactoryRegistry() DatasourceFactoryRegistry {
	once.Do(func() {
		registryInstance = &datasourceFactoryRegistry{
			factoryRegistry: make(map[string]func() Datasource),
		}
	})
	return registryInstance
}

func (dfr *datasourceFactoryRegistry) RegisterDatasource(name string, factory func() Datasource) {
	dfr.registryFactoryMutex.Lock()
	defer dfr.registryFactoryMutex.Unlock()
	dfr.factoryRegistry[name] = factory
}

func (dfr *datasourceFactoryRegistry) GetDatasource(name string) (Datasource, error) {
	dfr.registryFactoryMutex.RLock()
	defer dfr.registryFactoryMutex.RUnlock()
	factory, exists := dfr.factoryRegistry[name]
	if !exists {
		return nil, errors.New("datasource not found")
	}
	return factory(), nil
}
