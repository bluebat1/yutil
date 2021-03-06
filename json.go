package yutil

import "encoding/json"

// 将对象转成json字符串
func JsonObjToStr(v interface{}) string{
	bts,err := json.Marshal(v)
	_conf.log("json","JsonObjToStr",err)
	return string(bts)
}

// 将对象转成[]byte
func JsonObjToBytes(v interface{}) []byte{
	bts,err := json.Marshal(v)
	_conf.log("json","JsonObjToBytes",err)
	return bts
}

// 将[]byte转成对象
func JsonBytesToObj(bts []byte,v interface{}){
	err := json.Unmarshal(bts,v)
	_conf.log("json","JsonBytesToObj",err)
}

// 将json字符串转成对象
func JsonStrToObj(s string,v interface{}){
	err := json.Unmarshal([]byte(s),v)
	_conf.log("json","JsonStrToObj",err)
}

