package lib

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/EatherGo/eather/lib/types"
)

// Module of type ModuleXML
type Module types.ModuleXML

var (
	sortedModules []Module
	modConf       map[string]bool
)

// LoadModules will load all modules inside modulesDir directory
func LoadModules() {
	files := getListOfModuleFolders()

	modConf = loadModuleConf()

	orderByPriorities(getListOfModuleConfigs(files))

	for _, m := range sortedModules {
		m.processModule()
	}
}

func getListOfModuleFolders() []os.FileInfo {
	files, err := ioutil.ReadDir(types.ModulesDir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func getListOfModuleConfigs(files []os.FileInfo) (moduleConfigs map[string]Module) {
	moduleConfigs = make(map[string]Module)

	for _, f := range files {
		module, err := loadModule(f.Name())

		if err != nil {
			fmt.Println(err)
			continue
		}

		if modConf[module.Name] {
			moduleConfigs[module.Name] = module
		}
	}

	return
}

func orderByPriorities(moduleConfigs map[string]Module) {
	var parents []string
	for _, m := range moduleConfigs {
		m.addAllDependencies(moduleConfigs, parents)
	}
}

func callFunc(events map[string]types.EventFunc, name string) types.EventFunc {
	if val, ok := events[name]; ok {
		return val
	}

	return func(data ...interface{}) {}
}

func loadModule(name string) (module Module, err error) {

	file := fmt.Sprintf("%s/%s/etc/module.xml", types.ModulesDir, name)

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

func loadModuleConf() (modConfigs map[string]bool) {

	file := "./app/modules.xml"

	xmlFile, err := os.Open(file)

	if err != nil {
		return
	}

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	module := types.ModulesConfigXML{}

	err = xml.Unmarshal(byteValue, &module)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	modConfigs = map[string]bool{}
	for _, c := range module.Modules {
		modConfigs[c.Name] = c.Enabled
	}

	return
}

func syncVersions(module Module, mod types.Module) {
	version := module.GetVersion()

	if version == "" {
		fmt.Println("Version not found. Installing " + module.Name + " version " + module.Version + "...")
		mod.Install()
		module.UpdateVersion()
		fmt.Println(module.Name + " was installed")
		return
	}

	if version != module.Version {
		fmt.Println("Upgrading " + module.Name + " to version " + module.Version + "...")
		module.UpdateVersion()
		mod.Upgrade(module.Version)
		fmt.Println(module.Name + " was upgraded")
	}
}
