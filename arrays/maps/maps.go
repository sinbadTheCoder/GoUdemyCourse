package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func Main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	fmt.Println(websites["LinkedIn"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites["LinkedIn"])
	delete(websites, "LinkedIn")
	fmt.Println(websites["LinkedIn"])
}
