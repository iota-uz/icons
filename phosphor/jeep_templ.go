// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Jeep(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M224,168v32a8,8,0,0,1-8,8H192a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M72,168v32a8,8,0,0,1-8,8H40a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"16\" y1=\"96\" x2=\"240\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"144\" y1=\"128\" x2=\"144\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"112\" y1=\"128\" x2=\"112\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M224,168H32V96L42.64,46.32A8,8,0,0,1,50.47,40H205.53a8,8,0,0,1,7.83,6.32L224,96Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><circle cx=\"68\" cy=\"132\" r=\"8\"></circle><circle cx=\"188\" cy=\"132\" r=\"8\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M224,168v32a8,8,0,0,1-8,8H192a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M72,168v32a8,8,0,0,1-8,8H40a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"16\" y1=\"96\" x2=\"240\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"144\" y1=\"128\" x2=\"144\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"112\" y1=\"128\" x2=\"112\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M224,168H32V96L42.64,46.32A8,8,0,0,1,50.47,40H205.53a8,8,0,0,1,7.83,6.32L224,96Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><circle cx=\"68\" cy=\"132\" r=\"10\"></circle><circle cx=\"188\" cy=\"132\" r=\"10\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M248,103.47A8.17,8.17,0,0,0,239.73,96H232a8,8,0,0,0-.18-1.68L221.18,44.65A16.08,16.08,0,0,0,205.53,32H50.47A16.08,16.08,0,0,0,34.82,44.65L24.18,94.32A8,8,0,0,0,24,96H16.27A8.17,8.17,0,0,0,8,103.47,8,8,0,0,0,16,112h8v88a16,16,0,0,0,16,16H64a16,16,0,0,0,16-16V184h20a4,4,0,0,0,4-4V128.27a8.17,8.17,0,0,1,7.47-8.25,8,8,0,0,1,8.53,8v52a4,4,0,0,0,4,4h8a4,4,0,0,0,4-4V128.27a8.17,8.17,0,0,1,7.47-8.25,8,8,0,0,1,8.53,8v52a4,4,0,0,0,4,4h20v16a16,16,0,0,0,16,16h24a16,16,0,0,0,16-16V112h8A8,8,0,0,0,248,103.47ZM68,144a12,12,0,1,1,12-12A12,12,0,0,1,68,144Zm120,0a12,12,0,1,1,12-12A12,12,0,0,1,188,144ZM40.18,96,50.47,48H205.53l10.29,48Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M32,96,42.64,46.32A8,8,0,0,1,50.47,40H205.53a8,8,0,0,1,7.83,6.32L224,96Z\" opacity=\"0.2\" fill=\"currentColor\"></path><path d=\"M224,168v32a8,8,0,0,1-8,8H192a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M72,168v32a8,8,0,0,1-8,8H40a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"16\" y1=\"96\" x2=\"240\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"144\" y1=\"128\" x2=\"144\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"112\" y1=\"128\" x2=\"112\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M224,168H32V96L42.64,46.32A8,8,0,0,1,50.47,40H205.53a8,8,0,0,1,7.83,6.32L224,96Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><circle cx=\"68\" cy=\"132\" r=\"12\"></circle><circle cx=\"188\" cy=\"132\" r=\"12\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M224,168v32a8,8,0,0,1-8,8H188a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M76,168v32a8,8,0,0,1-8,8H40a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"16\" y1=\"96\" x2=\"240\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"128\" y1=\"132\" x2=\"128\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"168\" y1=\"132\" x2=\"168\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"88\" y1=\"132\" x2=\"88\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M224,168H32V96L42.64,46.32A8,8,0,0,1,50.47,40H205.53a8,8,0,0,1,7.83,6.32L224,96Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M224,168v32a8,8,0,0,1-8,8H192a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M72,168v32a8,8,0,0,1-8,8H40a8,8,0,0,1-8-8V168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"16\" y1=\"96\" x2=\"240\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"144\" y1=\"128\" x2=\"144\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"112\" y1=\"128\" x2=\"112\" y2=\"168\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M224,168H32V96L42.64,46.32A8,8,0,0,1,50.47,40H205.53a8,8,0,0,1,7.83,6.32L224,96Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><circle cx=\"68\" cy=\"132\" r=\"12\"></circle><circle cx=\"188\" cy=\"132\" r=\"12\"></circle>")
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
