package interactions_test

import (
	"testing"

	"github.com/Marcelixoo/go_specs_greet/domain/interactions"
	"github.com/Marcelixoo/go_specs_greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
