package read

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type User struct {
	Name       string
	Occupation string
}

func read() {
	log.Println("Starting yaml unmarshaller...")

	testYaml, err := ioutil.ReadFile("./yaml/test1.yaml")

	if err != nil {
		log.Error(err)
	}

	data := make(map[string]User)

	err2 := yaml.Unmarshal(testYaml, &data)

	if err2 != nil {
		log.Error(err2)
	}

	for k, v := range data {
		fmt.Printf("%s's name is %s, and they are a %s.\n", k, v.Name, v.Occupation)
	}
}
