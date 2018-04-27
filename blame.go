package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Nerdmaster/text-generator/pkg/filter/iafix"
	"github.com/Nerdmaster/text-generator/pkg/filter/substitution"
	"github.com/Nerdmaster/text-generator/pkg/filter/variation"
	"github.com/Nerdmaster/text-generator/pkg/generator"
	"github.com/Nerdmaster/text-generator/pkg/template"
	"github.com/go-yaml/yaml"
)

// EnterpriseBlameSystemVersionOnePointFivePointThree helps management teams
// adopt a customer-focused paradigm and stuff
type EnterpriseBlameSystemVersionOnePointFivePointThree struct {
	People []struct {
		Name    string
		Pronoun string
	}
	Apps    []string
	Reasons []string
	Formats []string
}

func main() {
	var data, err = ioutil.ReadFile("blame.yml")
	if err != nil {
		log.Fatalf("Unable to open blame.yml: %s", err)
	}

	var seed = time.Now().UTC().UnixNano()
	rand.Seed(seed)

	var ebs EnterpriseBlameSystemVersionOnePointFivePointThree
	err = yaml.Unmarshal(data, &ebs)
	if err != nil {
		log.Fatalf("Unable to parse yaml: %s", err)
	}
	var t = setupTemplate(ebs)

	fmt.Println(t.Execute())
}

func setupTemplate(ebs EnterpriseBlameSystemVersionOnePointFivePointThree) *template.Template {
	var t = template.FromString("{{format}}")
	var subFilter = substitution.New()

	// People are tricky; we actually just pull one and their pronoun rather than
	// having an actual random list, since we have to assign blame to a single
	// individual
	rand.Shuffle(len(ebs.People), func(i, j int) {
		ebs.People[i], ebs.People[j] = ebs.People[j], ebs.People[i]
	})
	subFilter.SetGenerator("person", &generator.Static{ebs.People[0].Name})
	subFilter.SetGenerator("pronoun", &generator.Static{ebs.People[0].Pronoun})

	var list = generator.MakeRandom()
	subFilter.SetGenerator("app", list)
	for _, str := range ebs.Apps {
		list.Append(strings.TrimSpace(str))
	}

	list = generator.MakeRandom()
	subFilter.SetGenerator("reason", list)
	for _, str := range ebs.Reasons {
		list.Append(strings.TrimSpace(str))
	}

	list = generator.MakeRandom()
	subFilter.SetGenerator("format", list)
	for _, str := range ebs.Formats {
		list.Append(strings.TrimSpace(str))
	}

	t.AddFilter(subFilter)
	t.AddFilter(variation.New())
	t.AddFilter(iafix.New())

	return t
}
