// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Brain(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M88,136a40,40,0,1,1-40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M168,136a40,40,0,1,0,40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M72,172H64A48,48,0,0,1,48,78.73V72a40,40,0,0,1,80,0V176\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M184,172h8a48,48,0,0,0,16-93.27V72a40,40,0,0,0-80,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M200,112h-4a28,28,0,0,1-28-28V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M56,112h4A28,28,0,0,0,88,84V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M212,76V72a44,44,0,0,0-74.86-31.31,3.93,3.93,0,0,0-1.14,2.8v88.72a4,4,0,0,0,6.2,3.33A47.67,47.67,0,0,1,167.68,128a8.18,8.18,0,0,1,8.31,7.58,8,8,0,0,1-8,8.42,32,32,0,0,0-32,32v33.88a4,4,0,0,0,1.49,3.12,47.92,47.92,0,0,0,74.21-17.16,4,4,0,0,0-4.49-5.56A68.06,68.06,0,0,1,192,192h-7.73a8.18,8.18,0,0,1-8.25-7.47,8,8,0,0,1,8-8.53h8a51.6,51.6,0,0,0,24-5.88v0A52,52,0,0,0,212,76Zm-12,36h-4a36,36,0,0,1-36-36V72a8,8,0,0,1,16,0v4a20,20,0,0,0,20,20h4a8,8,0,0,1,0,16ZM88,28A44.05,44.05,0,0,0,44,72v4a52,52,0,0,0-4,94.12h0A51.6,51.6,0,0,0,64,176h7.73A8.18,8.18,0,0,1,80,183.47,8,8,0,0,1,72,192H64a67.48,67.48,0,0,1-15.21-1.73,4,4,0,0,0-4.5,5.55A47.93,47.93,0,0,0,118.51,213a4,4,0,0,0,1.49-3.12V176a32,32,0,0,0-32-32,8,8,0,0,1-8-8.42A8.18,8.18,0,0,1,88.32,128a47.67,47.67,0,0,1,25.48,7.54,4,4,0,0,0,6.2-3.33V43.49a4,4,0,0,0-1.14-2.81A43.85,43.85,0,0,0,88,28Zm8,48a36,36,0,0,1-36,36H56a8,8,0,0,1,0-16h4A20,20,0,0,0,80,76V72a8,8,0,0,1,16,0Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M208,78.73V72a40,40,0,0,0-80,0,40,40,0,0,0-80,0v6.73a48,48,0,0,0,0,90.54h0V176a40,40,0,0,0,80,0,40,40,0,0,0,80,0v-6.73h0a48,48,0,0,0,0-90.54Z\" opacity=\"0.2\" fill=\"currentColor\"></path><path d=\"M88,136a40,40,0,1,1-40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M168,136a40,40,0,1,0,40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M72,172H64A48,48,0,0,1,48,78.73V72a40,40,0,0,1,80,0V176\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M184,172h8a48,48,0,0,0,16-93.27V72a40,40,0,0,0-80,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M200,112h-4a28,28,0,0,1-28-28V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M56,112h4A28,28,0,0,0,88,84V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Bold {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M72,172H64A48,48,0,0,1,48,78.73V72a40,40,0,0,1,80,0V176\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M184,172h8a48,48,0,0,0,16-93.27V72a40,40,0,0,0-80,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M196,116a28,28,0,0,1-28-28V84\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M60,116A28,28,0,0,0,88,88V84\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M152,139.33A40,40,0,1,0,208,176v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M104,139.33A40,40,0,1,1,48,176v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Thin {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M88,136a40,40,0,1,1-40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M168,136a40,40,0,1,0,40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M72,172H64A48,48,0,0,1,48,78.73V72a40,40,0,0,1,80,0V176\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M184,172h8a48,48,0,0,0,16-93.27V72a40,40,0,0,0-80,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M200,112h-4a28,28,0,0,1-28-28V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M56,112h4A28,28,0,0,0,88,84V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M88,136a40,40,0,1,1-40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M168,136a40,40,0,1,0,40,40v-6.73\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M72,172H64A48,48,0,0,1,48,78.73V72a40,40,0,0,1,80,0V176\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M184,172h8a48,48,0,0,0,16-93.27V72a40,40,0,0,0-80,0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M200,112h-4a28,28,0,0,1-28-28V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M56,112h4A28,28,0,0,0,88,84V80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
