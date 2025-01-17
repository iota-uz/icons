// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Champagne(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"96\" y1=\"240\" x2=\"144\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M97.7,16h44.6s52,160-22.3,160S97.7,16,97.7,16Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"120\" y1=\"176\" x2=\"120\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><circle cx=\"220\" cy=\"52\" r=\"8\"></circle><circle cx=\"196\" cy=\"20\" r=\"8\"></circle><circle cx=\"196\" cy=\"100\" r=\"8\"></circle><line x1=\"84.2\" y1=\"72\" x2=\"155.8\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"96\" y1=\"240\" x2=\"144\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M97.7,16h44.6s52,160-22.3,160S97.7,16,97.7,16Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"120\" y1=\"176\" x2=\"120\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><circle cx=\"220\" cy=\"52\" r=\"10\"></circle><circle cx=\"196\" cy=\"20\" r=\"10\"></circle><circle cx=\"196\" cy=\"100\" r=\"10\"></circle><line x1=\"84.2\" y1=\"72\" x2=\"155.8\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M149.91,13.53A8,8,0,0,0,142.3,8H97.71a8,8,0,0,0-7.61,5.53,451,451,0,0,0-14.21,59.7c-7.26,44.25-4.35,75.76,8.65,93.66A40,40,0,0,0,112,183.42V232H96a8,8,0,1,0,0,16h48a8,8,0,0,0,0-16H128V183.42a39.94,39.94,0,0,0,27.46-16.53c13-17.9,15.92-49.41,8.66-93.66A451,451,0,0,0,149.91,13.53ZM93.8,64c3-15.58,6.73-29.81,9.79-40h32.83c3.06,10.19,6.77,24.42,9.8,40ZM232,52a12,12,0,1,1-12-12A12,12,0,0,1,232,52ZM184,20a12,12,0,1,1,12,12A12,12,0,0,1,184,20Zm24,80a12,12,0,1,1-12-12A12,12,0,0,1,208,100Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M84.2,72c-7.9,46.13-8.9,104,35.8,104s43.7-57.87,35.8-104Z\" opacity=\"0.2\"></path><line x1=\"96\" y1=\"240\" x2=\"144\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M97.7,16h44.6s52,160-22.3,160S97.7,16,97.7,16Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"120\" y1=\"176\" x2=\"120\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"220\" cy=\"52\" r=\"12\"></circle><circle cx=\"196\" cy=\"20\" r=\"12\"></circle><circle cx=\"196\" cy=\"100\" r=\"12\"></circle><line x1=\"84.2\" y1=\"72\" x2=\"155.8\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"96\" y1=\"240\" x2=\"144\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M97.7,16h44.6s52,160-22.3,160S97.7,16,97.7,16Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"120\" y1=\"176\" x2=\"120\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><circle cx=\"224\" cy=\"56\" r=\"16\"></circle><circle cx=\"196\" cy=\"20\" r=\"16\"></circle><circle cx=\"200\" cy=\"104\" r=\"16\"></circle><line x1=\"84.2\" y1=\"72\" x2=\"155.8\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"96\" y1=\"240\" x2=\"144\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M97.7,16h44.6s52,160-22.3,160S97.7,16,97.7,16Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"120\" y1=\"176\" x2=\"120\" y2=\"240\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"220\" cy=\"52\" r=\"12\"></circle><circle cx=\"196\" cy=\"20\" r=\"12\"></circle><circle cx=\"196\" cy=\"100\" r=\"12\"></circle><line x1=\"84.2\" y1=\"72\" x2=\"155.8\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
