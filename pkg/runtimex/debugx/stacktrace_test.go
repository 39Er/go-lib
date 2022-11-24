package debugx

import (
	"fmt"
	"testing"
)

func TestGetStackTrace(t *testing.T) {
	fmt.Println(GetStackTrace(0))
	fmt.Println("------")
	fmt.Println(GetStackTrace(2))
	fmt.Println("------")
	fmt.Println(LocatePanic())
}
