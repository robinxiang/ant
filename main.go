package main

import (
	"add_number_tie/common"
	"fmt"
	"path"

	"github.com/spf13/pflag"
)

func main() {

	ascii_art := `
      _                          _              _ 
     (_)                        | |            | |
__  ___  __ _ _ __   __ _ ______| |_ ___   ___ | |
\ \/ / |/ _' | '_ \ / _' |______| __/ _ \ / _ \| |
 >  <| | (_| | | | | (_| |      | || (_) | (_) | |
/_/\_\_|\__,_|_| |_|\__, |       \__\___/ \___/|_|
                     __/ |                        
                    |___/
自动填充和拼接数字内容字符到文本行工具 202403
--------------------------------`

	info_help_cn := `
--------------------------------
参数说明：
	--between=?,? 指定拼接数字字符范围，默认为9999，例如：在每一行结尾拼接上从0001～9999的数字，则使用参数 --range=0,9999
	--exp=?,?,... 指定要排除拼接的数字，多个则用逗号分隔，默认为空，例如：排除1111，1234，则使用参数 --exp=1111,1234
	例：ant --between=0,9999 --exp=1111,2222,3333 source.txt
	`
	// 如果没有带任何参数，则打印程序帮助信息
	fmt.Println(ascii_art)

	// println(args_length, os.Args)

	// for i := 0; i < len(os.Args); i++ {
	// 	println(os.Args[i])
	// }

	// var flag_range int

	// flag.Var(&flag_range, "name", "help message for flagname")

	// 通过第三方库pflag获取命令行参数的值
	// 设置错误处理函数

	//	定义要获取的命令行参数名，默认值，帮助信息
	flag_range_between := pflag.String("between", "0,9999", "指定生成数字的范围,默认是0,9999")
	flag_exception := pflag.String("exp", "", "指定生成数字的范围,默认是99999")
	// 解析
	pflag.Parse()
	// 获取值
	value_range := *flag_range_between // 指定的生成数字范围
	value_exception := *flag_exception // 指定的排除列表

	// 获取未被指定flag的参数，其中应该包括文件路径
	args := pflag.Args()
	if len(args) < 1 {
		fmt.Println(info_help_cn)
		return
	}
	read_path := args[0]
	if len(read_path) < 5 {
		fmt.Println(info_help_cn)
	}

	//读取指定文件
	// fmt.Println(read_path)
	// fmt.Println(value_range)
	// fmt.Println(value_exception)
	// // res := string_comma_int(value_exception, ",")

	// // println(res[1])
	// 读取文本
	res, _ := common.Read_txt(read_path)
	// fmt.Println(res)
	// // // println(string_comma_int(value_exception, ","))
	// fmt.Println(common.String_comma_int(value_exception, ","))

	// tmp_list := common.String_comma_int("111,222,333", ",")
	// is_in := common.Int_in_slice(11, tmp_list)
	// fmt.Println(is_in)

	number_range := common.String_comma_int(value_range, ",")

	// tmp_int_slice := common.Make_range_list(number_range, []int{11, 13}) // 获得排除列表之后的切片
	// fmt.Println(tmp_int_slice)

	tmp_resList, _ := common.Make_result_txt(res, number_range, common.String_comma_int(value_exception, ","))
	// fmt.Println(tmp_resList)

	// fmt.Println(common.String_comma_int(value_range, ","))
	// err:=common.SaveSliceResult(tmp_resList,filepath+"")

	// 测试解析文件名
	tmp_filepath := read_path
	ext := path.Ext(tmp_filepath)                                   // 获取源文件后缀
	fileNameWithoutExt := tmp_filepath[:len(tmp_filepath)-len(ext)] // 去掉源文件后缀的路径

	saveFilePath := fileNameWithoutExt + "_result.txt"
	err := common.SaveSliceResult(tmp_resList, saveFilePath)
	if err != nil {
		fmt.Printf("写入结果文件出错：%s\n", err)
	}

}
