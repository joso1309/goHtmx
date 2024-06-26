// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package users

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"goHtmx/internal/types"
	"time"
)

func UserForm(user types.User) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-post=\"/users\" hx-push-url=\"true\" hx-swap=\"outerHTML\" class=\"space-y-4\"><input type=\"hidden\" name=\"id\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", user.ID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/users/user_form.templ`, Line: 16, Col: 67}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"grid grid-cols-1 gap-4 sm:grid-cols-2\"><div><label class=\"text-md py-2 font-medium text-gray-700\" for=\"firstName\">First name</label> <input class=\"block w-full rounded-lg border border-black p-3 text-sm\" placeholder=\"First name\" type=\"text\" id=\"firstName\" name=\"firstname\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(user.FirstName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/users/user_form.templ`, Line: 26, Col: 27}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div><label class=\"text-md py-2 font-medium text-gray-700\" for=\"lastName\">Last name</label> <input class=\"block w-full rounded-lg border border-black p-3 text-sm\" placeholder=\"Last name\" type=\"text\" id=\"lastName\" name=\"lastname\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(user.LastName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/users/user_form.templ`, Line: 37, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"grid grid-cols-1 gap-4 sm:grid-cols-2\"><div><label class=\"text-md py-2 font-medium text-gray-700\" for=\"email\">Email</label> <input class=\"w-full rounded-lg border border-black p-3 text-sm\" placeholder=\"Email address\" type=\"email\" id=\"email\" name=\"email\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(user.Email)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/users/user_form.templ`, Line: 50, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div><label class=\"text-md py-2 font-medium text-gray-700\" for=\"phone\">Phone</label> <input class=\"w-full rounded-lg border border-black p-3 text-sm\" placeholder=\"Phone Number\" type=\"tel\" id=\"phone\"></div></div><div class=\"grid grid-cols-1 gap-4 sm:grid-cols-2\"><div><label class=\"text-md py-2 font-medium text-gray-700\" for=\"date\">Date</label> <input class=\"block w-full rounded-lg border border-black p-3 text-sm\" type=\"date\" name=\"date\" id=\"date\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(time.Now().Format("2006-01-02"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/users/user_form.templ`, Line: 71, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div><div><label class=\"text-md py-2 font-medium text-gray-700\" for=\"time\">Time</label> <input class=\"block w-full rounded-lg border border-black p-3 text-sm\" type=\"time\" name=\"time\" id=\"time\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(time.Now().Format("15:04"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/users/user_form.templ`, Line: 81, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></div></div><div class=\"grid grid-cols-1 gap-4 text-center sm:grid-cols-3\"><div><label for=\"Option1\" class=\"block w-full cursor-pointer rounded-lg border border-black p-3 text-gray-600 hover:border-black has-[:checked]:border-black has-[:checked]:bg-black has-[:checked]:text-white\" tabindex=\"0\"><input class=\"sr-only\" id=\"Option1\" type=\"radio\" tabindex=\"-1\" name=\"option\"> <span class=\"text-sm\">Option 1 </span></label></div><div><label for=\"Option2\" class=\"block w-full cursor-pointer rounded-lg border border-black p-3 text-gray-600 hover:border-black has-[:checked]:border-black has-[:checked]:bg-black has-[:checked]:text-white\" tabindex=\"0\"><input class=\"sr-only\" id=\"Option2\" type=\"radio\" tabindex=\"-1\" name=\"option\"> <span class=\"text-sm\">Option 2 </span></label></div><div><label for=\"Option3\" class=\"block w-full cursor-pointer rounded-lg border border-black p-3 text-gray-600 hover:border-black has-[:checked]:border-black has-[:checked]:bg-black has-[:checked]:text-white\" tabindex=\"0\"><input class=\"sr-only\" id=\"Option3\" type=\"radio\" tabindex=\"-1\" name=\"option\"> <span class=\"text-sm\">Option 3 </span></label></div></div><div><label class=\"text-md py-2 font-medium text-gray-700\" for=\"select\">Select</label> <select class=\"block w-1/2 rounded-lg border border-black p-3 text-sm\" name=\"select\" id=\"select\"><option value=\"1\">Value 1</option> <option value=\"2\">Value 2</option> <option value=\"3\">Value 3</option></select></div><div><label class=\"sr-only\" for=\"message\">Message</label> <textarea class=\"w-full rounded-lg border border-black p-3 text-sm\" placeholder=\"Message\" rows=\"8\" id=\"message\"></textarea></div><div><button type=\"submit\" class=\"inline-block rounded-lg border border-black bg-black px-12 py-3 text-sm font-medium text-white hover:bg-transparent hover:text-black\">Spremi</button> <button hx-get=\"/users\" hx-target=\"#content\" hx-swap=\"innerHTML\" class=\"inline-block rounded-lg border border-black bg-black px-12 py-3 text-sm font-medium text-white hover:bg-transparent hover:text-black\">Cancel</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
