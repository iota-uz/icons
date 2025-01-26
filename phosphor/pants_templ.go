// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Pants(props Props) templ.Component {
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
			if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M193.94,39A8,8,0,0,0,186,32H70a8,8,0,0,0-7.94,7L54,104c.68,0,1.36,0,2,0A40,40,0,0,0,96,64h64a40,40,0,0,0,40,40c.69,0,1.37,0,2.05,0Z\" opacity=\"0.2\" fill=\"currentColor\"></path><path d=\"M186,32a8,8,0,0,1,7.94,7l22,176a8,8,0,0,1-7.94,9H167.31a8,8,0,0,1-7.76-6.06L128,96,96.45,217.94A8,8,0,0,1,88.69,224H48a8,8,0,0,1-7.94-9l22-176A8,8,0,0,1,70,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"128\" y1=\"96\" x2=\"128\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"58.94\" y1=\"64\" x2=\"197.06\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M96,64a40,40,0,0,1-40,40c-.69,0-1.37,0-2,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M202.06,104c-.68,0-1.37,0-2.06,0a40,40,0,0,1-40-40\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M92,68H57.94a4,4,0,0,0-4,3.5l-6,48a4,4,0,0,0,4,4.5A44.05,44.05,0,0,0,96,80V72A4,4,0,0,0,92,68Z\" fill=\"currentColor\"></path><path d=\"M208,119.5l-6-48a4,4,0,0,0-4-3.5H164a4,4,0,0,0-4,4v8a44.06,44.06,0,0,0,44.06,44,4,4,0,0,0,4-4.5Z\" fill=\"currentColor\"></path><path d=\"M186,32a8,8,0,0,1,7.94,7l22,176a8,8,0,0,1-7.94,9H167.31a8,8,0,0,1-7.76-6.06L128,96,96.45,217.94A8,8,0,0,1,88.69,224H48a8,8,0,0,1-7.94-9l22-176A8,8,0,0,1,70,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"128\" y1=\"96\" x2=\"128\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"57.94\" y1=\"72\" x2=\"198.06\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Thin {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M186,32a8,8,0,0,1,7.94,7l22,176a8,8,0,0,1-7.94,9H167.31a8,8,0,0,1-7.76-6.06L128,96,96.45,217.94A8,8,0,0,1,88.69,224H48a8,8,0,0,1-7.94-9l22-176A8,8,0,0,1,70,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"128\" y1=\"96\" x2=\"128\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"58.94\" y1=\"64\" x2=\"197.06\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M96,64a40,40,0,0,1-40,40c-.69,0-1.37,0-2,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M202.06,104c-.68,0-1.37,0-2.06,0a40,40,0,0,1-40-40\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M186,32a8,8,0,0,1,7.94,7l22,176a8,8,0,0,1-7.94,9H167.31a8,8,0,0,1-7.76-6.06L128,96,96.45,217.94A8,8,0,0,1,88.69,224H48a8,8,0,0,1-7.94-9l22-176A8,8,0,0,1,70,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"128\" y1=\"96\" x2=\"128\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"58.94\" y1=\"64\" x2=\"197.06\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M96,64a40,40,0,0,1-40,40c-.69,0-1.37,0-2,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M202.06,104c-.68,0-1.37,0-2.06,0a40,40,0,0,1-40-40\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M53.44,43.5,54.12,38A16,16,0,0,1,70,24H186a16,16,0,0,1,15.88,14l.68,5.49a4,4,0,0,1-4,4.5H57.41A4,4,0,0,1,53.44,43.5ZM169,64a32.06,32.06,0,0,0,31,24h3.59a4,4,0,0,0,4-4.5l-2-16a4,4,0,0,0-4-3.5ZM52.41,88H56A32.06,32.06,0,0,0,87,64H54.41a4,4,0,0,0-4,3.5l-2,16A4,4,0,0,0,52.41,88ZM223.88,214,210.56,107.5a4,4,0,0,0-4-3.5H200a48.07,48.07,0,0,1-47.32-40H136v39.73a8.18,8.18,0,0,1-7.47,8.25,8,8,0,0,1-8.53-8V64H103.32A48.07,48.07,0,0,1,56,104H49.41a4,4,0,0,0-4,3.5L32.12,214a16,16,0,0,0,6.71,15.09A16.56,16.56,0,0,0,48.39,232h40.3a16,16,0,0,0,15.51-12.06l23.8-92,23.79,91.94A16,16,0,0,0,167.31,232h40.3a16.54,16.54,0,0,0,9.56-2.89A16,16,0,0,0,223.88,214Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M186,32a8,8,0,0,1,7.94,7l22,176a8,8,0,0,1-7.94,9H167.31a8,8,0,0,1-7.76-6.06L128,96,96.45,217.94A8,8,0,0,1,88.69,224H48a8,8,0,0,1-7.94-9l22-176A8,8,0,0,1,70,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"128\" y1=\"96\" x2=\"128\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"58.94\" y1=\"64\" x2=\"197.06\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M96,64a40,40,0,0,1-40,40c-.69,0-1.37,0-2,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M202.06,104c-.68,0-1.37,0-2.06,0a40,40,0,0,1-40-40\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
