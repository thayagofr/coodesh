package utils

type Status int

const (
	TRASH Status = iota
	DRAFT
	PUBLISHED
)

func GetStatus(status Status) string {
	strings := [...]string{"trash", "draft", "published"}
	return strings[status]
}

type Collection int

const (
	HISTORY Collection = iota
	PRODUCTS
)

func GetCollection(collection Collection) string {
	strings := [...]string{"history", "products"}
	return strings[collection]
}
