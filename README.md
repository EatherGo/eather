# EATHER Project
Application that loading Go plugins that are created by predefined interface using built-in buildmode=plugin.  

## Install
Simply use the HelloWorld application or start fresh by running or get eather command line tools.

```
go get -u github.com/EatherGo/cmd/eather
```

## Create New Application
```
eather create -n NewApp
```

## Create New Module
Make sure that you are in your application directory.
```
eather module -n EmptyModule
```

This will create all necessary files for an empty module. New module will be stored to your env variable `CUSTOM_MODULES_DIR`

### With Controller 
```
eather module -n EmptyModule -c
```

### With Events
```
eather module -n EmptyModule -e
```

## Development

### Create module

Modules are loading automatically from folders defined in ENV CORE_MODULES_DIR and CUSTOM_MODULES_DIR. If you are using more than two directories for modules there is a posibility to add new modules directory through config. Each module need to be defined by `/etc/module.xml` in module directory and enabled in global `./config/modules.xml`.
Modules are builded using buildmode=plugin and module.so is loaded by application. Modules can be Eventable, Installable, Upgradable, Routable, Cronable, Callable described in module.go

#### Empty module
Create `etc` directory with module.xml 
```
<?xml version="1.0" encoding="UTF-8"?>
<module>
    <name>Empty</name>
    <version>1.0.0</version>
</module>

```
Name is used for defining a name for module and version for versioning.

Create `./config/modules.xml` and set module to enabled.
```
<?xml version="1.0" encoding="UTF-8"?>
<modules>
    <module>
        <name>Empty</name>
        <enabled>true</enabled>
    </module>
</modules>

```

Create main.go with function called same as name of module which should return module itself and error.
```
package main

import (
	"github.com/EatherGo/eather"
)

type module struct{}

// Empty to export in plugin
func Empty() (f eather.Module, err error) {
	f = module{}
	return
}

```

Run your application and you should see this output. 
```
Module Empty is not builded. Building...
Module Empty was builded
Module Empty is running 

```
Now your Empty module is running.

#### Evantable module
To make module Eventable add function GetEventFuncs() to your main.go
```
.
.
.

func (m module) GetEventFuncs() []eather.Fire {
	return eventFuncs
}

var eventFuncs = []eather.Fire{
	eather.Fire{Call: "added", Func: added},
	eather.Fire{Call: "removed", Func: removed},
}

var added = func(data ...interface{}) {
	fmt.Println(data)
	time.Sleep(2 * time.Second)
	fmt.Println("Running event after added")
}

var removed = func(data ...interface{}) {
	fmt.Println(data)
	time.Sleep(2 * time.Second)
	fmt.Println("Running event after removed")
}


.
.
.
```

And add to module.xml listeners for this events
```
<?xml version="1.0" encoding="UTF-8"?>
<module>
    <name>Empty</name>
    <version>1.0.0</version>
    <events>
        <listener for="test_added" call="added" name="add_some_stuff"></listener>
        <listener for="test_removed" call="removed" name="remove_some_stuff"></listener>
    </events>
</module>
```

Now your functions `added` can be triggered by calling `eather.GetEvents().Emmit("test_added", map[string]interface{}{"code": "test_code"})`

#### Installable module

To make module Installable you need to add function Install() to your module. This function will run only once when module is installing and is stored to database. Usually used for migrating to database.

```
.
.
.

func (m module) Install() {
	eather.GetDb().AutoMigrate(&YourModel{})
}

// YourModel struct
type YourModel struct {
	eather.ModelBase
}

.
.
.
```

#### Upgradable module

Add Upgrade(version) function to your module to make it Upgradable. This function is called every time you upgrade module version in `/etc/module.xml`. 
```
.
.
.

func (m module) Upgrade(version string) {
	
}

.
.
.
```

#### Routable module

Add function MapRoutes() to make module Routable.
```
.
.
.

func (m module) MapRoutes() {
	router := eather.GetRouter()

	router.HandleFunc("/", controllerEmpty).Methods("GET")
}

func controllerEmpty(w http.ResponseWriter, r *http.Request) {
	eather.SendJSONResponse(w, eather.Response{Message: "Empty controller"})
}

.
.
.
``` 
This will set controllerEmpty for path `/`.


#### Cronable module

Add function Crons() and module will be Cronable.
```
.
.
.

func (m module) Crons() eather.CronList {
	return eather.CronList{
		eather.Cron{Spec: "* * * * *", Cmd: func() { fmt.Println("test") }},
	}
}

.
.
.
``` 

This will run every second function Cmd and will print `test` into the terminal.

#### Callable module

Add function GetPublicFuncs() to make module Callable. This will allow to call you any function of your module in other parts of application or in another module.

```
.
.
.

func (m module) GetPublicFuncs() eather.PublicFuncList {
	list := make(eather.PublicFuncList)

	list["test"] = test

	return list
}

func test(data ...interface{}) (interface{}, error) {
	return []string{"testing public function of module"}, nil
}

.
.
.
```

Now it is possible to call this function from any part of application. It will print the return of test function to terminal.
```
    if callable := eather.GetRegistry().Get("Empty").GetCallable(); callable != nil {
		data, _ := callable.GetPublicFuncs().Call("test")

		fmt.Println(data)
	}
```
