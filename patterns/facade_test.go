package patterns

import (
	"fmt"
	"testing"
)

func printEvent(n interface{}) {
	fmt.Println(n)
}

func TestFacade(t *testing.T) {
	drive := Drive{}
	power := Power{}
	notification := Notification{}
	microwave := Microwave{&drive, &power, &notification}

	drive.AddHandler(printEvent)
	power.AddHandler(printEvent)
	notification.AddHandler(printEvent)

	microwave.Defrost()
}
