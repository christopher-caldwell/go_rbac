package session

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		log.Fatalf("Failed to enforce: %v", err)
	}

	if ok {
		fmt.Println("permit alice to read data1")
	} else {
		fmt.Println("deny the request, show an error")
	}
}
