package main

import (
"strconv"
	"go/src/fmt"
)

var c=`unsigned char* URLDecode(unsigned char* pSrc, std::string &strDest)
{
char szTemp[2];

while (*pSrc)
{
memset(szTemp, '/0', 2 );
if( *pSrc != '%' && *pSrc != ' ')
{
strDest += *pSrc++;
continue;
}
else if( *pSrc == ' ' )
{
strDest += '+';
continue;
}
//忽略%
++pSrc;
unsigned char cTmp = 0;

for (int k = 0; k < 2 && *pSrc; ++k, ++pSrc)　　　　//例如：将%3f ---> 0x3f
{
if(*pSrc <= '9')//0x30~0x39
{
cTmp |= (*pSrc & 0x0f) << (1 - k) * 4;
}
else if(*pSrc >= 'a')//0x61~7a
{
cTmp |= (*pSrc - 0x57) << (1 - k) * 4;
}
else//0x41~5a
{
cTmp |= (*pSrc - 0x37) << (1 - k) * 4;
}
}
strDest += cTmp;
}
return 0;
}`
func main(){
	fmt.Println(int('3'))
        fmt.Println(UrlDecode("%3e%3c%7c%2b%2b"))
	fmt.Println(UrlDecode("abc1314-_.!~*'()"))
	fmt.Println(UrlDecode("abc1314-_.!~*'()++"))
	fmt.Println(UrlDecode("abc1314-_.!~*'()++%3e%3c%7c%2b%2b"))
	fmt.Println(HexToBye("030e"))
	fmt.Println([]byte{'>'>>4,'>'&0x0f})
	fmt.Println(string([]byte{'0',3<<4,'0',3<<4|14&0x57}))
	fmt.Println(string([]byte{3*16+0}))
	fmt.Println(">")
	fmt.Println('>')
	fmt.Println('9')
	fmt.Println('a')
	fmt.Println('e'-16)
	fmt.Println('e')
	var m =map[byte]int{
		'a':10,
		'b':11,
		'c':12,
		'd':13,
		'e':14,
		'f':15,
	}
	fmt.Println(m)
	fmt.Println(string(3*16+m['e']))

}

func UrlDecode(src string) string {
	var des []byte
	b:=[]byte(src)
	for i:=0;i<len(b);i++{
		fmt.Println("\n i=%d",i)
		if b[i]!='+'&& b[i]!='%'{
			des=append(des,b[i])
		}else if b[i]=='+' {
			des = append(des, ' ')
		}else{
			i++

			var m =map[byte]int{
				'a':10,
				'b':11,
				'c':12,
				'd':13,
				'e':14,
				'f':15,
			}
			//fmt.Println(b[i])
			//fmt.Println(b[i+1])
			//fmt.Println(m)
			var s string

			fmt.Println(int(b[i]))
			fmt.Println(int(b[i+1]))
			if b[i+1]<='9' && b[i+1]>='0'{
				s=string((int(b[i])-int('0'))*16+int(b[i+1])-int('0'))
				fmt.Println(int(b[i]) * 16 +int(b[i + 1]))
			}else {
				s= string((int(b[i])-int('0')) * 16 + m[b[i + 1]])
				fmt.Println(int(b[i]) * 16 + m[b[i + 1]])
			}

			fmt.Println(s)
                        i=i+1

			des = append(des, []byte(s)[0])

			//var d uint
			//var j uint
			//for j=0;j<2;j++ {
			//	if b[i]<='9'{//0x30~0x39
			//		d|=uint((b[i]&0x0f))<<(1-j)*4
			//	}else if b[i]>='a'{//0x61~7a
			//		d|=uint((b[i]-0x57))<<(1-j)*4
			//	}else {//0x41~5a
			//		d|=uint((b[i]-0x37))<<(1-j)*4
			//	}
			//i++
			//}
			//des = append(des, byte(d))

			//var c []byte
			//for j:=0;j<2;j++ {
			//	c=append(c,'0')
			//	c=append(c,b[i])
			//	i++
			//}
			//ds:=HexToBye(string(c))
			//fmt.Println(ds)
			//var e byte
			//for _,d:=range ds {
			//	e|=d<<4
			//	e=e&0x0f
			//}
			//des = append(des, byte(e))
		}
	}
	return  string(des)
}
//16进制字符串转[]byte
func HexToBye(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)

	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}