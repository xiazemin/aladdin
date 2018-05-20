package main
import (
	"fmt"
	"strings"
	"encoding/hex"
	"bytes"
	"strconv"
)
var c=`unsigned char CHAR_TO_HEX( unsigned char ch )
{
//0xa(10)转化为字符'A'(65)，要加上55 //0x0(0)转化为字符'0'(48)，要加上48
return (unsigned char)(ch > 9 ? ch + 55: ch + 48);
}

std::string URLEncode(const char* pSrc)
{
unsigned char chTemp;
std::string strDest;

while (*pSrc)
{
chTemp = (unsigned char)*pSrc;
if ( (chTemp >= 'a' && chTemp <= 'z') || (chTemp >= 'A' && chTemp <= 'Z') || (chTemp >= '0' && chTemp <= '9')
|| strchr("-_.!~*'()", chTemp))
{
strDest += chTemp;
}
else if (chTemp == ' ')
{
*pDest++ = '+';
}
else
{
strDest += '%';
strDest += CHAR_TO_HEX( (unsigned char)(chTemp >> 4) );
strDest += CHAR_TO_HEX( (unsigned char)(chTemp & 0x0f) );
}
++pSrc;
}

return strDest;
}`

//十进制切片数组转换为16进制字符串

func CharToHex(DecimalSlice []byte) string {
	var sa = make([]string, 0)
	for _, v := range DecimalSlice {
		sa = append(sa, fmt.Sprintf("%02X", v))
	}
	ss := strings.Join(sa, "")
	return ss
}
func main(){
fmt.Println(UrlEncode("abc1314-_.!~*'()"))
	fmt.Println(UrlEncode("abc1314-_.!~*'()  "))
	fmt.Println(UrlEncode("abc1314-_.!~*'()  ><|++"))
	fmt.Println(string('|'>>4)+string('|'&0x0f))
	var a [] byte
	a=append(a,'|'>>4)
	a=append(a,'|'&0x0f)
	fmt.Println(hex.EncodeToString(a))
	fmt.Println(hex.EncodeToString([]byte{'|'>>4}))
}

func UrlEncode(src string) string {
 var des []byte
	for _,b:=range []byte(src){
		if (int(b)>=int('a') && int(b)<=int('z'))||
		(int(b)>=int('A') && int(b)<=int('Z'))||
		(int(b)>=int('0') && int(b)<=int('9'))||
		strings.Contains("-_.!~*'()",string(b)){
			des=append(des,b)
			//if ( (chTemp >= 'a' && chTemp <= 'z') || (chTemp >= 'A' && chTemp <= 'Z')
			// || (chTemp >= '0' && chTemp <= '9')
			//|| strchr("-_.!~*'()", chTemp))
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
			//var a =[]byte{b>>4,b&0x0f}
			//s:=hex.EncodeToString(a)
			//for _,c:=range []byte(s){
			//	des=append(des,c)
			//}
			// a=[]byte{b&0x0f}
			// s =hex.EncodeToString(a)
			//for _,c:=range []byte(s){
			//	des=append(des,c)
			//}
			//strDest += CHAR_TO_HEX( (unsigned char)(chTemp >> 4) );
		//strDest += CHAR_TO_HEX( (unsigned char)(chTemp & 0x0f) );
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
