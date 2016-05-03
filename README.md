## What is Configo

Configo is the simplist JSON config file reader for Go.

## Code Example

You can access values with Get...() functions & dot(.) seperated path string.

```
    config, err := configo.NewConfigFromFile("config.json")
    if err != nil {
        // handle error
    }

    value := config.GetString("path.to.the.key", "default value")
```

## Installation

go get github.com/hongseokyoon/configo