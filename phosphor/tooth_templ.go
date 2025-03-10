// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Tooth(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"64\" x2=\"168\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"168\" y1=\"64\" x2=\"88\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M40,79.76C39.93,151.18,63.48,208,84,222.54a8,8,0,0,0,12.52-6C98,196.31,104,160,128,160s30,36.31,31.51,56.53a8,8,0,0,0,12.52,6c20.49-14.55,44-71.36,44-142.78A47.94,47.94,0,0,0,168,32H88A47.94,47.94,0,0,0,40,79.76Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"64\" x2=\"168\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"168\" y1=\"64\" x2=\"88\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M40,79.76C39.93,151.18,63.48,208,84,222.54a8,8,0,0,0,12.52-6C98,196.31,104,160,128,160s30,36.31,31.51,56.53a8,8,0,0,0,12.52,6c20.49-14.55,44-71.36,44-142.78A47.94,47.94,0,0,0,168,32H88A47.94,47.94,0,0,0,40,79.76Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M168,24H88A56,56,0,0,0,32,79.75c0,42.72,8,75.4,14.7,95.28,8.72,25.8,20.62,45.49,32.64,54A15.67,15.67,0,0,0,88.47,232a16.09,16.09,0,0,0,16-14.9c.85-11.52,5-49.11,23.53-49.11s22.68,37.59,23.53,49.11a16.09,16.09,0,0,0,9.18,13.36,15.69,15.69,0,0,0,15.95-1.41c12-8.53,23.92-28.22,32.64-54C216,155.15,224,122.47,224,79.75A56,56,0,0,0,168,24Zm3,56.57A8,8,0,1,1,165,95.42L128,80.61,91,95.42A8,8,0,1,1,85,80.57L106.46,72,85,63.42A8,8,0,1,1,91,48.57l37,14.81,37-14.81A8,8,0,1,1,171,63.42L149.54,72Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M40,79.76C39.93,151.18,63.48,208,84,222.54a8,8,0,0,0,12.52-6C98,196.31,104,160,128,160s30,36.31,31.51,56.53a8,8,0,0,0,12.52,6c20.49-14.55,44-71.36,44-142.78A47.94,47.94,0,0,0,168,32H88A47.94,47.94,0,0,0,40,79.76Z\" opacity=\"0.2\" fill=\"currentColor\"></path><line x1=\"88\" y1=\"64\" x2=\"168\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"168\" y1=\"64\" x2=\"88\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M40,79.76C39.93,151.18,63.48,208,84,222.54a8,8,0,0,0,12.52-6C98,196.31,104,160,128,160s30,36.31,31.51,56.53a8,8,0,0,0,12.52,6c20.49-14.55,44-71.36,44-142.78A47.94,47.94,0,0,0,168,32H88A47.94,47.94,0,0,0,40,79.76Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"72\" x2=\"168\" y2=\"104\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"168\" y1=\"72\" x2=\"88\" y2=\"104\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M40,79.76C39.93,151.18,63.48,208,84,222.54a8,8,0,0,0,12.52-6C98,196.31,104,160,128,160s30,36.31,31.51,56.53a8,8,0,0,0,12.52,6c20.49-14.55,44-71.36,44-142.78A47.94,47.94,0,0,0,168,32H88A47.94,47.94,0,0,0,40,79.76Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"64\" x2=\"168\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"168\" y1=\"64\" x2=\"88\" y2=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M40,79.76C39.93,151.18,63.48,208,84,222.54a8,8,0,0,0,12.52-6C98,196.31,104,160,128,160s30,36.31,31.51,56.53a8,8,0,0,0,12.52,6c20.49-14.55,44-71.36,44-142.78A47.94,47.94,0,0,0,168,32H88A47.94,47.94,0,0,0,40,79.76Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
