package top

type taobaoMap map[string]interface{}

func (m taobaoMap) unwrap() (r interface{}, count string) {
	convertedMap := map[string]interface{}(m)
	return taobaoMap(convertedMap).unwrap(), count
}
