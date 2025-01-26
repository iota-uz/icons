// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Crop(props Props) templ.Component {
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
			if props.Variant == Thin {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><polyline points=\"64 24 64 192 232 192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></polyline><line x1=\"24\" y1=\"64\" x2=\"64\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><polyline points=\"96 64 192 64 192 160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></polyline><line x1=\"192\" y1=\"192\" x2=\"192\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><polyline points=\"64 24 64 192 232 192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></polyline><line x1=\"24\" y1=\"64\" x2=\"64\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><polyline points=\"96 64 192 64 192 160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></polyline><line x1=\"192\" y1=\"192\" x2=\"192\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M208,32H48A16,16,0,0,0,32,48V208a16,16,0,0,0,16,16H208a16,16,0,0,0,16-16V48A16,16,0,0,0,208,32ZM120,80h48a8,8,0,0,1,8,8v48a8,8,0,0,1-16,0V96H120a8,8,0,0,1,0-16Zm72,96H176v16a8,8,0,0,1-16,0V176H88a8,8,0,0,1-8-8V96H64a8,8,0,0,1,0-16H80V64a8,8,0,0,1,16,0v96h96a8,8,0,0,1,0,16Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"64\" y=\"64\" width=\"128\" height=\"128\" opacity=\"0.2\"></rect><polyline points=\"64 24 64 192 232 192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></polyline><line x1=\"24\" y1=\"64\" x2=\"64\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><polyline points=\"96 64 192 64 192 160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></polyline><line x1=\"192\" y1=\"192\" x2=\"192\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><polyline points=\"64 24 64 192 232 192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></polyline><line x1=\"24\" y1=\"64\" x2=\"64\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><polyline points=\"104 64 192 64 192 152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></polyline><line x1=\"192\" y1=\"192\" x2=\"192\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><polyline points=\"64 24 64 192 232 192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></polyline><line x1=\"24\" y1=\"64\" x2=\"64\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><polyline points=\"96 64 192 64 192 160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></polyline><line x1=\"192\" y1=\"192\" x2=\"192\" y2=\"232\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
