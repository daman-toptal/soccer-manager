package id

type IDPrefix string

var uuidRE = "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"

const (
	IDPrefixNone = IDPrefix("")
	IDPrefixRequest = IDPrefix("req-")
	IDPrefixTeam        = IDPrefix("tea-")
	IDPrefixUser        = IDPrefix("usr-")
	IDPrefixPlayer      = IDPrefix("ply-")
	IDPrefixTransaction = IDPrefix("txn-")
)

func (pr IDPrefix) String() string {
	return string(pr)
}

func (pr IDPrefix) REMatch() string {
	return string(pr) + uuidRE
}

type ID interface {
	Prefix() IDPrefix
	String() string
	UUIDString() string
	IsZero() bool
	JSONString() string
}
