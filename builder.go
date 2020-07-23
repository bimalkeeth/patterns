package main

import (
	"fmt"
	"strings"
)

const indentSize =2

type HtmlElements struct {
	name,text string
	elements []HtmlElements
}

func (e *HtmlElements) String() string {
	return e.string(0)
}

func (e *HtmlElements) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent+1))
	}
	return sb.String()
}


type HtmlBuilder struct {
	rootName string
	root HtmlElements
}

func NewHtmlBuilder(rootName string)*HtmlBuilder{
	return &HtmlBuilder{rootName,HtmlElements{rootName,"",[]HtmlElements{}}}
}

func (b *HtmlBuilder) String() string{
	return b.root.String()
}

func(b * HtmlBuilder)AddChild(childName,childText string){
	e:=HtmlElements{childName,childText,[]HtmlElements{}}
	b.root.elements=append(b.root.elements,e)
}

func main() {
	b:=NewHtmlBuilder("ul")
	b.AddChild("li","Hello")
	b.AddChild("li","World")
	b.AddChild("ul","")
	fmt.Println(b.String())
}
