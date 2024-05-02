package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/xmlpath/path"
)

func searchUsers(xmlFile string, username string) ([]string, error) {
	data, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return nil, err
	}

	expr := fmt.Sprintf("//user[username='%s']", username)
	root, err := path.Parse(string(data))
	if err != nil {
		return nil, err
	}

	var results []string
	iter := root.Iter("//user")
	for node := iter.Next(); node != nil; node = iter.Next() {
		username, ok := node.SelectElements("//username")[0].String()
		if ok && username == username {
			results = append(results, username)
		}
	}

	return results, nil
}

func main() {
	username := "admin"
	users, err := searchUsers("users.xml", username)
	if err != nil {
		fmt.Println("Error searching users:", err)
		return
	}

	if len(users) > 0 {
		fmt.Println("Found users:", strings.Join(users, ", "))
	} else {
		fmt.Println("User not found.")
	}
}
