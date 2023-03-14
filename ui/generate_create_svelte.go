package ui

import (
	"fmt"
	"strings"
)

func handleSubmit() string {
	var sb strings.Builder
	sb.WriteString("\n\tasync function handleSubmit() {") // start
	sb.WriteString("\n\t\t// todo...please add your here to submit...")
	sb.WriteString("\n\t}\n\n") // end
	return sb.String()
}

func inputTypes(arg string) string {
	var sb strings.Builder
	// type text
	if strings.Contains(strings.ToLower(arg), "t_text") {
		sb.WriteString(" type=\"text\"")
	}
	// type email
	if strings.Contains(strings.ToLower(arg), "t_email") {
		sb.WriteString(" type=\"email\"")
	}
	// type number
	if strings.Contains(strings.ToLower(arg), "t_number") {
		sb.WriteString(" type=\"number\"")
	}
	// type password
	if strings.Contains(strings.ToLower(arg), "t_password") {
		sb.WriteString(" type=\"password\"")
	}

	// type checkbox
	if strings.Contains(strings.ToLower(arg), "t_checkbox") {
		sb.WriteString(" type=\"checkbox\"")
	}
	// type radio
	if strings.Contains(strings.ToLower(arg), "t_radio") {
		sb.WriteString(" type=\"radio\"")
	}
	return sb.String()
}

func addInputNameAndID(arg string) string {
	var sb strings.Builder
	if strings.Contains(strings.ToLower(arg), "n_") {
		for _, n := range strings.Split(arg, "::") {
			if strings.HasPrefix(n, "n_") {
				// split n_name
				sb.WriteString(fmt.Sprintf(" name=\"%v\"", strings.Split(n, "_")[1]))
				sb.WriteString(fmt.Sprintf(" id=\"%v\"", strings.Split(n, "_")[1]))
				return sb.String()
			}
		}
	}
	return sb.String()
}

func determineInput(arg string) string {
	var sb strings.Builder

	// <input />
	if strings.HasPrefix(strings.ToLower(arg), "f_input") {
		sb.WriteString("<input ")
		sb.WriteString(inputTypes(arg))
		// name
		sb.WriteString(addInputNameAndID(arg))
		sb.WriteString(" />")
		return sb.String()
	}

	// <select />
	if strings.HasPrefix(strings.ToLower(arg), "f_select") {
		sb.WriteString("<select ")
		sb.WriteString(inputTypes(arg))
		// name
		sb.WriteString(addInputNameAndID(arg))
		sb.WriteString(" />")
		return sb.String()
	}

	return sb.String()
}

func GenerateCreateSvelteComponent(args []string) string {
	// 1 - fileName
	// 2 - name of model to create
	// 3 - f_input::t_text::n_name

	var sb strings.Builder
	sb.WriteString("<script lang=\"ts\">\n") // start script
	sb.WriteString(fmt.Sprintf("%v", handleSubmit()))
	sb.WriteString("</script>") // end script

	sb.WriteString("\n\n<NoRights entityName=\"\" action=\"\">")         // open form
	sb.WriteString("\n\t<form on:submit|preventDefault={handleSubmit}>") // open form

	for _, arg := range args {
		sb.WriteString(fmt.Sprintf("\n\t%v\n", determineInput(arg)))
	}
	sb.WriteString(fmt.Sprintf("\n\t<FourthColumnGridOfFour>\n\t\t<IosButtonSmallGreen btnLabel=\"save\" btnType=\"submit\" />\n\t</FourthColumnGridOfFour>\n"))

	sb.WriteString("\n\t</form>\n\n") // close form
	sb.WriteString("</NoRights>")     // close form

	return sb.String()
}
