package api

import (
	"strings"
	"bytes"
	"strconv"
)
func NeedUrlEncode(src string) bool {
	for i:=0;i<len([]byte(src));i++{
		b:=[]byte(src)[i]
		if (int(b) >= int('a') && int(b) <= int('z')) ||
			(int(b) >= int('A') && int(b) <= int('Z')) ||
			(int(b) >= int('0') && int(b) <= int('9')) ||
			strings.Contains("-_.!~*'()", string(b)) || b == '+'{
			continue
		}else if b=='%' &&( ([]byte(src)[i+1]>='0' &&[]byte(src)[i+1]<='9') ||
			([]byte(src)[i+1]>='a' && []byte(src)[i+1]<='f'))  &&
			(([]byte(src)[i+2]>='0' &&[]byte(src)[i+2]<='9') ||
				([]byte(src)[i+2]>='a' && []byte(src)[i+2]<='f')){
			i=i+2
			continue
		}else {
			return true
		}
	}
	return  false
}


func UrlEncode(src string) string {
	var des []byte
	for _,b:=range []byte(src){
		if (int(b)>=int('a') && int(b)<=int('z'))||
			(int(b)>=int('A') && int(b)<=int('Z'))||
			(int(b)>=int('0') && int(b)<=int('9'))||
			strings.Contains("-_.!~*'()",string(b)){
			des=append(des,b)
		}else if b==' '{
			des=append(des,'+')
		}else{
			des=append(des,'%')
			for _,c:=range []byte(ByteToHex([]byte{b>>4,b&0x0f})){
				if c=='0'{
					continue
				}
				des=append(des,c)
			}
		}
	}
	return  string(des)
}
func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}
var m =map[byte]int{
	'a':10,
	'b':11,
	'c':12,
	'd':13,
	'e':14,
	'f':15,
}
func UrlDecode(src string) string {
	var des []byte
	b:=[]byte(src)
	for i:=0;i<len(b);i++{
		if b[i]!='+'&& b[i]!='%'{
			des=append(des,b[i])
		}else if b[i]=='+' {
			des = append(des, ' ')
		}else{
			i++
			var s string
			if b[i+1]<='9' && b[i+1]>='0'{
				s=string((int(b[i])-int('0'))*16+int(b[i+1])-int('0'))
			}else {
				s= string((int(b[i])-int('0')) * 16 + m[b[i + 1]])
			}
			i=i+1

			des = append(des, []byte(s)[0])
		}
	}
	return  string(des)
}