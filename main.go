package main

import (
  "net/http"
  "fmt"
  "github.com/gin-gonic/gin"
)

type MyTransaction struct { // json struct
	Cost int64 `json:"cost" binding: "required"`
}

func health (c *gin.Context) { // monitor health of api
    c.JSON(http.StatusOK, gin.H{"running":true})  
	fmt.Printf("replied to client health check with success status.")
  }
func transaction (c *gin.Context){ // process customer transaction
	var myjson MyTransaction 
	if c.BindJSON(&myjson) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment, please make sure you use a integer with no decimal places."})
		fmt.Printf("invalid payment, user needs to instead use a integer with no decimal places.")
		return
	}
	fmt.Printf("sending cost to calculate function %d", myjson.Cost )
	points := calculateTransaction(myjson.Cost)
	c.JSON(http.StatusCreated, gin.H{"reward": points})
	fmt.Printf("sent reward client %d  points \n", points )
	return
	//c.Data(200, "application/json; charset=utf-8", []byte(string(dat)))	
}

func calculateTransaction (cost int64)(int64) { // calculate customer reward points
    fmt.Printf("cost before conversion %d \n", cost )
    //convertedCost, err := strconv.Atoi(cost)
	//fmt.Printf("convertedCost after conversion %d  \n", convertedCost )
	var first = ((cost - 50) * 1)
	var second =  ((cost - 100) * 1)
	if first < 0 {//account for less than 50 dollars spent
		first = 0
	}
	if second < 0{ // account for less than 100 dollars spent
		second = 0
	}
	var total = first + second
	fmt.Printf("rewarded client %d  total \n", total )
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





func main() { // setup router and api
	
  r := gin.Default()
  r.GET("/api/health", health)
  r.POST("/api/transaction", transaction)
  
  r.Run(":8080")
}
