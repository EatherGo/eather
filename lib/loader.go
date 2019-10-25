package lib

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"project/lib/interfaces"
)

// Module of type ModuleXML
type Module interfaces.ModuleXML

var sortedModules []Module

// LoadModules will load all modules inside modulesDir directory
func LoadModules() {
	files := GetListOfModuleFolders()

	orderByPriorities(getLlistOfModuleConfigs(files))

	for _, m := range sortedModules {
		m.processModule()
	}
}

func GetListOfModuleFolders() []os.FileInfo {
	files, err := ioutil.ReadDir(interfaces.ModulesDir)
	if err != nil {
		log.Fatal(err)
	}

	return files
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

func callFunc(events map[string]interfaces.EventFunc, name string) interfaces.EventFunc {
	if val, ok := events[name]; ok {
		return val
	}

	return func(data ...interface{}) {}
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

func syncVersions(module Module, mod interfaces.Module) {
	version := module.GetVersion()

	if version == "" {
		fmt.Println("Version not found. Installing " + module.Name + " version " + module.Version + "...")
		mod.Install()
		module.UpdateVersion()
		return
	}

	if version != module.Version {
		fmt.Println("Upgrading " + module.Name + " to version " + module.Version + "...")
		module.UpdateVersion()
		mod.Upgrade(module.Version)
	}
}
