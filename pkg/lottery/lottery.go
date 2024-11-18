package lottery

// Number is the type of lottery number, currently ranging from 1 to 90
type Number = uint8

// PlayerID is the type that represents a sequential player ID. Currently, up to 10 million.
type PlayerID = int32

// MaxNumber is the maximum lottery number, inclusive
const MaxNumber = 90

// NumberBitSize the lottery number bit size, necessary for parsing textual input.
const NumberBitSize = 8

// NumPicks is the number of distinct lottery picks, currently 5 distinct numbers.
const NumPicks = 5
