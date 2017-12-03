package intutils

const (
	MinUint uint = 0
	MaxUint uint = ^MinUint
	MaxInt  int  = int(MaxUint >> 1)
	MinInt  int  = ^MaxInt
)
