// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func DiamondsFour(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"99.72\" y=\"35.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-7.76 109.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect><rect x=\"99.72\" y=\"163.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-98.27 146.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect><rect x=\"163.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-34.27 173.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect><rect x=\"35.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-71.76 82.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"99.72\" y=\"35.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-7.76 109.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect><rect x=\"99.72\" y=\"163.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-98.27 146.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect><rect x=\"163.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-34.27 173.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect><rect x=\"35.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-71.76 82.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M82.34,69.66a8,8,0,0,1,0-11.32l40-40a8,8,0,0,1,11.32,0l40,40a8,8,0,0,1,0,11.32l-40,40a8,8,0,0,1-11.32,0Zm51.32,76.68a8,8,0,0,0-11.32,0l-40,40a8,8,0,0,0,0,11.32l40,40a8,8,0,0,0,11.32,0l40-40a8,8,0,0,0,0-11.32Zm104-24-40-40a8,8,0,0,0-11.32,0l-40,40a8,8,0,0,0,0,11.32l40,40a8,8,0,0,0,11.32,0l40-40A8,8,0,0,0,237.66,122.34Zm-128,0-40-40a8,8,0,0,0-11.32,0l-40,40a8,8,0,0,0,0,11.32l40,40a8,8,0,0,0,11.32,0l40-40A8,8,0,0,0,109.66,122.34Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"99.72\" y=\"35.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-7.76 109.25) rotate(-45)\" opacity=\"0.2\"></rect><rect x=\"99.72\" y=\"163.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-98.27 146.75) rotate(-45)\" opacity=\"0.2\"></rect><rect x=\"163.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-34.27 173.25) rotate(-45)\" opacity=\"0.2\"></rect><rect x=\"35.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-71.76 82.75) rotate(-45)\" opacity=\"0.2\"></rect><rect x=\"99.72\" y=\"35.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-7.76 109.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><rect x=\"99.72\" y=\"163.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-98.27 146.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><rect x=\"163.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-34.27 173.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><rect x=\"35.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-71.76 82.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"101.13\" y=\"35.13\" width=\"53.74\" height=\"53.74\" transform=\"translate(-6.35 108.67) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect><rect x=\"101.13\" y=\"167.13\" width=\"53.74\" height=\"53.74\" transform=\"translate(-99.69 147.33) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect><rect x=\"35.13\" y=\"101.13\" width=\"53.74\" height=\"53.74\" transform=\"translate(-72.35 81.33) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect><rect x=\"167.13\" y=\"101.13\" width=\"53.74\" height=\"53.74\" transform=\"translate(-33.69 174.67) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"99.72\" y=\"35.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-7.76 109.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><rect x=\"99.72\" y=\"163.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-98.27 146.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><rect x=\"163.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-34.27 173.25) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><rect x=\"35.72\" y=\"99.72\" width=\"56.57\" height=\"56.57\" transform=\"translate(-71.76 82.75) rotate(-45)\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect>")
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
