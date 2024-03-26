package tws

/**
 * string OR int64
 */
type id interface{}

type Node struct {
	RaterId  id
	TargetId id
	Deep     int64
	Score    int64
}
