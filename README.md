# microconfig
go micro config parser

## Features

- Tested and used in production
- Zero memory allocations
- KISS principle "Keep it simple, stupid"
- No dependencies


## Installing

1. Get package:

	```Shell
	go get github.com/microcodebase/microconfig
	```

2. Import it in your code:

	```Go
	import "github.com/microcodebase/microconfig"
	```

## Usage

1. Load config file:

	```Go
	conf, err := microconfig.ParseFile("zzz.conf")
	```

2. Parse config from string or []byte:

	```Go
	conf := microconfig.Parse([]byte("a = 1"))
	```

3. Config are read to string map:

	```Go
	var conf map[string]string = microconfig.Parse([]byte("key = value"))
	value := conf["key"]
	```

### Configuration example

```
# a config comment

# support the name = value format
Port=8000

# support spaces
  key1 = some value
  key2 = some other value
```

