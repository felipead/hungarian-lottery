package parsing

import "errors"

var ErrInvalidQuantityOfPicks = errors.New("invalid quantity of lottery picks")
var ErrPickedNumberOutOfRange = errors.New("picked lottery number is out of range")
