package evaluator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ayobami0/yoruba/src/object"
)

var builtins = map[string]*object.Builtin{
	"ka": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) > 1 {
				return newError(WRONG_ARGUMENT_NO, 1, 5)
			}

			if len(args) == 1 {
				switch arg := args[0].(type) {
				case *object.String:
				case *object.Boolean:
				case *object.Number:
				default:
					return newError(INVALID_TYPE, arg.Type())
				}
				fmt.Print(args[0].Inspect())
			}

			var b string
			reader := bufio.NewReader(os.Stdin)

			b, err := reader.ReadString('\n')

			if err != nil {
				return newError(UNKNOWN_ERROR)
			}

			return &object.String{Value: strings.Trim(b, "\n")}
		},
	},
	"ko": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var strbuf strings.Builder

			if len(args) == 0 {
				return newError(WRONG_ARGUMENT_NO, "ju ọkan lọ", 0)
			}

			for i, v := range args {
				switch v.(type) {
				case *object.String:
				case *object.Boolean:
				case *object.Number:
				default:
					return newError(UNKNOWN_ERROR)
				}
				strbuf.WriteString(v.Inspect())
				if i != len(args)-1 {
					strbuf.WriteByte(' ')
				}
			}

			fmt.Println(strbuf.String())
			return nil
		},
	},
}
