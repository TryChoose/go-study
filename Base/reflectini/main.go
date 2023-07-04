package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//ini 配置文件解析器

//MysqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//RedisConfig 配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(f string, data interface{}) (err error) {
	//0.参数的校验
	//0.1传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t1 := reflect.TypeOf(data)
	fmt.Println(t1, t1.Kind())
	if t1.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") //新创建一个错误
		return
	}
	//0.2判断必须是一个结构体指针(因为配置文件中各种键值对对需要赋值给结构体的字段)
	if t1.Elem().Kind() != reflect.Struct { //反射中使用专有的Elem()方法来获取指针对应的值
		err = errors.New("data param should be a struct pointer") //新创建一个错误
		return
	}
	//1.读取文件得到字节类型数据
	b, err := ioutil.ReadFile(f) //file 字符串变量
	if err != nil {
		return
	}
	s := string(b) //将文件内容转为字符串
	lineSlice := strings.Split(s, "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	//2.一行一行得到数据
	var structName string
	for idx, line := range lineSlice {
		//过滤空格
		line = strings.TrimSpace(line)
		//如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		//2.1如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//2.2如果是[开头的就表示是节
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line :%d syntax error", idx+1)
				return
			}
			//把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line :%d syntax error", idx+1)
			}
			//根据字符串sectionName去data里面根据反射找到结构体
			//Vof := reflect.ValueOf(data)
			for i := 0; i < t1.Elem().NumField(); i++ {
				field := t1.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}

		} else {
			//2.3如果不是[开头的就是=分割的键值对
			//1.以等号分割这一行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//2.根据structName 去 data把对应嵌套结构体的值取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是有个结构体", structName)

			}
			var fieldName string
			var fileType reflect.StructField
			//3.遍历嵌套结构体的每一个字段，判断没一个tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) //tag信息是存储在类型信息当中的
				fileType = filed
				if filed.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = filed.Name
					break
				}
			}
			//4.如果key=tag ,给这个字段赋值
			//4.1根据fieldName 去取出这个字段
			if len(fieldName) == 0 {
				//在结构体中找不到对应的字段
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			//4.2对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line :%d value type error ", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line :%d value type error ", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line :%d value type error ", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}

	}

	return
}
func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v", err)
		return
	}
	fmt.Println(cfg)
}
