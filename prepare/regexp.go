package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	const text = `My email is ahanru723@163.com@haha.com My email is hanru723@163.com@haha.com
email is wangergou@sina.com
email is kongyixueyuan@cldy.org.cn`
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)
	allMatch := re.FindAllString(text, -1)
	fmt.Println(allMatch)
	//fmt.Println(re.MatchString(text))


	// 身份证：  练习3：身份证号：18位。0不能开头第一位：非0,16位。最后一位：(数字|X)
	//fmt.Println(regexp.MustCompile(`[1-9]\d{17}|[1-9]\d{16}X`).MatchString(`13141219880712321X`))
	//fmt.Println(regexp.MustCompile(`^[1-9]\d{16}(\d|X)$`).MatchString(`13141219880712321X`))


	s1 := `wangergou@163.com`
	s2 := `sanpang@qq.com` // 828384848@qq.com
	s3 := `lixiaohua@sina.com`

	fmt.Println(regexp.MustCompile(`\w+@(163|qq)\.com`).MatchString(s1))
	fmt.Println(regexp.MustCompile(`\w+@(163|qq)\.com`).MatchString(s2))
	fmt.Println(regexp.MustCompile(`\w+@(163|qq)\.com`).MatchString(s3))
	fmt.Println(regexp.MustCompile(`\w+@(163|sina|yahoo|qq)\.(com|cn)`).MatchString(`sanpang@163.com`))

	fmt.Println("----------------------")

	s4 := `<html><h1>helloworld</h1></html>李小花李小花`
	re1 := regexp.MustCompile(`<(.+)><.+>(.+)</.+></.+>`)

	fmt.Println(re1.NumSubexp()) //2
	names := re1.SubexpNames()
	fmt.Println(names)

	res1 := re1.FindAllStringSubmatch(s4, -1)
	fmt.Println(res1)
	fmt.Println(res1[0])
	fmt.Println(res1[0][0])
	fmt.Println(res1[0][1])
	fmt.Println(res1[0][2])

	fmt.Println("----------------------")

	s5 := `<html><body><h1>hello</h1></body></html>`
	re2 := regexp.MustCompile(`<(?P<t1>.+)><(?P<t2>.+)><(?P<t3>.+)>(?P<t4>.+)</(.+)></(.+)></(.+)>`)
	fmt.Println(re2.NumSubexp())
	fmt.Println(re2.SubexpNames())

	//获取分组名称
	for i := 0; i <= re2.NumSubexp(); i++ {
		fmt.Printf("%d: \t %q\n ", i, re2.SubexpNames()[i])
	}
	res2 := re2.FindAllStringSubmatch(s5, -1)
	fmt.Println(res2)

	s := "This is a number 234-245-236"
	//获取数字部分
	b5 := regexp.MustCompile(`(.+?)(\d+-\d+-\d+)`).FindAllStringSubmatch(s,-1)
	fmt.Println(b5)
	//fmt.Println(b5[0][1])
	//fmt.Println(b5[0][2])

	hhdir, _ := ioutil.TempDir("", "hhtmp")
	f, _ := ioutil.TempFile(hhdir, "hihi.txt")
	fmt.Fprint(f,"sss")
	f.Close()

}
