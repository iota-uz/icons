// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Engine(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M224,176l-37.66,37.66a8,8,0,0,1-5.65,2.34H103.31a8,8,0,0,1-5.65-2.34L58.34,174.34A8,8,0,0,1,56,168.69V80a8,8,0,0,1,8-8H180.69a8,8,0,0,1,5.65,2.34L224,112h16a8,8,0,0,1,8,8v48a8,8,0,0,1-8,8Z\" opacity=\"0.2\" fill=\"currentColor\"></path><line x1=\"16\" y1=\"108\" x2=\"16\" y2=\"172\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"100\" y1=\"32\" x2=\"164\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M224,176l-37.66,37.66a8,8,0,0,1-5.65,2.34H103.31a8,8,0,0,1-5.65-2.34L58.34,174.34A8,8,0,0,1,56,168.69V80a8,8,0,0,1,8-8H180.69a8,8,0,0,1,5.65,2.34L224,112h16a8,8,0,0,1,8,8v48a8,8,0,0,1-8,8Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"16\" y1=\"140\" x2=\"56\" y2=\"140\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"132\" y1=\"32\" x2=\"132\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"16\" y1=\"108\" x2=\"16\" y2=\"172\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"100\" y1=\"32\" x2=\"164\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M220,176l-37.66,37.66a8,8,0,0,1-5.65,2.34H103.31a8,8,0,0,1-5.65-2.34L58.34,174.34A8,8,0,0,1,56,168.69V80a8,8,0,0,1,8-8H176.69a8,8,0,0,1,5.65,2.34L220,112h16a8,8,0,0,1,8,8v48a8,8,0,0,1-8,8Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"16\" y1=\"140\" x2=\"56\" y2=\"140\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"132\" y1=\"32\" x2=\"132\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Thin {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"16\" y1=\"108\" x2=\"16\" y2=\"172\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"100\" y1=\"32\" x2=\"164\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M224,176l-37.66,37.66a8,8,0,0,1-5.65,2.34H103.31a8,8,0,0,1-5.65-2.34L58.34,174.34A8,8,0,0,1,56,168.69V80a8,8,0,0,1,8-8H180.69a8,8,0,0,1,5.65,2.34L224,112h16a8,8,0,0,1,8,8v48a8,8,0,0,1-8,8Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"16\" y1=\"140\" x2=\"56\" y2=\"140\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"132\" y1=\"32\" x2=\"132\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"16\" y1=\"108\" x2=\"16\" y2=\"172\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"100\" y1=\"32\" x2=\"164\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M224,176l-37.66,37.66a8,8,0,0,1-5.65,2.34H103.31a8,8,0,0,1-5.65-2.34L58.34,174.34A8,8,0,0,1,56,168.69V80a8,8,0,0,1,8-8H180.69a8,8,0,0,1,5.65,2.34L224,112h16a8,8,0,0,1,8,8v48a8,8,0,0,1-8,8Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"16\" y1=\"140\" x2=\"56\" y2=\"140\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"132\" y1=\"32\" x2=\"132\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M256,120v48a16,16,0,0,1-16,16H227.31L192,219.31A15.86,15.86,0,0,1,180.69,224H103.31A15.86,15.86,0,0,1,92,219.31L52.69,180A15.86,15.86,0,0,1,48,168.69V148H24v24a8,8,0,0,1-16,0V108a8,8,0,0,1,16,0v24H48V80A16,16,0,0,1,64,64h60V40H100a8,8,0,0,1,0-16h64a8,8,0,0,1,0,16H140V64h40.69A15.86,15.86,0,0,1,192,68.69L227.31,104H240A16,16,0,0,1,256,120Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"16\" y1=\"108\" x2=\"16\" y2=\"172\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"100\" y1=\"32\" x2=\"164\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M224,176l-37.66,37.66a8,8,0,0,1-5.65,2.34H103.31a8,8,0,0,1-5.65-2.34L58.34,174.34A8,8,0,0,1,56,168.69V80a8,8,0,0,1,8-8H180.69a8,8,0,0,1,5.65,2.34L224,112h16a8,8,0,0,1,8,8v48a8,8,0,0,1-8,8Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"16\" y1=\"140\" x2=\"56\" y2=\"140\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"132\" y1=\"32\" x2=\"132\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
