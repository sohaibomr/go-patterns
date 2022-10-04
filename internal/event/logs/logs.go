package logs

import "fmt"

func HandleAuditLog(st string) error {
	fmt.Println(st)
	fmt.Println("Handling audit logs")
	return nil
}
