// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Aperture(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M232,128A104,104,0,0,0,54.46,54.46,104,104,0,0,0,128,232h.09A104,104,0,0,0,232,128ZM49.18,88.92l51.21,9.35L46.65,161.53A88.39,88.39,0,0,1,49.18,88.92Zm160.17,5.54a88.41,88.41,0,0,1-2.53,72.62l-51.21-9.35Zm-8.08-15.2L167.55,119,139.63,40.78a87.38,87.38,0,0,1,50.6,25A88.74,88.74,0,0,1,201.27,79.26ZM122.43,40.19l17.51,49L58.3,74.32a89.28,89.28,0,0,1,7.47-8.55A87.37,87.37,0,0,1,122.43,40.19ZM54.73,176.74,88.45,137l27.92,78.18a88,88,0,0,1-61.64-38.48Zm78.84,39.06-17.51-49L139.14,171h0l58.52,10.69a87.5,87.5,0,0,1-64.13,34.12Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M195.88,60.12a96,96,0,1,0,0,135.76A96,96,0,0,0,195.88,60.12Zm-31.2,74.58-24.14,28.41h0l-36.68-6.69h0L91.32,121.3h0l24.14-28.41h0l36.68,6.69h0l12.54,35.12Z\" opacity=\"0.2\"></path><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><line x1=\"128\" y1=\"32\" x2=\"164.68\" y2=\"134.7\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"44.86\" y1=\"80\" x2=\"152.14\" y2=\"99.58\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"44.86\" y1=\"176\" x2=\"115.46\" y2=\"92.89\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"224\" x2=\"91.32\" y2=\"121.3\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"211.14\" y1=\"176\" x2=\"103.86\" y2=\"156.42\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"211.14\" y1=\"80\" x2=\"140.54\" y2=\"163.11\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></circle><line x1=\"128\" y1=\"32\" x2=\"164.68\" y2=\"134.7\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"44.86\" y1=\"80\" x2=\"152.14\" y2=\"99.58\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"44.86\" y1=\"176\" x2=\"115.46\" y2=\"92.89\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"128\" y1=\"224\" x2=\"91.32\" y2=\"121.3\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"211.14\" y1=\"176\" x2=\"103.86\" y2=\"156.42\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"211.14\" y1=\"80\" x2=\"140.54\" y2=\"163.11\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></circle><line x1=\"128\" y1=\"32\" x2=\"164.68\" y2=\"134.7\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"44.86\" y1=\"80\" x2=\"152.14\" y2=\"99.58\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"44.86\" y1=\"176\" x2=\"115.46\" y2=\"92.89\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"128\" y1=\"224\" x2=\"91.32\" y2=\"121.3\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"211.14\" y1=\"176\" x2=\"103.86\" y2=\"156.42\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"211.14\" y1=\"80\" x2=\"140.54\" y2=\"163.11\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></circle><line x1=\"128\" y1=\"32\" x2=\"164.68\" y2=\"134.7\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"44.86\" y1=\"80\" x2=\"152.14\" y2=\"99.58\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"44.86\" y1=\"176\" x2=\"115.46\" y2=\"92.89\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"128\" y1=\"224\" x2=\"91.32\" y2=\"121.3\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"211.14\" y1=\"176\" x2=\"103.86\" y2=\"156.42\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"211.14\" y1=\"80\" x2=\"140.54\" y2=\"163.11\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><line x1=\"128\" y1=\"32\" x2=\"164.68\" y2=\"134.7\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"44.86\" y1=\"80\" x2=\"152.14\" y2=\"99.58\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"44.86\" y1=\"176\" x2=\"115.46\" y2=\"92.89\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"224\" x2=\"91.32\" y2=\"121.3\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"211.14\" y1=\"176\" x2=\"103.86\" y2=\"156.42\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"211.14\" y1=\"80\" x2=\"140.54\" y2=\"163.11\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
