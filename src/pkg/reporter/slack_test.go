package reporter

import (
	"storage/src/config"
	"testing"

	"github.com/joho/godotenv"
)

func TestReport(t *testing.T) {
	// print working directory here
	godotenv.Load("../../../.env")
	config.Setup("../../config")

	tests := []struct {
		testName string
		title    string
		text     string
	}{
		{
			testName: "Good report",
			title:    "Happy message",
			text:     "This is regular happy message",
		},
		{
			testName: "Warning report",
			title:    "Concerning message",
			text:     "This is concerning, but not very urgent",
		},
		{
			testName: "Danger report",
			title:    "Crictial message",
			text:     "This is very high priority message",
		},
	}
	t.Run(tests[0].testName, func(t *testing.T) {
		Info(tests[0].title, tests[0].text)
	})
	t.Run(tests[1].testName, func(t *testing.T) {
		Warning(tests[1].title, tests[1].text)
	})
	t.Run(tests[2].testName, func(t *testing.T) {
		Error(tests[2].title, tests[2].text)
	})

}
