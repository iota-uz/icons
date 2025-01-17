// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Hoodie(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M237.31,120.53,183,39.12A16,16,0,0,0,169.73,32H86.27A16,16,0,0,0,73,39.12L18.69,120.53a16,16,0,0,0-2.13,13.09L38,212.21A16,16,0,0,0,53.43,224H80a16,16,0,0,0,16-16V192h64v16a16,16,0,0,0,16,16h26.57A16,16,0,0,0,218,212.21l21.44-78.59A16,16,0,0,0,237.31,120.53ZM80,208H53.43L32,129.41l32-48V176a16,16,0,0,0,16,16Zm40-72a8,8,0,0,1-16,0V97.14a8,8,0,1,1,16,0Zm32-8a8,8,0,0,1-16,0V97.14a8,8,0,1,1,16,0ZM128,78.71,83.35,52.39,86.27,48h83.46l2.92,4.39ZM202.57,208H176V192a16,16,0,0,0,16-16V81.41l32,48Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M230.66,125,184,55h0L128,88,72,55h0L25.34,125a8,8,0,0,0-1.06,6.54L45.72,210.1a8,8,0,0,0,7.71,5.9H80a8,8,0,0,0,8-8V184h80v24a8,8,0,0,0,8,8h26.57a8,8,0,0,0,7.71-5.9l21.44-78.59A8,8,0,0,0,230.66,125Z\" opacity=\"0.2\"></path><path d=\"M168,184v24a8,8,0,0,0,8,8h26.57a8,8,0,0,0,7.71-5.9l21.44-78.59a8,8,0,0,0-1.06-6.54L176.38,43.56A8,8,0,0,0,169.73,40H86.27a8,8,0,0,0-6.65,3.56L25.34,125a8,8,0,0,0-1.06,6.54L45.72,210.1a8,8,0,0,0,7.71,5.9H80a8,8,0,0,0,8-8V184\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M128,88l56-33V176a8,8,0,0,1-8,8H80a8,8,0,0,1-8-8V55Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"144\" y1=\"78.57\" x2=\"144\" y2=\"128\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"112\" y1=\"78.57\" x2=\"112\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M172,184v24a8,8,0,0,0,8,8h26.57a8,8,0,0,0,7.71-5.9l21.44-78.59a8,8,0,0,0-1.06-6.54L180.38,43.56A8,8,0,0,0,173.73,40H82.27a8,8,0,0,0-6.65,3.56L21.34,125a8,8,0,0,0-1.06,6.54L41.72,210.1a8,8,0,0,0,7.71,5.9H76a8,8,0,0,0,8-8V184\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M128,84l56-35V176a8,8,0,0,1-8,8H80a8,8,0,0,1-8-8V49Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"148\" y1=\"71.49\" x2=\"148\" y2=\"128\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"108\" y1=\"71.49\" x2=\"108\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M168,184v24a8,8,0,0,0,8,8h26.57a8,8,0,0,0,7.71-5.9l21.44-78.59a8,8,0,0,0-1.06-6.54L176.38,43.56A8,8,0,0,0,169.73,40H86.27a8,8,0,0,0-6.65,3.56L25.34,125a8,8,0,0,0-1.06,6.54L45.72,210.1a8,8,0,0,0,7.71,5.9H80a8,8,0,0,0,8-8V184\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M128,88l56-33V176a8,8,0,0,1-8,8H80a8,8,0,0,1-8-8V55Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"144\" y1=\"78.57\" x2=\"144\" y2=\"128\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"112\" y1=\"78.57\" x2=\"112\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M168,184v24a8,8,0,0,0,8,8h26.57a8,8,0,0,0,7.71-5.9l21.44-78.59a8,8,0,0,0-1.06-6.54L176.38,43.56A8,8,0,0,0,169.73,40H86.27a8,8,0,0,0-6.65,3.56L25.34,125a8,8,0,0,0-1.06,6.54L45.72,210.1a8,8,0,0,0,7.71,5.9H80a8,8,0,0,0,8-8V184\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M128,88l56-33V176a8,8,0,0,1-8,8H80a8,8,0,0,1-8-8V55Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"144\" y1=\"78.57\" x2=\"144\" y2=\"128\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"112\" y1=\"78.57\" x2=\"112\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M168,184v24a8,8,0,0,0,8,8h26.57a8,8,0,0,0,7.71-5.9l21.44-78.59a8,8,0,0,0-1.06-6.54L176.38,43.56A8,8,0,0,0,169.73,40H86.27a8,8,0,0,0-6.65,3.56L25.34,125a8,8,0,0,0-1.06,6.54L45.72,210.1a8,8,0,0,0,7.71,5.9H80a8,8,0,0,0,8-8V184\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M128,88l56-33V176a8,8,0,0,1-8,8H80a8,8,0,0,1-8-8V55Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"144\" y1=\"78.57\" x2=\"144\" y2=\"128\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"112\" y1=\"78.57\" x2=\"112\" y2=\"136\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
