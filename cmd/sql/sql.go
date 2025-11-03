// Package sql @Author:冯铁城 [17615007230@163.com] 2025-10-31 20:12:16
package sql

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"go-ftc-console/common"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// flag变量
var (
	csvFile    string
	outputFile string
	db         string
	table      string
	listTable  bool
)

// NewSqlCommand 创建sql命令
func NewSqlCommand() *cobra.Command {

	//1.设置Flags
	sqlCmd.Flags().StringVarP(&csvFile, "csvFile", "c", "", "CSV文件路径")
	sqlCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "输出文件路径")
	sqlCmd.Flags().StringVarP(&db, "db", "d", "", "数据库")
	sqlCmd.Flags().StringVarP(&table, "table", "t", "", "表")
	sqlCmd.Flags().BoolVarP(&listTable, "list table", "l", false, "输出内置表信息")

	//2.返回
	return sqlCmd
}

// sql命令 将csv数据转换成sql
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "transform data from csv to sql",
	Run: func(cmd *cobra.Command, args []string) {

		//1.如果打印表信息，则打印并返回，否则执行SQL命令
		if listTable {
			consoleTables()
			return
		} else {
			runCommand()
		}
	},
}

// 打印表信息
func consoleTables() {

	//1.打印分割线
	fmt.Println("--------------------------------------------------------------------------------")

	//2.打印表头
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "| 表名                  \t| 列名 \t|")
	fmt.Fprintln(w, "--------------------------------------------------------------------------------")

	//3.打印项目信息
	for tableName, columns := range tableColumnMap {
		fmt.Fprintf(w, "| %-18s\t|%-18s\t|\n", tableName, strings.Join(columns, ", "))
		fmt.Fprintln(w, "--------------------------------------------------------------------------------")
	}

	//4.写入控制台
	w.Flush()
}

// 运行命令
func runCommand() {

	//1.获取读取器
	err, reader, sourceFile := getReader()
	defer sourceFile.Close()
	if err != nil {
		fmt.Printf("获取 CSV 读取器失败:%s", err)
		return
	}

	//2.获取写入器
	err, sqlFile, writer := getWriter()
	defer sqlFile.Close()
	defer writer.Flush()
	if err != nil {
		fmt.Printf("获取 SQL 写入器失败:%s", err)
		return
	}

	//3.读取表头（跳过）
	_, err = reader.Read()
	if err != nil {
		fmt.Printf("读取 CSV 表头失败:%s", err)
		return
	}

	//4.如果数据库或表为空，则使用默认库表
	if db == "" {
		db = defaultDB
	}
	if table == "" {
		table = defaultTable
	}

	//5.如果根据表无法获取表字段，则返回
	columns, exist := tableColumnMap[table]
	if !exist {
		fmt.Printf("无法获取表字段，请检查表名:%s", table)
		return
	}

	//5.循环读取每行数据,写入到输出文件中
	line := writeDataToOutputFile(reader, writer, columns, table)

	//6.打印完成信息
	fmt.Printf("SQL 生成完成，共处理 %d 行数据\n", line-2)
}

// 获取CSV读取器
func getReader() (error, *csv.Reader, *os.File) {

	//1.如果CSV文件为空，则返回错误
	if csvFile == "" {
		return errors.New("请指定CSV文件"), nil, nil
	}

	//2.打开 CSV 文件
	fmt.Printf("CSV文件路径: %s\n", csvFile)
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("打开 CSV 文件失败:%s", err)
		return err, nil, nil
	}

	//3.创建 CSV 读取器
	reader := csv.NewReader(file)

	//4.设置读取参数
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	//5.返回
	return nil, reader, file
}

// 获取SQL写入器
func getWriter() (error, *os.File, *bufio.Writer) {

	//1.如果输出文件路径为空，则使用默认输出文件路径
	if outputFile == "" {
		outputFile = defaultOutput
	}

	//2.打开输出文件
	fmt.Printf("输出文件路径: %s\n", outputFile)
	sqlFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("创建输出文件失败:%s", err)
		return err, nil, nil
	}

	//3.创建writer
	writer := bufio.NewWriter(sqlFile)

	//4.返回
	return nil, sqlFile, writer
}

// 写入数据到输出文件中
func writeDataToOutputFile(reader *csv.Reader, writer *bufio.Writer, columns []string, table string) int {

	//1.定义初始读取行数为第2行
	line := 2

	//2.循环读取excel
	for {

		//3.读取数据
		row, err := reader.Read()
		if err != nil {

			//4.如果读取完毕，退出循环
			if err.Error() == "EOF" {
				break
			}

			//5.如果是其他异常，则打印并继续
			fmt.Printf("读取第 %d 行失败: %v", line, err)
			line++
			continue
		}

		//6.如果数据列数不足，则打印并跳过
		if len(row) < len(columns) {
			fmt.Printf("警告: 第 %d 行数据列数不足（%d < %d），跳过\n", line, len(row), len(columns))
			line++
			continue
		}

		//7.处理数据列
		values := parseRow(columns, row)

		//8.解析SQL
		sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);\n",
			table,
			strings.Join(columns, ", "),
			strings.Join(values, ", "),
		)

		//9.写入文件
		_, err = writer.WriteString(sql)
		if err != nil {
			fmt.Printf("写入 SQL 失败: %v", err)
		}

		//10.line++
		line++
	}

	//11.返回行数
	return line
}

// 处理数据列
func parseRow(columns []string, row []string) []string {

	//1.定义列值切片
	values := make([]string, len(columns))

	//2.处理数据列
	for j, cellValue := range row[:len(columns)] {

		//3.读取列值
		cellValue = strings.TrimSpace(cellValue)

		//4.空值处理
		if cellValue == "" {
			values[j] = "NULL"
			continue
		}

		//5.根据不同类型，进行转换并存入切片
		//数字类型:直接存储
		//字符串:转义单引号 + 加引号
		if common.IsNumeric(cellValue) {
			values[j] = cellValue
		} else {
			cellValue = strings.ReplaceAll(cellValue, "'", "''")
			values[j] = "'" + cellValue + "'"
		}
	}

	//6.返回列值切片
	return values
}
