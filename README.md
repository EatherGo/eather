# EATHER Project
Application that loading Go plugins using built-in Go plugins created by predefined interface.  

## Overview

### Getting started 
Simply use the HelloWorld application or start fresh by running

```
go get -u github.com/EatherGo/eather
```

Then start server
```
package main

import (
	"github.com/EatherGo/eather"
)

func main() {
	eather.Start(nil)
}

```

Copy .env.example to your project as .env with your own settings.

Access http://localhost:8000/ or APP_URL from ENV, and you should see 404 page not found or `Hello world` and you ready to start.