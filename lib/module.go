package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"plugin"

	"github.com/EatherGo/eather/lib/types"

	"github.com/jinzhu/gorm"
)

// ModuleVersion struct - structure of moduleVersion in database
type ModuleVersion struct {
	gorm.Model
	Name    string
	Version string
}

// InitVersion - initialize version with automigration
func InitVersion() {
	db.AutoMigrate(&ModuleVersion{})
}

// GetVersion - load version from database
func (m Module) GetVersion() string {
	module := ModuleVersion{}
	db.Select("version").Where("name = ?", m.Name).First(&module)

	return module.Version
}

// UpdateVersion - set the new version of the module to the database
func (m Module) UpdateVersion() {
	if m.GetVersion() == "" {
		db.Create(&ModuleVersion{Name: m.Name, Version: m.Version})
	} else {
		var module ModuleVersion
		db.Where("name = ?", m.Name).First(&module)

		db.Model(&module).Update("version", m.Version)
	}
}

func (m Module) getPath(inclFilename bool) string {

	path := types.ModulesDir + "/" + m.Name + "/"

	if inclFilename {
		return path + "module.so"
	}

	return path
}

func (m Module) init() types.Module {
	plug, err := plugin.Open(m.getPath(true))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lookup, err := plug.Lookup(m.Func)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mod, _ := lookup.(func() (types.Module, error))()

	module, ok := mod.(types.Module)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	return module
}

func (m Module) addAllDependencies(moduleConfigs map[string]Module, parents []string) {
	var cParents []string
	copy(cParents, parents)

	if len(m.Dependencies.Dependencies) > 0 {
		for _, d := range m.Dependencies.Dependencies {
			if mod, ok := moduleConfigs[d]; ok {
				if sliceContains(parents, m.Name) {
					panic("Module " + d + " got into loop. Take a look on dependencies")
				}

				cParents = append(cParents, mod.Name)
				mod.addAllDependencies(moduleConfigs, cParents)
			} else {
				panic("Module" + d + " not set")
			}
		}
	}

	if !contains(sortedModules, m) {
		sortedModules = append(sortedModules, m)
	}
}

func (m Module) processModule() {

	registry := GetRegistry()
	eventer := GetEvents()

	m.build()

	mod := m.init()
	mod.MapRoutes()

	registry.Add(mod, m.Name)

	for _, listener := range m.Events.Listeners {
		eventer.Add(listener.For, callFunc(mod.GetEventFuncs(), listener.Call), listener.Name)
	}

	syncVersions(m, mod)
}

func (m Module) build() {
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
