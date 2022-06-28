package policy

import (
	"context"
	"time"

	"github.com/open-policy-agent/opa/sdk"
)

type Court struct {
	law   Law
	judge *sdk.OPA
}

type Law struct {
}

type Witness struct {
}

func New() *Court {
	return nil
}

func (c *Court) Decide(ctx context.Context, input map[string]interface{}) (bool, error) {
	result, err := c.judge.Decision(ctx, sdk.DecisionOptions{
		Now:   time.Now(),
		Path:  "/system/main/allow",
		Input: input,
	})
	if err != nil {
		return false, err
	}
	_ = result
	return false, nil
}
