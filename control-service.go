package main

import ("go.bug.st/serial.v1"
"github.com/gin-gonic/gin"
"net/http"
"os"
"fmt")

type changeRequest struct {
	Input int `json: "input"`
	Output int `json: "output"`
}

var port serial.Port

func putOutput(c *gin.Context) {
	var request changeRequest
	if err := c.BindJSON(&request); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request)
	command := fmt.Sprintf("V%d%dD\r", request.Output, request.Input);
	port.Write([]byte(command))
	c.IndentedJSON(http.StatusOK, "OK")
}

func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println(err)
	}
	if len(ports) == 0 {
		fmt.Printf("No serial ports found. Requires at least 1 serial port.")
		os.Exit(1)
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
	mode := &serial.Mode{}
	port, err = serial.Open(ports[0],mode)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer port.Close()

	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.PUT("/video", putOutput)
	router.Run(":8080")
}
