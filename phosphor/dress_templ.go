// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Dress(props Props) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M66.26,80.23a15.26,15.26,0,0,1-1.65-12.17,15.54,15.54,0,0,1,2-4.76L88,32.7V8a8,8,0,0,1,8.53-8A8.17,8.17,0,0,1,104,8.27V32.42L109.25,39a23.91,23.91,0,0,0,19.13,9,24.67,24.67,0,0,0,18.71-9.43L152,32.42V8a8,8,0,0,1,8.53-8A8.17,8.17,0,0,1,168,8.27V32.7l21.42,30.6a15.54,15.54,0,0,1,2,4.76,15.26,15.26,0,0,1-1.65,12.17,1.74,1.74,0,0,0-.11.18l-13.85,21.74A4,4,0,0,1,172.4,104H83.6a4,4,0,0,1-3.38-1.85L66.37,80.41A1.74,1.74,0,0,0,66.26,80.23Zm148.5,129.56a2.52,2.52,0,0,0-.15-.34L173.69,122.3a4,4,0,0,0-3.63-2.3H85.94a4,4,0,0,0-3.63,2.3L41.39,209.45a2.52,2.52,0,0,0-.15.34A16.19,16.19,0,0,0,41.6,223,16,16,0,0,0,56,232H200a16,16,0,0,0,14.39-9A16.19,16.19,0,0,0,214.76,209.79Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M200,224a8,8,0,0,0,7.35-11.15L160,112l22.86-35.88a8,8,0,0,0,0-8.24L160,35.22,153,44a32,32,0,0,1-50,0l-7-8.77L73.14,67.88a8,8,0,0,0,0,8.24L96,112,48.66,212.85A8,8,0,0,0,56,224Z\" opacity=\"0.2\"></path><line x1=\"160\" y1=\"35.22\" x2=\"160\" y2=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"8\" x2=\"96\" y2=\"35.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M200,224a8,8,0,0,0,7.35-11.15L160,112l22.86-35.88a8,8,0,0,0,0-8.24L160,35.22,153,44a32,32,0,0,1-50,0l-7-8.77L73.14,67.88a8,8,0,0,0,0,8.24L96,112,48.66,212.85A8,8,0,0,0,56,224Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"96\" y1=\"112\" x2=\"160\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"160\" y1=\"35.22\" x2=\"160\" y2=\"12\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"96\" y1=\"12\" x2=\"96\" y2=\"35.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M200,224a8,8,0,0,0,7.35-11.15L160,112l22.86-35.88a8,8,0,0,0,0-8.24L160,35.22,153,44a32,32,0,0,1-50,0l-7-8.77L73.14,67.88a8,8,0,0,0,0,8.24L96,112,48.66,212.85A8,8,0,0,0,56,224Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"96\" y1=\"112\" x2=\"160\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"160\" y1=\"35.22\" x2=\"160\" y2=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"96\" y1=\"8\" x2=\"96\" y2=\"35.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M200,224a8,8,0,0,0,7.35-11.15L160,112l22.86-35.88a8,8,0,0,0,0-8.24L160,35.22,153,44a32,32,0,0,1-50,0l-7-8.77L73.14,67.88a8,8,0,0,0,0,8.24L96,112,48.66,212.85A8,8,0,0,0,56,224Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"96\" y1=\"112\" x2=\"160\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"160\" y1=\"35.22\" x2=\"160\" y2=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"96\" y1=\"8\" x2=\"96\" y2=\"35.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M200,224a8,8,0,0,0,7.35-11.15L160,112l22.86-35.88a8,8,0,0,0,0-8.24L160,35.22,153,44a32,32,0,0,1-50,0l-7-8.77L73.14,67.88a8,8,0,0,0,0,8.24L96,112,48.66,212.85A8,8,0,0,0,56,224Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"96\" y1=\"112\" x2=\"160\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"160\" y1=\"35.22\" x2=\"160\" y2=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"8\" x2=\"96\" y2=\"35.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M200,224a8,8,0,0,0,7.35-11.15L160,112l22.86-35.88a8,8,0,0,0,0-8.24L160,35.22,153,44a32,32,0,0,1-50,0l-7-8.77L73.14,67.88a8,8,0,0,0,0,8.24L96,112,48.66,212.85A8,8,0,0,0,56,224Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"96\" y1=\"112\" x2=\"160\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			return nil
		})
		templ_7745c5c3_Err = svg(props).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
