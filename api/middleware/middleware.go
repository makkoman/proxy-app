package middleware

import (
	"bufio"
	"github.com/kataras/iris"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Queue
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que declaration
var Que []string

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var holder []string
	var queueArray []*Queue
	for scanner.Scan() {
		holder = append(holder, scanner.Text())
		if scanner.Text() == "" {
			w, _ := strconv.Atoi(strings.Split(holder[1], ":")[1])
			p, _ := strconv.Atoi(strings.Split(holder[2], ":")[1])
			queueArray = append(queueArray, &Queue{Domain: holder[0], Weight: w, Priority: p})
			holder = []string{}
			continue
		}
	}
	return queueArray
}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var low []string
	var medium []string
	var high []string
	lowD, mediumD, highD := prioritize(domain)
	low = append(low, lowD...)
	medium = append(medium, mediumD...)
	high = append(high, highD...)
	for _, queuedDomain := range Que{
		lowQ, mediumQ, highQ :=  prioritize(queuedDomain)
		low = append(low, lowQ...)
		medium = append(medium, mediumQ...)
		high = append(high, highQ...)
	}
	var newQueue []string
	newQueue = append(newQueue, high...)
	newQueue = append(newQueue, medium...)
	newQueue = append(newQueue, low...)
	Que = newQueue
	c.Next()
}


func prioritize(domain string) ([]string, []string, []string) {
	var low []string
	var medium []string
	var high []string
	var repo Repository
	repo = &Queue{}
	for _, row := range repo.Read() {
		// low
		if domain == row.Domain && (row.Priority < 5 || row.Weight < 5) {
			low = append(low, domain)
		}
		// med
		if domain == row.Domain && ((row.Priority > 5 && row.Weight < 5) || (row.Weight > 5 && row.Priority < 5) || (row.Weight == 5 && row.Priority == 5)) {
			medium = append(medium, domain)
		}
		// high
		if domain == row.Domain && (row.Priority > 5 && row.Weight > 5) {
			high = append(high, domain)
		}
	}
	return low, medium, high
}
