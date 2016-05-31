package hello

import "fmt"

func HelloWorld(name string) string {

    if name != "" {
        return fmt.Sprintf("Hello, %s!", name)
    } else {
        return "Hello, World!"
    }
}
