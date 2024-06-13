// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SideBar() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"hidden md:flex flex-col w-64 bg-gray-800\"><div class=\"flex items-center justify-center h-16 bg-gray-900\"><span class=\"text-white font-bold uppercase\">HROTE</span></div><div class=\"flex flex-col flex-1 overflow-y-auto\"><nav class=\"flex-1 px-2 py-4 bg-gray-800\"><a href=\"/\" class=\"flex items-center rounded-lg px-4 py-2 text-gray-100 hover:bg-gray-500\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6 mr-2\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h16M4 18h16\"></path></svg> Home</a> <a href=\"/users\" class=\"flex items-center rounded-lg px-4 py-2 mt-2 text-gray-100 hover:bg-gray-500\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6 mr-2\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M6 18L18 6M6 6l12 12\"></path></svg> Users</a> <a href=\"/login\" class=\"flex items-center rounded-lg px-4 py-2 mt-2 text-gray-100 hover:bg-gray-500\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6 mr-2\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 10V3L4 14h7v7l9-11h-7z\"></path></svg> Login</a></nav></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
