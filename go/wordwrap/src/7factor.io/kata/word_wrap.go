package kata

import (
	"fmt"
	"strings"
)

func Wrap(toWrap string, column int) (string) {
	if len(toWrap) <= column {
		return toWrap
	} else {
		return wrapOnBoundary(toWrap, column)
	}
}

func wrapOnBoundary(toWrap string, column int) (string){
	space := strings.LastIndex(toWrap[0:column]," ")
	fmt.Printf("%v", space)
	if space != -1 {
		fmt.Printf("%s", toWrap[0:space])
		return toWrap[0:space] + "\n" + Wrap(toWrap[(space+1):], column)
	} else {
		return toWrap[0:column] + "\n" + Wrap(toWrap[column:], column)
	}
}