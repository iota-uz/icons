// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func PottedPlant(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M184,152l-14.61,65.74a8,8,0,0,1-7.81,6.26H94.42a8,8,0,0,1-7.81-6.26L72,152Z\" opacity=\"0.2\" fill=\"currentColor\"></path><path d=\"M144.28,111.72c-25.08-41.81,8.36-83.61,79.43-79.43C227.89,103.36,186.09,136.8,144.28,111.72Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M98,114c18.24-30.41-6.08-60.81-57.77-57.77C37.17,107.9,67.57,132.22,98,114Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"56\" y1=\"152\" x2=\"200\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M184,152l-14.61,65.74a8,8,0,0,1-7.81,6.26H94.42a8,8,0,0,1-7.81-6.26L72,152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"144.28\" y1=\"111.72\" x2=\"104\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"97.98\" y1=\"113.98\" x2=\"120\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M147.9,108.1c-23.94-39.91,8-79.81,75.82-75.82C227.71,100.11,187.81,132,147.9,108.1Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M94.35,110.35c17.11-28.5-5.7-57-54.15-54.15C37.35,104.65,65.85,127.46,94.35,110.35Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"56\" y1=\"156\" x2=\"200\" y2=\"156\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M183.11,156l-13.72,61.74a8,8,0,0,1-7.81,6.26H94.42a8,8,0,0,1-7.81-6.26L72.89,156\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"147.9\" y1=\"108.1\" x2=\"100\" y2=\"156\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"94.35\" y1=\"110.35\" x2=\"120\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Thin {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M144.28,111.72c-25.08-41.81,8.36-83.61,79.43-79.43C227.89,103.36,186.09,136.8,144.28,111.72Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M98,114c18.24-30.41-6.08-60.81-57.77-57.77C37.17,107.9,67.57,132.22,98,114Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"56\" y1=\"152\" x2=\"200\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M184,152l-14.61,65.74a8,8,0,0,1-7.81,6.26H94.42a8,8,0,0,1-7.81-6.26L72,152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"144.28\" y1=\"111.72\" x2=\"104\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"97.98\" y1=\"113.98\" x2=\"120\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M144.28,111.72c-25.08-41.81,8.36-83.61,79.43-79.43C227.89,103.36,186.09,136.8,144.28,111.72Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M98,114c18.24-30.41-6.08-60.81-57.77-57.77C37.17,107.9,67.57,132.22,98,114Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"56\" y1=\"152\" x2=\"200\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M184,152l-14.61,65.74a8,8,0,0,1-7.81,6.26H94.42a8,8,0,0,1-7.81-6.26L72,152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"144.28\" y1=\"111.72\" x2=\"104\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"97.98\" y1=\"113.98\" x2=\"120\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M200,144h-76.7l22.41-22.41a59.55,59.55,0,0,0,26.1,6.36,49.56,49.56,0,0,0,25.89-7.22c23.72-14.36,36.43-47.6,34-88.92a8,8,0,0,0-7.52-7.52c-41.32-2.43-74.56,10.28-88.93,34-9.35,15.45-9.59,34.11-.86,52L120,124.68l-12.21-12.21c6-13.25,5.57-27-1.39-38.48C95.53,56,70.61,46.41,39.73,48.22a8,8,0,0,0-7.51,7.51C30.4,86.6,40,111.52,58,122.4A38.22,38.22,0,0,0,78,128a45,45,0,0,0,18.52-4.19L108.69,136l-8,8H56a8,8,0,0,0,0,16h9.59L78.8,219.47A15.89,15.89,0,0,0,94.42,232h67.17a15.91,15.91,0,0,0,15.62-12.53L190.42,160H200a8,8,0,0,0,0-16Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M144.28,111.72c-25.08-41.81,8.36-83.61,79.43-79.43C227.89,103.36,186.09,136.8,144.28,111.72Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M98,114c18.24-30.41-6.08-60.81-57.77-57.77C37.17,107.9,67.57,132.22,98,114Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"56\" y1=\"152\" x2=\"200\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M184,152l-14.61,65.74a8,8,0,0,1-7.81,6.26H94.42a8,8,0,0,1-7.81-6.26L72,152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"144.28\" y1=\"111.72\" x2=\"104\" y2=\"152\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"97.98\" y1=\"113.98\" x2=\"120\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
