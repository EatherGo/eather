package eather

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"plugin"
)

// ModuleVersion struct - structure of moduleVersion in database
type ModuleVersion struct {
	ID uint `gorm:"primary_key"`
	DatabaseCreatedAt
	DatabaseUpdatedAt
	DatabaseDeletedAt
	Name    string
	Version string
}

// Module interface
type Module interface{}

// Eventable interface for modules that are with events
// add this func to module to enable events
type Eventable interface {
	GetEventFuncs() []Fire
}

// Installable interface for modules that are with func Install
// add this func to run install func during installation of module
type Installable interface {
	Install()
}

// Upgradable interface for modules that are with func Upgrade
// add this func to run Upgrade after upgrading module version
type Upgradable interface {
	Upgrade(version string)
}

// Routable interface for modules that are with func Routable
// add this func to map routes for module
type Routable interface {
	MapRoutes()
}

// Cronable interface for modules that are with func Crons
// return cronlist to add custom crons from modules to global list
type Cronable interface {
	Crons() CronList
}

// Callable interface for modules that are with GetPublicFuncs
// posibility to add custom public functions
type Callable interface {
	GetPublicFuncs() PublicFuncList
}

// PublicFunc function type for public function
type PublicFunc func(data ...interface{}) (interface{}, error)

// PublicFuncList is a list of function
type PublicFuncList map[string]PublicFunc

// Call function to call the public function of module
func (pfl PublicFuncList) Call(name string, data ...interface{}) (i interface{}, err error) {
	if pf, ok := pfl[name]; ok {
		i, err = pf(data)
		return
	}

	return
}

// InitVersion - initialize version with automigration
func InitVersion() {
	db.AutoMigrate(&ModuleVersion{})
}

// GetVersion - load version from database
func (m ModuleXML) GetVersion() string {
	module := ModuleVersion{}
	db.Select("version").Where("name = ?", m.Name).First(&module)

	return module.Version
}

// UpdateVersion - set the new version of the module to the database
func (m ModuleXML) UpdateVersion() {
	if m.GetVersion() == "" {
		db.Create(&ModuleVersion{Name: m.Name, Version: m.Version})
	} else {
		var module ModuleVersion
		db.Where("name = ?", m.Name).First(&module)

		db.Model(&module).Update("version", m.Version)
	}
}

func (m ModuleXML) getPath(inclFilename bool) string {
	path := m.Dir + "/" + m.Name + "/"

	if inclFilename {
		return path + "module.so"
	}

	return path
}

func (m ModuleXML) init() Module {
	plug, err := plugin.Open(m.getPath(true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lookup, err := plug.Lookup(m.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mod, _ := lookup.(func() (Module, error))()

	module, ok := mod.(Module)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	return module
}

func (m ModuleXML) addAllDependencies(parents []string) {
	var cParents []string
	copy(cParents, parents)

	if len(m.Dependencies.Dependencies) > 0 {
		for _, d := range m.Dependencies.Dependencies {
			if mod, ok := allModuleConfigs[d]; ok {
				if sliceContains(parents, m.Name) {
					panic("Module " + d + " got into loop. Take a look on dependencies")
				}

				cParents = append(cParents, mod.Name)
				mod.addAllDependencies(cParents)
			} else {
				panic("Module " + d + " is not installed")
			}
		}
	}

	if !containsModule(sortedModules, m) {
		sortedModules = append(sortedModules, m)
	}
}

func (m ModuleXML) processModule() {

	registry := GetRegistry()
	eventer := GetEvents()

	m.build()

	mod := m.init()

	if routableModule, isRoutable := mod.(Routable); isRoutable {
		routableModule.MapRoutes()
	}

	registry.Add(mod, m.Name)

	if eventableModule, isEventable := mod.(Eventable); isEventable {
		for _, listener := range m.Events.Listeners {
			eventer.Add(listener.For, callFunc(eventableModule.GetEventFuncs(), listener.Call), listener.Call, listener.Name)
		}
	}

	syncVersions(m, mod)

	fmt.Println("Module " + m.Name + " is running \n")
}

func (m ModuleXML) build() {
	fullPath := m.getPath(true)

	if os.Getenv("REBUILD") == "1" {
		rmcmd := exec.Command("rm", fullPath)
		rmcmd.Run()
	}

	if _, err := os.Stat(fullPath); err != nil {
		fmt.Println("Module " + m.Name + " is not builded. Building...")

		cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", fullPath, m.getPath(false))

		var errb bytes.Buffer
		cmd.Stderr = &errb

		err := cmd.Run()
		if err != nil {
			log.Fatal(errb.String())
		}

		fmt.Println("Module " + m.Name + " was builded")
	}
}
