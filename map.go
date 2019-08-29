
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var i int
	fmt.Scan(&i)
	
    m := make(map[string]int)
	for j := 0; j < i; j++ {
		scanner.Scan()
		text := scanner.Text()
		s := strings.Split(text, " ")
		
		phoneNumber,_ := strconv.ParseInt(s[1], 10, 0)

		m[s[0]] = int(phoneNumber)
	}

	for k := 0; k < i; k++ {

	}
}