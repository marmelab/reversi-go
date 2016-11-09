package matrix

import (
	"bufio"
	"bytes"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"reflect"
	"strings"
)

func Render(matrix [][]string) string {

	var buffer bytes.Buffer
	xSize, _ := GetSize(matrix)
	underlined := color.New(color.Underline).SprintFunc()

	buffer.WriteString(strings.Repeat("_", int((xSize*2)+1)) + "\n")

	for _, row := range matrix {
		buffer.WriteString("|")
		for _, cellContent := range row {
			buffer.WriteString(underlined(cellContent) + "|")
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func RenderAscii(matrix [][]string) string {

	var buffer bytes.Buffer
	stringBuffer := bufio.NewWriter(&buffer)
	table := tablewriter.NewWriter(stringBuffer)

	for _, v := range matrix {
		table.Append(v)
	}

	table.Render()
	stringBuffer.Flush()

	return buffer.String()

}

func GetSize(matrix interface{}) (int, int) {
	if reflect.TypeOf(matrix).Kind() == reflect.Slice {
		matrixSlice := reflect.ValueOf(matrix)
		if matrixSlice.Len() == 0 {
			return 0, 0
		}
		return matrixSlice.Index(0).Len(), matrixSlice.Len()
	}
	return 0, 0
}
