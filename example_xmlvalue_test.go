package xmlvalue_test

import (
	"fmt"

	xmlvalue "github.com/Andrew-M-C/go.xmlvalue"
)

func ExampleV_SetString() {
	x := xmlvalue.New("xml")
	x.SetString("Hello, world!").At("data", "message")
	fmt.Println(x.MustMarshalString(xmlvalue.Opt{Indent: "  "}))
	// Output:
	// <xml>
	//   <data>
	//     <message>Hello, world!</message>
	//   </data>
	// </xml>
}

func ExampleV_AddString() {
	x := xmlvalue.New("xml")
	x.AddString("Hello, world!").At("data", "message")
	x.AddString("Hello, world again!").At("data", "message")
	fmt.Println(x.MustMarshalString(xmlvalue.Opt{Indent: "  "}))
	// Output:
	// <xml>
	//   <data>
	//     <message>Hello, world!</message>
	//     <message>Hello, world again!</message>
	//   </data>
	// </xml>
}

func ExampleV_SetText() {
	x := xmlvalue.New("xml")
	x.SetText("Hello, XML!")
	fmt.Println(x.MustMarshalString())
	// Output:
	// <xml>Hello, XML!</xml>
}
