// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Yarn(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"224\" x2=\"232\" y2=\"223.97\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></circle><path d=\"M92.3,217.14A176.3,176.3,0,0,1,214.81,87\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M57.34,193A216.52,216.52,0,0,1,188.59,53.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M131.8,223.93A136.3,136.3,0,0,1,224,126\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M32.31,120.25a135.07,135.07,0,0,1,53,14.33\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M44.19,81.15a175,175,0,0,1,67.17,22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M77.12,46.58A215,215,0,0,1,144,76.37\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"224\" x2=\"232\" y2=\"223.97\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></circle><path d=\"M92.3,217.14A176.3,176.3,0,0,1,214.81,87\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M57.34,193A216.52,216.52,0,0,1,188.59,53.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M131.8,223.93A136.3,136.3,0,0,1,224,126\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M32.31,120.25a135.07,135.07,0,0,1,53,14.33\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M44.19,81.15a175,175,0,0,1,67.17,22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M77.12,46.58A215,215,0,0,1,144,76.37\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M69.4,144.89a220.94,220.94,0,0,0-18.77,42.6,4,4,0,0,1-7,1.19,103.44,103.44,0,0,1-18.83-48.11,4,4,0,0,1,4.13-4.47A119,119,0,0,1,69.4,144.89ZM44,77.14a180.06,180.06,0,0,1,63,19.12,226.9,226.9,0,0,1,22.6-19.49,206.42,206.42,0,0,0-69.28-25.6,4,4,0,0,0-3.42,1A105.3,105.3,0,0,0,41.08,71,4,4,0,0,0,44,77.14ZM239.58,221.35A7.91,7.91,0,0,0,232,216H183.36A104.25,104.25,0,0,0,230.25,147a4,4,0,0,0-5.27-4.52A120.6,120.6,0,0,0,150.88,216H134a136.55,136.55,0,0,1,94.78-91.37,4,4,0,0,0,2.92-4.15,102.59,102.59,0,0,0-3.58-20.56,4,4,0,0,0-4.89-2.8A164.53,164.53,0,0,0,103,225a4,4,0,0,0,3.08,4.69A103.9,103.9,0,0,0,128,232h0l104,0A8,8,0,0,0,239.58,221.35Zm-211-101.27a134.51,134.51,0,0,1,49.39,11A224.44,224.44,0,0,1,95.52,108.4,164.28,164.28,0,0,0,33.36,92.28a4,4,0,0,0-4,2.75,103,103,0,0,0-4.63,20.61A4,4,0,0,0,28.57,120.08ZM85.84,40.66A222.81,222.81,0,0,1,144,66.8a221.3,221.3,0,0,1,38.8-19.67,4,4,0,0,0,.7-7.08,103.86,103.86,0,0,0-98.2-6.85A4,4,0,0,0,85.84,40.66ZM216,82.51a4,4,0,0,0,2.4-5.87,105,105,0,0,0-12.82-17.81,4,4,0,0,0-4.21-1.19A208.81,208.81,0,0,0,62.21,205.51a4,4,0,0,0,1.44,4.13A104.25,104.25,0,0,0,82.2,221.36a4,4,0,0,0,5.71-2.75A180.61,180.61,0,0,1,216,82.51Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"128\" cy=\"128\" r=\"96\" opacity=\"0.2\"></circle><line x1=\"128\" y1=\"224\" x2=\"232\" y2=\"223.97\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M92.3,217.14A176.3,176.3,0,0,1,214.81,87\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M57.34,193A216.52,216.52,0,0,1,188.59,53.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M131.8,223.93A136.3,136.3,0,0,1,224,126\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M32.31,120.25a135.07,135.07,0,0,1,53,14.33\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M44.19,81.15a175,175,0,0,1,67.17,22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M77.12,46.58A215,215,0,0,1,144,76.37\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"224\" x2=\"232\" y2=\"223.97\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></circle><path d=\"M92.3,217.14A176.3,176.3,0,0,1,214.81,87\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M57.34,193A216.52,216.52,0,0,1,188.59,53.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M131.8,223.93A136.3,136.3,0,0,1,224,126\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M32.31,120.25a135.07,135.07,0,0,1,53,14.33\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M44.19,81.15a175,175,0,0,1,67.17,22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M77.12,46.58A215,215,0,0,1,144,76.37\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"224\" x2=\"232\" y2=\"223.97\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"128\" cy=\"128\" r=\"96\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M92.3,217.14A176.3,176.3,0,0,1,214.81,87\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M57.34,193A216.52,216.52,0,0,1,188.59,53.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M131.8,223.93A136.3,136.3,0,0,1,224,126\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M32.31,120.25a135.07,135.07,0,0,1,53,14.33\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M44.19,81.15a175,175,0,0,1,67.17,22\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M77.12,46.58A215,215,0,0,1,144,76.37\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
