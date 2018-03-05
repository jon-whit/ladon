package ladon

import "github.com/ory/ladon"

/*
BooleanCondition is used to determine if a boolean context matches an expected
boolean condition.

BooleanCondition implements the ladon.Condition interface.
See https://github.com/ory/ladon/blob/master/condition.go
*/
type BooleanCondition struct {
	BooleanValue bool `json:"value"`
}

// GetName returns the name of the BooleanCondition
func (c *BooleanCondition) GetName() string {
	return "BooleanCondition"
}

// Fulfills determines if the BooleanCondition is fulfilled.
// The BooleanCondition is fulfilled if the provided boolean value matches the conditions boolean value.
func (c *BooleanCondition) Fulfills(value interface{}, _ *ladon.Request) bool {
	val, ok := value.(bool)

	return ok && val == c.BooleanValue
}
