package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
)

type MyTransaction struct {
	cost float64 `json:"cost" binding: "required"`
}

func health (c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"running":true})  
	fmt.Printf("replied to client health check.")
  }
func transaction (c *gin.Context){
	var json MyTransaction
	bool := c.BindJSON(&json) 
	if bool != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment, please make sure you use a decimal with no more than two decimal places."})
		return
	}
	points := calculateTransaction(json.cost)
	c.JSON(http.StatusCreated, gin.H{"reward": points})
	fmt.Printf("rewarded client %d  points \n", points )
	return
	//c.Data(200, "application/json; charset=utf-8", []byte(string(dat)))	
}

func calculateTransaction (cost float64)(int64) {
	var total = int64(((cost - 50) * 1) + ((cost - 100) * 1))
	
	if total >= 1 {
		return total
	}
	return 0
}

/* func formatCheck (){
	
	 if input == float64(int64(input)) {
		fmt.Println("passed data type test")
		return true
	} else {
		fmt.Println("failed data type test")
		return false
	} 
} */





func main() {
	
  r := gin.Default()
  r.GET("/api/health", health)
  r.POST("/api/transaction", transaction)
  
  r.Run(":8080")
}
