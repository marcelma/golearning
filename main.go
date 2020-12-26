package main

import (
	"github.com/marcelma/golearning/array"
	"github.com/marcelma/golearning/concurrency"
	"github.com/marcelma/golearning/concurrency/channel"
	"github.com/marcelma/golearning/errorhandling"
	"github.com/marcelma/golearning/function"
	"github.com/marcelma/golearning/json"
	"github.com/marcelma/golearning/maps"
	"github.com/marcelma/golearning/pointer"
	"github.com/marcelma/golearning/slice"
	"github.com/marcelma/golearning/stdout"
	"github.com/marcelma/golearning/structe"
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
	slice.Examples()
	stdout.Examples()
	structe.Examples()
}
