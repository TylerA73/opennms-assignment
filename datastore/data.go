package datastore

var (
	Data []interface{}
)

func init() {
	Data = make([]interface{}, 0)
}

func SaveStats(v interface{}) {
	Data = append(Data, v)
}

func GetStats() []interface{} {
	return Data
}
