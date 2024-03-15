package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 读取指定文件的函数，返回字符串列表和错误
func Read_txt(read_path string) ([]string, error) {
	// 打开文件
	file, err := os.Open(read_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//读取到的文件
	var lines []string

	// 逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("读取到%d行数据\n", len(lines))

	return lines, nil
}

// 生成结果文件的例子,old_txt:原始字符串列表；number_range：要生成的数字范围；exp_list:要排除指定数字列表
func Make_result_txt(old_txt []string, number_range []int, exp_list []int) ([]string, error) {
	// println("正在生成结果数据……")
	var result_list []string //原始字符串列表
	var range_list []string  // 生成的指定范围生成字符串列表
	for _, str := range old_txt {
		tmp_slice := Make_range_list(number_range, exp_list)
		for _, str2 := range tmp_slice {
			range_list = append(range_list, str+str2) // 将原读取的一行字符与生成的字符串拼接，写入列表
		}
		result_list = append(result_list, range_list...) // 将生成的拼接结果slice，与输出结果拼接
		range_list = range_list[:0]
	}
	return result_list, nil
}

// 生成指定范围数字组成的字符串，不足部分用0补齐,exception_list是排除数字列表
func Make_range_list(number_range []int, exception_list []int) []string {
	var result_list []string

	for i := number_range[0]; i < number_range[1]; i++ {
		// 如果当前数字在排除列表中，则跳过
		if Int_in_slice(i, exception_list) {
			continue
		}
		result_list = append(result_list, fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(number_range[1])))+"d\n", i))
	}

	return result_list

}

// 判断指定数字是否在切片中,返回布尔值
func Int_in_slice(number_find int, target_slice []int) bool {
	// 遍历切片循环查找
	for _, value := range target_slice {
		if value == number_find {
			return true
		}
	}
	return false
}

// 将逗号分割的字符串，转换为[]int,参数指定要转换的字符串和分割字符串
func String_comma_int(input_string string, str_split string) []int {
	if len(input_string) < 1 {
		return nil
	}
	// 使用 strings.Split() 函数将字符串分割为子串
	substrings := strings.Split(input_string, str_split)

	// 遍历子串并将其转换为整数
	var numbers []int
	for _, substr := range substrings {
		num, err := strconv.Atoi(substr)
		if err != nil {
			fmt.Printf("数据有误，无法作为数字使用: %v\n", err)
			return nil
		}
		numbers = append(numbers, num)
	}

	return numbers

}

// 将[]string保存到结果文件
func SaveSliceResult(result_slice []string, file_path string) error {
	fmt.Println("正在写入结果数据...")
	file, err := os.Create(file_path)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, str := range result_slice {
		_, err := fmt.Fprint(writer, str)
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return err
		}
	}
	fmt.Printf("写入完成：%s \n", file_path)
	return nil
}
