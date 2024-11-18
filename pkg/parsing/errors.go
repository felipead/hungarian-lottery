package parsing

import "errors"

var ErrInvalidQuantityOfNumbers = errors.New("invalid quantity of picked numbers")

var ErrNumberOutOfRange = errors.New("picked number is out of range")

var ErrNoRepeatedNumbers = errors.New("no repeated numbers should be picked")
