// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Eyedropper(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M224,67.3a35.79,35.79,0,0,0-11.26-25.66c-14-13.28-36.72-12.78-50.62,1.13L138.8,66.2a24,24,0,0,0-33.14.77l-5,5a16,16,0,0,0,0,22.64l2,2.06-51,51a39.75,39.75,0,0,0-10.53,38l-8,18.41A13.68,13.68,0,0,0,36,219.3a15.92,15.92,0,0,0,17.71,3.35L71.23,215a39.89,39.89,0,0,0,37.06-10.75l51-51,2.06,2.06a16,16,0,0,0,22.62,0l5-5a24,24,0,0,0,.74-33.18l23.75-23.87A35.75,35.75,0,0,0,224,67.3ZM97,193a24,24,0,0,1-24,6,8,8,0,0,0-5.55.31l-18.1,7.91L57,189.41a8,8,0,0,0,.25-5.75A23.88,23.88,0,0,1,63,159l51-51,33.94,34Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" opacity=\"0.2\"></path><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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

func EyedropperSample(props Props) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"52.28\" y1=\"160\" x2=\"141.25\" y2=\"160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"52.28\" y1=\"160\" x2=\"141.25\" y2=\"160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M224,67.3a35.79,35.79,0,0,0-11.26-25.66c-14-13.28-36.72-12.78-50.62,1.13L138.8,66.2a24,24,0,0,0-33.14.77l-5,5a16,16,0,0,0,0,22.64l2,2.06-51,51a39.75,39.75,0,0,0-10.53,38l-8,18.41A13.65,13.65,0,0,0,36,219.29a15.9,15.9,0,0,0,17.71,3.36L71.24,215a39.9,39.9,0,0,0,37.05-10.75l51-51,2.06,2.06a16,16,0,0,0,22.62,0l5-5a24,24,0,0,0,.74-33.18l23.75-23.87A35.75,35.75,0,0,0,224,67.3ZM138,152H70.07l44-44,33.94,34Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" opacity=\"0.2\"></path><path d=\"M49.67,186.24h0L40.48,207.3a5.74,5.74,0,0,0,1.2,6.36,8,8,0,0,0,8.86,1.67l20-8.74h0a32,32,0,0,0,32.06-8L141.25,160H52.3A32.07,32.07,0,0,0,49.67,186.24Z\" opacity=\"0.2\"></path><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"52.28\" y1=\"160\" x2=\"141.25\" y2=\"160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"52.28\" y1=\"160\" x2=\"141.25\" y2=\"160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M182.43,113.17l4.88,4.89a16,16,0,0,1,0,22.63l-9,9a8,8,0,0,1-11.31,0L106.34,89a8,8,0,0,1,0-11.31l9-9a16,16,0,0,1,22.63,0l4.89,4.88,25-25.1c10.79-10.79,28.37-11.45,39.44-1a28,28,0,0,1,.57,40.15Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M159.31,141.94l-56.68,56.69a32,32,0,0,1-32.06,8h0l-20,8.74a8,8,0,0,1-8.86-1.67h0a5.74,5.74,0,0,1-1.2-6.36l9.19-21.06h0a32,32,0,0,1,7.7-32.87l56.69-56.68\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"52.28\" y1=\"160\" x2=\"141.25\" y2=\"160\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			return nil
		})
		templ_7745c5c3_Err = svg(props).Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
