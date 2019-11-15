package eather

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/EatherGo/eather/types"
)

// Module of type ModuleXML
type Module types.ModuleXML

var (
	sortedModules    []Module
	allModuleConfigs map[string]Module = make(map[string]Module)
	modConf          map[string]bool   = loadModuleConf()
)

// LoadModules will load all modules inside modules directory
func LoadModules(dirs []string) {
	for _, dir := range dirs {
		loadDir(dir)
	}

	orderModulesByPriorities()

	for _, m := range sortedModules {
		m.processModule()
	}
}

func loadDir(dir string) {
	files := getListOfModuleFolders(dir)

	getListOfModuleConfigs(files, dir)
}

func getListOfModuleFolders(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func getListOfModuleConfigs(files []os.FileInfo, dir string) {

	for _, f := range files {
		module, err := loadModule(f.Name(), dir)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if modConf[module.Name] {
			module.Dir = dir
			allModuleConfigs[module.Name] = module
		}
	}

	return
}

func orderModulesByPriorities() {
	var parents []string
	for _, m := range allModuleConfigs {
		m.addAllDependencies(parents)
	}
}

func callFunc(events map[string]types.EventFunc, name string) types.EventFunc {
	if val, ok := events[name]; ok {
		return val
	}

	return func(data ...interface{}) {}
}

func loadModule(name string, dir string) (module Module, err error) {

	file := fmt.Sprintf("%s/%s/etc/module.xml", dir, name)

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

	file := types.ConfigDir + "modules.xml"

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
