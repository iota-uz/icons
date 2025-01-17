// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Racquet(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M230,26.05C202-1.88,151.53,3.16,117.4,37.3c-31.79,31.79-38.33,77.77-16.51,106.49L71.33,173.35l-.68-.68a16,16,0,0,0-22.64,0L20.69,200a16,16,0,0,0,0,22.64l12.69,12.69a16,16,0,0,0,22.63,0h0L83.34,208a16,16,0,0,0,0-22.63l-.69-.69,29.56-29.56c11.29,8.58,25.24,12.79,40,12.79,22.72,0,47.25-10,66.54-29.3C252.83,104.47,257.88,54,230,26.05ZM224.23,104H200.06v-32h32A72.45,72.45,0,0,1,224.23,104ZM136,149.61A44.15,44.15,0,0,1,106.39,120H136ZM104,104a72.24,72.24,0,0,1,7.86-32H136v32Zm48-32h32v32h-32Zm77.67-16H200.06V26.28a44.23,44.23,0,0,1,29.66,29.66Zm-45.82-32h.16v32h-32V31.76A72.47,72.47,0,0,1,183.9,23.9ZM136,42.06V55.94H122.16a89.72,89.72,0,0,1,6.56-7.32A93.17,93.17,0,0,1,136,42.06Zm16,109.92V120h32v24.16A72.24,72.24,0,0,1,152.05,152Zm48-18.14V120H214a91.62,91.62,0,0,1-6.56,7.32A89.64,89.64,0,0,1,200.06,133.84Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><ellipse cx=\"168\" cy=\"88\" rx=\"79.51\" ry=\"63.61\" transform=\"translate(-13.02 144.57) rotate(-45)\" opacity=\"0.2\"></ellipse><ellipse cx=\"168\" cy=\"88\" rx=\"79.51\" ry=\"63.61\" transform=\"translate(-13.02 144.57) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></ellipse><rect x=\"24.69\" y=\"187.03\" width=\"54.63\" height=\"33.94\" rx=\"8\" transform=\"translate(-129.02 96.52) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"71.31\" y1=\"184.69\" x2=\"111.78\" y2=\"144.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"192\" y1=\"16.5\" x2=\"192\" y2=\"148.98\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"144\" y1=\"27.02\" x2=\"144\" y2=\"159.5\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"107.02\" y1=\"64\" x2=\"239.5\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96.5\" y1=\"112\" x2=\"228.98\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><ellipse cx=\"168\" cy=\"88\" rx=\"79.51\" ry=\"63.61\" transform=\"translate(-13.02 144.57) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></ellipse><rect x=\"23.86\" y=\"179.37\" width=\"60.28\" height=\"45.25\" rx=\"8\" transform=\"translate(-127.02 97.35) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect><line x1=\"75.31\" y1=\"180.69\" x2=\"111.78\" y2=\"144.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"196\" y1=\"17.13\" x2=\"196\" y2=\"146.58\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"140\" y1=\"29.41\" x2=\"140\" y2=\"158.87\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"109.42\" y1=\"60\" x2=\"238.87\" y2=\"60\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"97.13\" y1=\"116\" x2=\"226.58\" y2=\"116\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><ellipse cx=\"168\" cy=\"88\" rx=\"79.51\" ry=\"63.61\" transform=\"translate(-13.02 144.57) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></ellipse><rect x=\"24.69\" y=\"187.03\" width=\"54.63\" height=\"33.94\" rx=\"8\" transform=\"translate(-129.02 96.52) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect><line x1=\"71.31\" y1=\"184.69\" x2=\"111.78\" y2=\"144.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"192\" y1=\"16.5\" x2=\"192\" y2=\"148.98\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"144\" y1=\"27.02\" x2=\"144\" y2=\"159.5\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"107.02\" y1=\"64\" x2=\"239.5\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"96.5\" y1=\"112\" x2=\"228.98\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><ellipse cx=\"168\" cy=\"88\" rx=\"79.51\" ry=\"63.61\" transform=\"translate(-13.02 144.57) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></ellipse><rect x=\"24.69\" y=\"187.03\" width=\"54.63\" height=\"33.94\" rx=\"8\" transform=\"translate(-129.02 96.52) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect><line x1=\"71.31\" y1=\"184.69\" x2=\"111.78\" y2=\"144.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"192\" y1=\"16.5\" x2=\"192\" y2=\"148.98\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"144\" y1=\"27.02\" x2=\"144\" y2=\"159.5\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"107.02\" y1=\"64\" x2=\"239.5\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"96.5\" y1=\"112\" x2=\"228.98\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><ellipse cx=\"168\" cy=\"88\" rx=\"79.51\" ry=\"63.61\" transform=\"translate(-13.02 144.57) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></ellipse><rect x=\"24.69\" y=\"187.03\" width=\"54.63\" height=\"33.94\" rx=\"8\" transform=\"translate(-129.02 96.52) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"71.31\" y1=\"184.69\" x2=\"111.78\" y2=\"144.22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"192\" y1=\"16.5\" x2=\"192\" y2=\"148.98\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"144\" y1=\"27.02\" x2=\"144\" y2=\"159.5\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"107.02\" y1=\"64\" x2=\"239.5\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96.5\" y1=\"112\" x2=\"228.98\" y2=\"112\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
