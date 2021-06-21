package main

import (
	"github.com/marcelma/golearning/basic/array"
	"github.com/marcelma/golearning/basic/concurrency"
	"github.com/marcelma/golearning/basic/concurrency/channel"
	"github.com/marcelma/golearning/basic/errorhandling"
	"github.com/marcelma/golearning/basic/function"
	"github.com/marcelma/golearning/basic/json"
	"github.com/marcelma/golearning/basic/maps"
	"github.com/marcelma/golearning/basic/pointer"
	"github.com/marcelma/golearning/basic/slice"
	"github.com/marcelma/golearning/basic/stdout"
	"github.com/marcelma/golearning/basic/structe"
	"github.com/marcelma/golearning/basic/testing"
)

func main() {
	array.Examples()
	channel.Examples()
	concurrency.Examples()
	errorhandling.Examples()
	function.Examples()
	json.Examples()
	maps.Examples()
	pointer.Examples()
	testing.Examples()
	slice.Examples()
	stdout.Examples()
	structe.Examples()
}
