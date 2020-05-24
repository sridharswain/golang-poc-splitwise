package main

import (
	"splitwise-gateway/service/user"

	"github.com/gin-gonic/gin"
)

/* TODO
1. Add Cluster of server of service
2. Check if multiple concurrent forked processes can run sharing same port
*/
func main() {
	r := gin.Default()
	user.Print()
	StartRouter(r)
	r.Run(":8080")
}
