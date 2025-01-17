// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func LastfmLogo(props Props) templ.Component {
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
			if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M108.67,168A40,40,0,0,1,72,192H64a48,48,0,0,1-48-48V120A48,48,0,0,1,64,72h9.43a48,48,0,0,1,43.5,27.7l30.14,64.6a48,48,0,0,0,43.5,27.7H208a32,32,0,0,0,32-32h0a32,32,0,0,0-32-32H188a28,28,0,0,1-28-28h0a28,28,0,0,1,28-28h20a24,24,0,0,1,24,24h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M108.67,168A40,40,0,0,1,72,192H64a48,48,0,0,1-48-48V120A48,48,0,0,1,64,72h9.43a48,48,0,0,1,43.5,27.7l30.14,64.6a48,48,0,0,0,43.5,27.7H208a32,32,0,0,0,32-32h0a32,32,0,0,0-32-32H188a28,28,0,0,1-28-28h0a28,28,0,0,1,28-28h20a24,24,0,0,1,24,24h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M108.67,168A40,40,0,0,1,72,192H64a48,48,0,0,1-48-48V120A48,48,0,0,1,64,72h9.43a48,48,0,0,1,43.5,27.7l30.14,64.6a48,48,0,0,0,43.5,27.7H208a32,32,0,0,0,32-32h0a32,32,0,0,0-32-32H188a28,28,0,0,1-28-28h0a28,28,0,0,1,28-28h20a24,24,0,0,1,24,24h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M216,40H40A16,16,0,0,0,24,56V200a16,16,0,0,0,16,16H216a16,16,0,0,0,16-16V56A16,16,0,0,0,216,40ZM184,184H172.61a40.09,40.09,0,0,1-36.42-23.45l-23-50.48A24,24,0,0,0,91.39,96H80a24,24,0,0,0-24,24v24a24,24,0,0,0,24,24h8a23.92,23.92,0,0,0,18.74-9,8,8,0,1,1,12.48,10A39.83,39.83,0,0,1,88,184H80a40,40,0,0,1-40-40V120A40,40,0,0,1,80,80H91.39a40.09,40.09,0,0,1,36.42,23.45l22.95,50.48A24,24,0,0,0,172.61,168H184a16,16,0,0,0,0-32h-8a28,28,0,0,1,0-56h12a20,20,0,0,1,20,20,8,8,0,0,1-16,0,4,4,0,0,0-4-4H176a12,12,0,0,0,0,24h8a32,32,0,0,1,0,64Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M208,192H64a48,48,0,0,1-48-48V120A48,48,0,0,1,64,72H208a24,24,0,0,1,24,24l8,64A32,32,0,0,1,208,192Z\" opacity=\"0.2\"></path><path d=\"M108.67,168A40,40,0,0,1,72,192H64a48,48,0,0,1-48-48V120A48,48,0,0,1,64,72h9.43a48,48,0,0,1,43.5,27.7l30.14,64.6a48,48,0,0,0,43.5,27.7H208a32,32,0,0,0,32-32h0a32,32,0,0,0-32-32H188a28,28,0,0,1-28-28h0a28,28,0,0,1,28-28h20a24,24,0,0,1,24,24h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M108.67,168A40,40,0,0,1,72,192H64a48,48,0,0,1-48-48V120A48,48,0,0,1,64,72h9.43a48,48,0,0,1,43.5,27.7l30.14,64.6a48,48,0,0,0,43.5,27.7H208a32,32,0,0,0,32-32h0a32,32,0,0,0-32-32H188a28,28,0,0,1-28-28h0a28,28,0,0,1,28-28h20a24,24,0,0,1,24,24h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
