package tick_test

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	pipe, _, query := BatchQuery("select cpu_usage from cpu")
	logger := query.Log()
	logger.Level = "ERROR"
	logger.Prefix = "oh no"
	logger.Log() // default options

	got, err := PipelineTick(pipe)
	if err != nil {
		t.Fatalf("Unexpected error building pipeline %v", err)
	}
	want := `batch
    |query('select cpu_usage from cpu')
    |log()
        .level('ERROR')
        .prefix('oh no')
    |log()
        .level('INFO')
`
	if got != want {
		t.Errorf("TestQuery = %v, want %v", got, want)
		fmt.Println(got) // print is helpful to get the correct format.
	}
}