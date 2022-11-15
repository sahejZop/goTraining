package packageInsidePackage

import (
	"github.com/sahejZop/goTraining/models"
	"github.com/sahejZop/goTraining/models/addition"
)

func NestedCall() {
	models.CallInsidePackage()
	addition.Add(10, 20)
}
