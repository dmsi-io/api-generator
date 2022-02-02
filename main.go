package main

import (
	"os"

	"github.com/alehechka/api-generator/jenshared"
)

func main() {
	packageName := os.Args[1]
	fileName := os.Args[2]

	jenshared.CreateStructs(packageName, fileName)
}
