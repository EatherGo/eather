# EATHER Project
Application that loading Go plugins using built-in Go plugins created by predefined interface.  

## Overview

### Getting started 
Simply use the HelloWorld application or start fresh by running

```
go get -u github.com/EatherGo/eather
```

Then initialize application and start server
```
package main

import (
	"github.com/EatherGo/eather"
)

func main() {
	config := eather.GetConfig()
    
	eather.Start(config)
}

```

Access http://localhost:8000/ or APP_URL from ENV.
