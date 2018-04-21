package main

import (
	"bufio"
	"fmt"
	"github.com/zhouziqunzzq/InfoSecHomework/homework1"
	"github.com/zhouziqunzzq/InfoSecHomework/homework2"
	"github.com/zhouziqunzzq/InfoSecHomework/homework3"
	"os"
)

const (
	ClearScreenCMD = "\033[2J\033[1;1H"
)

func Pause() {
	fmt.Print("按下回车以继续...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Clear() {
	fmt.Print(ClearScreenCMD)
	return
}

func ScanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	return string(b)
}

func PrintHeader() {
	fmt.Println("==============================================")
	fmt.Println("              信息安全基础作业演示")
	fmt.Println()
	fmt.Println("  姓名：周子群  学号：20154491  班级：计算机1502")
	fmt.Println("==============================================")
}

func PrintMenu() {
	PrintHeader()
	fmt.Printf("请输入欲演示的项目编号：\n")
	fmt.Printf("1. 作业1（双轨、钥控）\n")
	fmt.Printf("2. 作业2（逆元、换位）\n")
	fmt.Printf("3. 作业3（Vigenere）\n")
	fmt.Printf("4. 退出\n")
	fmt.Printf("> ")
}

func ShowHomework1() {
	choice := 0
	for choice != 3 {
		Clear()
		fmt.Println("==============================================")
		fmt.Println("                    作业1")
		fmt.Println("==============================================")
		fmt.Printf("请输入欲演示的项目编号：\n")
		fmt.Printf("1. 双轨密码\n")
		fmt.Printf("2. 钥控密码\n")
		fmt.Printf("3. 返回\n")
		fmt.Printf("> ")
		fmt.Scanf("%v", &choice)
		switch choice {
		case 1:
			Clear()
			fmt.Println("==============================================")
			fmt.Println("               作业1 - 双轨密码")
			fmt.Println("==============================================")
			fmt.Printf("请输入明文：\n")
			fmt.Printf("> ")
			m := ScanLine()
			fmt.Printf("加密结果：%v\n", homework1.DoubleRailEncryptor(m))
			Pause()
		case 2:
			Clear()
			fmt.Println("==============================================")
			fmt.Println("               作业1 - 钥控密码")
			fmt.Println("==============================================")
			fmt.Printf("请输入明文：\n")
			fmt.Printf("> ")
			m := ScanLine()
			fmt.Printf("请输入密钥：\n")
			fmt.Printf("> ")
			var k string
			fmt.Scanf("%v", &k)
			fmt.Printf("加密结果：%v\n", homework1.KeyControlledEncryptor(m, k))
			Pause()
		case 3:
			break
		default:
			continue
		}
	}
}

func ShowHomework2() {
	choice := 0
	for choice != 3 {
		Clear()
		fmt.Println("==============================================")
		fmt.Println("                    作业2")
		fmt.Println("==============================================")
		fmt.Printf("请输入欲演示的项目编号：\n")
		fmt.Printf("1. 乘法逆元\n")
		fmt.Printf("2. 换位密码\n")
		fmt.Printf("3. 返回\n")
		fmt.Printf("> ")
		fmt.Scanf("%v", &choice)
		switch choice {
		case 1:
			Clear()
			fmt.Println("==============================================")
			fmt.Println("               作业2 - 乘法逆元")
			fmt.Println("==============================================")
			fmt.Printf("请输入欲求逆元的数：\n")
			fmt.Printf("> ")
			var m int
			fmt.Scanf("%v", &m)
			fmt.Printf("请输入模数：\n")
			fmt.Printf("> ")
			var n int
			fmt.Scanf("%v", &n)
			if rst, err := homework2.MultiplicativeInverse(m, n); err != nil {
				fmt.Printf("错误：%v\n", err)
			} else {
				fmt.Printf("结果：%v\n", rst)
			}
			Pause()
		case 2:
			Clear()
			fmt.Println("==============================================")
			fmt.Println("               作业2 - 换位密码")
			fmt.Println("==============================================")
			fmt.Printf("请输入明文：\n")
			fmt.Printf("> ")
			m := ScanLine()
			fmt.Printf("请输入换位密钥长度：\n")
			fmt.Printf("> ")
			var n = -1
			fmt.Scanf("%v", &n)
			for n <= 0 {
				fmt.Printf("输入错误！请输入一个大于0的整数：\n")
				fmt.Printf("> ")
				fmt.Scanf("%v", &n)
			}
			for len(m)%n != 0 {
				fmt.Printf("输入错误！明文长度不是换位密钥长度的倍数，请重试：\n")
				fmt.Printf("> ")
				fmt.Scanf("%v", &n)
			}
			key := make([]int, n)
			fmt.Printf("请输入%v个不重复的换位密钥，每输入完一个密钥按下回车：\n", n)
			for i := 0; i < n; i++ {
				fmt.Printf("> ")
				fmt.Scanf("%v", &key[i])
			}
			if rst, err := homework2.PositionSwitchEncryptor(m, key); err != nil {
				fmt.Printf("错误：%v\n", err)
			} else {
				fmt.Printf("加密结果：%v\n", rst)
			}
			Pause()
		case 3:
			break
		default:
			continue
		}
	}
}

func ShowHomework3() {
	choice := 0
	for choice != 3 {
		Clear()
		fmt.Println("==============================================")
		fmt.Println("                    作业3")
		fmt.Println("==============================================")
		fmt.Printf("请输入欲演示的项目编号：\n")
		fmt.Printf("1. 使用作者的姓名全拼缩写（ZZQ）进行Vigenere加密\n")
		fmt.Printf("2. 使用自定义的密钥进行Vigenere加密\n")
		fmt.Printf("3. 返回\n")
		fmt.Printf("> ")
		fmt.Scanf("%v", &choice)
		switch choice {
		case 1:
			Clear()
			fmt.Println("==============================================")
			fmt.Println("         作业3 - 使用ZZQ进行Vigenere加密")
			fmt.Println("==============================================")
			fmt.Printf("请输入明文(仅包含字母)：\n")
			fmt.Printf("> ")
			m := ScanLine()
			fmt.Printf("加密结果：%v\n", homework3.VigenereEncryptor(m, "ZZQ"))
			Pause()
		case 2:
			Clear()
			fmt.Println("==============================================")
			fmt.Println("     作业3 - 使用自定义的密钥进行Vigenere加密")
			fmt.Println("==============================================")
			fmt.Printf("请输入明文(仅包含字母)：\n")
			fmt.Printf("> ")
			m := ScanLine()
			fmt.Printf("请输入密钥：\n")
			fmt.Printf("> ")
			var k string
			fmt.Scanf("%v", &k)
			fmt.Printf("加密结果：%v\n", homework3.VigenereEncryptor(m, k))
			Pause()
		case 3:
			break
		default:
			continue
		}
	}
}

func main() {
	choice := 0
	for choice != 4 {
		Clear()
		PrintMenu()
		fmt.Scanf("%v", &choice)
		switch choice {
		case 1:
			ShowHomework1()
		case 2:
			ShowHomework2()
		case 3:
			ShowHomework3()
		case 4:
			fmt.Println("感谢使用，再见")
			break
		default:
			continue
		}
	}
}
