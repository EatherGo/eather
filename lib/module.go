package lib

import "github.com/jinzhu/gorm"

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
func GetVersion(name string) string {
	module := ModuleVersion{}
	db.Select("version").Where("name = ?", name).First(&module)

	return module.Version
}

// SetVersion - set the new version of the module to the database
func SetVersion(name string, version string) {
	if GetVersion(name) == "" {
		db.Create(&ModuleVersion{Name: name, Version: version})
	} else {
		var module ModuleVersion
		db.Where("name = ?", name).First(&module)

		db.Model(&module).Update("version", version)
	}
}
