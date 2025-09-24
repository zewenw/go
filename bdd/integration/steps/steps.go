package steps

import (
	"bdd"
	"fmt"
	"github.com/cucumber/godog"
)

var calc bdd.Calculator

func theFirstNumberIs(arg1 int) error {
	calc.A = arg1
	return nil
}

func theSecondNumberIs(arg1 int) error {
	calc.B = arg1
	return nil
}

func theNumbersAreAdded() error {
	calc.Add()
	return nil
}

func theResultShouldBe(expected int) error {
	if calc.Result != expected {
		return fmt.Errorf("expected %d but got %d", expected, calc.Result)
	}
	return nil
}

func InitializeAddScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the first number is (\d+)$`, theFirstNumberIs)
	ctx.Step(`^the second number is (\d+)$`, theSecondNumberIs)
	ctx.Step(`^the numbers are added$`, theNumbersAreAdded)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
}
