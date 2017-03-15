package weightrandom

type Items struct {
	Weight int
	Data   interface{}
}

func NewItem(weight int, data interface{}) *Items {
	return &Items{
		Weight: weight,
		Data:   data,
	}
}
