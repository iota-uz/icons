// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func SwimmingPool(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"52\" x2=\"168\" y2=\"52\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"88\" y1=\"88\" x2=\"168\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"88\" y1=\"124\" x2=\"168\" y2=\"124\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"88\" y1=\"32\" x2=\"88\" y2=\"137.39\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"168\" y1=\"32\" x2=\"168\" y2=\"139.29\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M32,164c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M32,208c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Thin {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"56\" x2=\"168\" y2=\"56\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"88\" y1=\"88\" x2=\"168\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"88\" y1=\"120\" x2=\"168\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"88\" y1=\"32\" x2=\"88\" y2=\"141.39\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"168\" y1=\"32\" x2=\"168\" y2=\"143.29\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M32,168c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M32,208c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"56\" x2=\"168\" y2=\"56\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"88\" y1=\"88\" x2=\"168\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"88\" y1=\"120\" x2=\"168\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"88\" y1=\"32\" x2=\"88\" y2=\"141.39\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"168\" y1=\"32\" x2=\"168\" y2=\"143.29\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M32,168c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M32,208c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M24,168a8,8,0,0,1,8-8c14.42,0,22.19,5.18,28.44,9.34C66,173.06,70.42,176,80,176s14-2.94,19.56-6.66c6.24-4.16,14-9.34,28.43-9.34s22.2,5.18,28.44,9.34c5.58,3.72,10,6.66,19.57,6.66s14-2.94,19.56-6.66c6.25-4.16,14-9.34,28.44-9.34a8,8,0,0,1,0,16c-9.58,0-14,2.94-19.56,6.66-6.25,4.16-14,9.34-28.44,9.34s-22.2-5.18-28.44-9.34C142,178.94,137.57,176,128,176s-14,2.94-19.56,6.66c-6.24,4.16-14,9.34-28.43,9.34s-22.19-5.18-28.44-9.34C46,178.94,41.58,176,32,176A8,8,0,0,1,24,168Zm200,32c-14.42,0-22.19,5.18-28.44,9.34C190,213.06,185.58,216,176,216s-14-2.94-19.57-6.66c-6.24-4.16-14-9.34-28.44-9.34s-22.19,5.18-28.43,9.34C94,213.06,89.57,216,80,216s-14-2.94-19.56-6.66C54.19,205.18,46.42,200,32,200a8,8,0,0,0,0,16c9.58,0,14,2.94,19.56,6.66,6.25,4.16,14,9.34,28.44,9.34s22.19-5.18,28.43-9.34c5.58-3.72,10-6.66,19.56-6.66s14,2.94,19.57,6.66c6.24,4.16,14,9.34,28.44,9.34s22.19-5.18,28.44-9.34c5.57-3.72,10-6.66,19.56-6.66a8,8,0,0,0,0-16ZM80,141.39V32a8,8,0,0,1,16,0v8h64V32a8,8,0,0,1,16,0V143.29a8,8,0,0,1-16,0V136H96v5.39a8,8,0,0,1-16,0ZM96,72a8,8,0,0,0,8,8h48a8,8,0,0,0,0-16H104A8,8,0,0,0,96,72Zm0,32a8,8,0,0,0,8,8h48a8,8,0,0,0,0-16H104A8,8,0,0,0,96,104Z\" fill=\"currentColor\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == DuoTone {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M32,208c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16V168c-24,0-24,16-48,16s-24-16-48-16-24,16-48,16-24-16-48-16Z\" opacity=\"0.2\" fill=\"currentColor\"></path><line x1=\"88\" y1=\"56\" x2=\"168\" y2=\"56\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"88\" y1=\"88\" x2=\"168\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"88\" y1=\"120\" x2=\"168\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"88\" y1=\"32\" x2=\"88\" y2=\"141.39\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"168\" y1=\"32\" x2=\"168\" y2=\"143.29\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M32,168c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M32,208c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"88\" y1=\"56\" x2=\"168\" y2=\"56\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"88\" y1=\"88\" x2=\"168\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"88\" y1=\"120\" x2=\"168\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"88\" y1=\"32\" x2=\"88\" y2=\"141.39\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"168\" y1=\"32\" x2=\"168\" y2=\"143.29\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M32,168c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M32,208c24,0,24,16,48,16s24-16,48-16,24,16,48,16,24-16,48-16\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
