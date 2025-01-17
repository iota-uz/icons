// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CowboyHat(props Props) templ.Component {
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
			if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M40,128c40,64,88,64,88,64s48,0,88-64a32,32,0,0,1,0,64H40a32,32,0,0,1,0-64Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M65.47,160.18,85.54,54.53A8,8,0,0,1,98.4,49.76L123,70.25a8,8,0,0,0,10,0l24.6-20.49a8,8,0,0,1,12.86,4.77l20.07,105.65\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"70.06\" y1=\"136\" x2=\"185.94\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M216,120a8,8,0,0,0-6.78,3.76A179.9,179.9,0,0,1,195.41,143l-1.63-8.57v0L178.32,53.07a16,16,0,0,0-25.72-9.55l-.13.1L128,64,103.53,43.62l-.13-.1a16,16,0,0,0-25.72,9.53L62.23,134.38v0L60.59,143a179.27,179.27,0,0,1-13.81-19.25A8,8,0,0,0,40,120a40,40,0,0,0,0,80H216a40,40,0,0,0,0-80ZM76.68,144H179.31l2.54,13.35a113.28,113.28,0,0,1-27.35,19C139.1,183.77,128.06,184,128,184c-.33,0-25.49-.4-53.86-26.6Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M65.47,160.18C97.87,192,128,192,128,192s30.13,0,62.53-31.82L185.94,136H70.06Z\" opacity=\"0.2\"></path><path d=\"M40,128c40,64,88,64,88,64s48,0,88-64a32,32,0,0,1,0,64H40a32,32,0,0,1,0-64Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M65.47,160.18,85.54,54.53A8,8,0,0,1,98.4,49.76L123,70.25a8,8,0,0,0,10,0l24.6-20.49a8,8,0,0,1,12.86,4.77l20.07,105.65\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"70.06\" y1=\"136\" x2=\"185.94\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M44,128c40,64,84,64,84,64s44,0,84-64a32,32,0,0,1,0,64H44a32,32,0,0,1,0-64Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"70.06\" y1=\"136\" x2=\"185.94\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M66.05,157.08,85.54,54.53A8,8,0,0,1,98.4,49.76L123,70.25a8,8,0,0,0,10,0l24.6-20.49a8,8,0,0,1,12.86,4.77L190,157.08\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M40,128c40,64,88,64,88,64s48,0,88-64a32,32,0,0,1,0,64H40a32,32,0,0,1,0-64Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M65.47,160.18,85.54,54.53A8,8,0,0,1,98.4,49.76L123,70.25a8,8,0,0,0,10,0l24.6-20.49a8,8,0,0,1,12.86,4.77l20.07,105.65\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"70.06\" y1=\"136\" x2=\"185.94\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M40,128c40,64,88,64,88,64s48,0,88-64a32,32,0,0,1,0,64H40a32,32,0,0,1,0-64Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M65.47,160.18,85.54,54.53A8,8,0,0,1,98.4,49.76L123,70.25a8,8,0,0,0,10,0l24.6-20.49a8,8,0,0,1,12.86,4.77l20.07,105.65\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"70.06\" y1=\"136\" x2=\"185.94\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
