package mongo

type Connection struct {
	MongoDriver string
}

func NewConnection(md string) *Connection {
	return &Connection{MongoDriver: md}
}
