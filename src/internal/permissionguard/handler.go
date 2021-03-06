package permissionguard

import (
	"fmt"
	"os"
)

// Process takes action on a Conclusion
func Process(conclusion Conclusion) {
	exitCode := 0
	linebreak := "======================================================================"
	fmt.Println(linebreak)
	switch result := conclusion.Result; result {
	case Pass:
		fmt.Println("No permission is changed.")
	case Fail:
		more := conclusion.More
		less := conclusion.Less
		fmt.Println("Failure!")
		fmt.Println(fmt.Sprintf("\n%d new permission(s) added:", len(more)))
		for i := range more {
			fmt.Println(fmt.Sprintf("    %s", more[i].Name))
		}
		if len(less) > 0 {
			fmt.Println(fmt.Sprintf("\n%d old permission(s) removed:", len(less)))
			for i := range less {
				fmt.Println(fmt.Sprintf("    %s", less[i].Name))
			}
		}
		exitCode = 1
	case Warn:
		less := conclusion.Less
		fmt.Println("Warning!")
		fmt.Println(fmt.Sprintf("\n%d old permission(s) removed:", len(less)))
		for i := range less {
			fmt.Println(fmt.Sprintf("    %s", less[i].Name))
		}
		fmt.Println("\nA new snapshot needs to be taken.")
		exitCode = 1
	}
	fmt.Println(linebreak)
	os.Exit(exitCode)
}

// TakeSnapshot saves a slice of permissions to a file (JSON-encoded)
func TakeSnapshot(permissions []Permission, file string) {
	fileExist := true
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fileExist = false
	}
	PersistOntoDisk(ToJsonFromList(permissions), file)
	if fileExist {
		fmt.Println("Snapshot file has been updated.")
	} else {
		fmt.Println("Snapshot file has been generated.")
	}
}
