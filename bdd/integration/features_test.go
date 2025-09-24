package integration

import (
	"bdd/integration/steps"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

var scenarioInitializers = []func(ctx *godog.ScenarioContext){
	steps.InitializeAddScenario,
	// Add more step initializers here as needed
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
		Output: colors.Colored(os.Stdout),
	}

	status := godog.TestSuite{
		Name: "calculator",
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			for _, initFn := range scenarioInitializers {
				initFn(ctx)
			}
		},
		Options: &opts,
	}.Run()
	os.Exit(status)
}
