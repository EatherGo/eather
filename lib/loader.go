package lib

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"plugin"
	"project/lib/interfaces"
	"reflect"
)

// Module of type ModuleXML
type Module interfaces.ModuleXML

var sortedModules []Module

// LoadModules will load all modules inside modulesDir directory
func LoadModules() {
	files, err := ioutil.ReadDir(interfaces.ModulesDir)
	if err != nil {
		log.Fatal(err)
	}

	orderByPriorities(getLlistOfModuleConfigs(files))

	for _, m := range sortedModules {
		m.processModule()
	}
}

func getLlistOfModuleConfigs(files []os.FileInfo) (moduleConfigs map[string]Module) {
	moduleConfigs = make(map[string]Module)

	for _, f := range files {
		module, err := loadModule(f.Name())

		if err != nil {
			fmt.Println(err)
			continue
		}

		moduleConfigs[module.Name] = module
	}

	return
}

func orderByPriorities(moduleConfigs map[string]Module) {
	var parents []string
	for _, m := range moduleConfigs {
		m.addAllDependencies(moduleConfigs, parents)
	}
}

func callFunc(events map[string]func(), name string) func() {
	if val, ok := events[name]; ok {
		return val
	}

	return func() {}
}

// loadModule will load module by name
func loadModule(name string) (module Module, err error) {

	file := fmt.Sprintf("%s/%s/etc/module.xml", interfaces.ModulesDir, name)

	xmlFile, err := os.Open(file)

	if err != nil {
		return
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	err = xml.Unmarshal(byteValue, &module)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}

func (m Module) getPath(inclFilename bool) string {

	path := interfaces.ModulesDir + "/" + m.Name + "/"

	if inclFilename {
		return path + "module.so"
	}

	return path
}

func (m Module) init() interfaces.Module {
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

	mod, _ := lookup.(func() (interfaces.Module, error))()

	module, ok := mod.(interfaces.Module)
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

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Module " + m.Name + " was builded")
	}
}

func syncVersions(module Module, mod interfaces.Module) {
	version := GetVersion(module.Name)

	if version == "" {
		fmt.Println("Version not found. Installing " + module.Name + " version " + module.Version + "...")
		mod.Install()
		SetVersion(module.Name, module.Version)
		return
	}

	if version != module.Version {
		fmt.Println("Upgrading " + module.Name + " to version " + module.Version + "...")
		SetVersion(module.Name, module.Version)
		mod.Upgrade(module.Version)
	}
}

func contains(s []Module, e Module) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}
