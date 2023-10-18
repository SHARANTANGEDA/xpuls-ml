package types

type SortOrders string

const (
	ASC  SortOrders = "asc"
	DESC SortOrders = "desc"
)

type UniqueIdPrefixes string

const (
	PROMPT_ID_PREFIX  UniqueIdPrefixes = "xpr"
	PROJECT_ID_PREFIX UniqueIdPrefixes = "xproj-"
)
