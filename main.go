package main

import (
	"os"
)

func main() {
	if len(os.Args) >= 2 && len(os.Args) <= 4 {
		var arg3 string
		var arg4 string
		arg1 := os.Args[1]
		arg2 := DefineArg()
		if len(os.Args) == 4 || len(os.Args) == 3 || len(os.Args) == 2 {
			arg3 = os.Args[1]
			if IsFile(arg3) {
				arg1 = os.Args[2]
				arg2 = DefineArg()
			} else if IsColor(arg3) {
				if len(os.Args) > 4 {
					errorColor()
				}
				if len(os.Args) == 3 {
					arg1 = os.Args[2]
					arg2 = DefineArg()
				} else {
					arg1 = os.Args[3]
					arg2 = DefineArg()
					arg4 = os.Args[2]
				}
			} else {
				arg3 = ""
			}
		}

		if len(os.Args) == 4 && !IsFile(arg3) && !IsColor(arg3) {
			error()
		}
		if CaseResolved(arg1, arg3) {
			return
		}
		texte := foundFile(arg2)

		asciiArt(arg1, arg3, arg4, texte)
	} else if len(os.Args) > 1 {
		if IsColor(os.Args[1]) && len(os.Args) > 4 {
			errorColor()
		} else if IsFile(os.Args[1]) && len(os.Args) > 4 {
			errorOutput()
		}
	}
}
