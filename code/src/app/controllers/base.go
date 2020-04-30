package controllers

func sendRes(code int, msg string, data interface{}) (m map[string]interface{}){
	m = map[string]interface{}{}
	m["code"] = code
	m["msg"] = msg
	if data == nil {
		m["data"] = []interface{}{}//map[string]interface{}{}
	} else {
		m["data"] = data
	}

	return
}
