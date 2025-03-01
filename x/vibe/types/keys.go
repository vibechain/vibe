package types

const (
	// ModuleName defines the module name
	ModuleName = "vibe"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vibe"
)

var (
	ParamsKey = []byte("p_vibe")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
