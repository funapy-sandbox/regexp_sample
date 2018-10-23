package main

import (
	"fmt"
	"regexp"
)

var html = `
<html>
<head>
  <meta charset="utf-8">
  <title>タイトル</title>
</head>
<body>
  <p><br>test1
  <br>test2
  <BR>test3
  </p>
</body>
</html>
`

func main() {
	r := regexp.MustCompile("<(?:br|BR)>(.*)")
	for _, matched := range r.FindAllStringSubmatch(html, -1) {
		if len(matched) > 1 {
			fmt.Println(matched[1])
		}
	}
}
