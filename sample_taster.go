package sample_taster

import "github.com/dev-cycles/contextive-demo-go-common"
import "fmt"

func Taste() string {
	return fmt.Sprintf("The sample tastes like: '%s'", sample.Sample())
}