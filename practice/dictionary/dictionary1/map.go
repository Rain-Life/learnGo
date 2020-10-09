package dictionary

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

func (d *Dictionary) Get(key interface{}) interface{}  {
	return d.data[key]
}

func (d *Dictionary) Visit(callback func(k,v interface{}) bool) {
	if callback == nil {
		return
	}

	for k,v := range d.data {
		if !callback(k,v) {
			return
		}
	}
}

func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()

	return d
}