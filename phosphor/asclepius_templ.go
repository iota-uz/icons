// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Asclepius(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M136,128v72h11.62A12.25,12.25,0,0,0,160,188.53a12,12,0,0,0-5.93-10.87,4.08,4.08,0,0,1-2.06-3.49v-8.79a4,4,0,0,1,5.25-3.81A28.06,28.06,0,0,1,176,187.71c.16,15.56-13,28.29-28.59,28.29H136v16a8,8,0,0,1-16,0V216H96a8,8,0,0,1-8-8.53A8.17,8.17,0,0,1,96.27,200H120V128H104.46c-8.6,0-16,6.6-16.44,15.19a16,16,0,0,0,12.87,16.51,3.94,3.94,0,0,1,3.11,3.89V172a4,4,0,0,1-4,4,36,36,0,0,1-36-36.87C64.47,119.48,81,104,100.68,104H120V24a8,8,0,0,1,16,0v80h32a16,16,0,0,0,16-16.81C183.56,78.6,176.14,72,167.54,72H156a4,4,0,0,1-4-4V44a4,4,0,0,1,4-4h15.22c24.62,0,45.2,20.15,44.77,44.76A44,44,0,0,1,172,128ZM92.66,72H100a4,4,0,0,0,4-4V44a4,4,0,0,0-4-4H64A40,40,0,0,0,24,80v8a8,8,0,0,0,8,8H56A40,40,0,0,0,92.66,72Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M88,48v8A32,32,0,0,1,56,88H32V80A32,32,0,0,1,64,48Z\" opacity=\"0.2\"></path><line x1=\"128\" y1=\"24\" x2=\"128\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M160,166.11A24,24,0,0,1,168,184h0a24,24,0,0,1-24,24H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M160,48h17a31,31,0,0,1,31,31v1a32,32,0,0,1-32,32H96a24,24,0,0,0-24,24h0a24,24,0,0,0,24,24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M88,48v8A32,32,0,0,1,56,88H32V80A32,32,0,0,1,64,48H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"24\" x2=\"128\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M80,48v8A32,32,0,0,1,48,88H28V80A32,32,0,0,1,60,48H92\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M164,48h17a31,31,0,0,1,31,31v1a32,32,0,0,1-32,32H100a24,24,0,0,0-24,24h0a24,24,0,0,0,16,22.63\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M164,166.11A24,24,0,0,1,172,184h0a24,24,0,0,1-24,24H92\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"24\" x2=\"128\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M160,166.11A24,24,0,0,1,168,184h0a24,24,0,0,1-24,24H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M160,48h17a31,31,0,0,1,31,31v1a32,32,0,0,1-32,32H96a24,24,0,0,0-24,24h0a24,24,0,0,0,24,24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M88,48v8A32,32,0,0,1,56,88H32V80A32,32,0,0,1,64,48H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"24\" x2=\"128\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M160,166.11A24,24,0,0,1,168,184h0a24,24,0,0,1-24,24H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M160,48h17a31,31,0,0,1,31,31v1a32,32,0,0,1-32,32H96a24,24,0,0,0-24,24h0a24,24,0,0,0,24,24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M88,48v8A32,32,0,0,1,56,88H32V80A32,32,0,0,1,64,48H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"24\" x2=\"128\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M160,166.11A24,24,0,0,1,168,184h0a24,24,0,0,1-24,24H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M160,48h17a31,31,0,0,1,31,31v1a32,32,0,0,1-32,32H96a24,24,0,0,0-24,24h0a24,24,0,0,0,24,24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M88,48v8A32,32,0,0,1,56,88H32V80A32,32,0,0,1,64,48H96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
