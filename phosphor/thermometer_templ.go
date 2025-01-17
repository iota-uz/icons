// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Thermometer(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M212,56a28,28,0,1,0,28,28A28,28,0,0,0,212,56Zm0,40a12,12,0,1,1,12-12A12,12,0,0,1,212,96Zm-60,50.08V40a32,32,0,0,0-64,0V146.08a56,56,0,1,0,64,0ZM136,104H104V40a16,16,0,0,1,32,0Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"212\" cy=\"84\" r=\"20\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M152,138V48a32,32,0,0,0-64,0v90a56,56,0,1,0,64,0Zm-32,70a24,24,0,1,1,24-24A24,24,0,0,1,120,208Z\" opacity=\"0.2\"></path><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"212\" cy=\"84\" r=\"20\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></circle><line x1=\"108\" y1=\"188\" x2=\"108\" y2=\"92\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M68,154.78V56a40,40,0,0,1,80,0v98.77h0a52,52,0,1,1-80,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><circle cx=\"108\" cy=\"188\" r=\"24\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"212\" cy=\"84\" r=\"20\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></circle><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"212\" cy=\"84\" r=\"20\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></circle><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"212\" cy=\"84\" r=\"20\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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

func ThermometerCold(props Props) templ.Component {
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
			if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M248.91,77.72l-20,6.49,12.34,17a8,8,0,1,1-12.94,9.4L216,93.61l-12.34,17a8,8,0,0,1-12.94-9.4l12.34-17-20-6.49A8,8,0,0,1,188,62.5L208,69V48a8,8,0,0,1,16,0V69l20-6.49a8,8,0,0,1,4.95,15.22ZM176,192a56,56,0,1,1-88-46V40a32,32,0,0,1,64,0V146A56.23,56.23,0,0,1,176,192Zm-95.18-8h78.36A40.16,40.16,0,0,0,140,157.35a8,8,0,0,1-4-6.93V40a16,16,0,0,0-32,0V150.42a8,8,0,0,1-4,6.93A40.16,40.16,0,0,0,80.82,184Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M152,138V48a32,32,0,0,0-64,0v90a56,56,0,1,0,64,0Zm-32,70a24,24,0,1,1,24-24A24,24,0,0,1,120,208Z\" opacity=\"0.2\"></path><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"216\" y1=\"48\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"185.57\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"197.19\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"234.81\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"246.43\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M68,154.78V56a40,40,0,0,1,80,0v98.77h0a52,52,0,1,1-80,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"108\" y1=\"188\" x2=\"108\" y2=\"132\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"212\" y1=\"48\" x2=\"212\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"181.57\" y1=\"70.11\" x2=\"212\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"193.19\" y1=\"105.89\" x2=\"212\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"230.81\" y1=\"105.89\" x2=\"212\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"242.43\" y1=\"70.11\" x2=\"212\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><circle cx=\"108\" cy=\"188\" r=\"24\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"216\" y1=\"48\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"185.57\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"197.19\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"234.81\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"246.43\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"216\" y1=\"48\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"185.57\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"197.19\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"234.81\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"246.43\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"120\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"216\" y1=\"48\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"185.57\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"197.19\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"234.81\" y1=\"105.89\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"246.43\" y1=\"70.11\" x2=\"216\" y2=\"80\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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

func ThermometerHot(props Props) templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var6 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M152,146.08V40a32,32,0,0,0-64,0V146.08a56,56,0,1,0,64,0ZM136,56H104V40a16,16,0,0,1,32,0Zm41.3,24.77a8,8,0,0,1,2.33-11.07c15-9.79,26.87-4.75,35.51-1.06C223,72,227.76,74,235.63,68.89a8,8,0,0,1,8.74,13.41C237.88,86.53,232,88,226.69,88c-7,0-12.92-2.54-17.83-4.63C201,80,196.24,78,188.37,83.11A8,8,0,0,1,177.3,80.77Zm69.4,22.46a8,8,0,0,1-2.33,11.07C237.88,118.53,232,120,226.69,120c-7,0-12.92-2.54-17.83-4.63-7.87-3.36-12.62-5.38-20.49-.25a8,8,0,0,1-8.74-13.41c15-9.79,26.87-4.75,35.51-1.06,7.87,3.36,12.62,5.39,20.49.25A8,8,0,0,1,246.7,103.23Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M152,138V48a32,32,0,0,0-64,0v90a56,56,0,1,0,64,0Zm-32,70a24,24,0,1,1,24-24A24,24,0,0,1,120,208Z\" opacity=\"0.2\"></path><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"48\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M184,76.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M184,108.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"108\" y1=\"188\" x2=\"108\" y2=\"56\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M68,154.78V56a40,40,0,0,1,80,0v98.77h0a52,52,0,1,1-80,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><circle cx=\"108\" cy=\"188\" r=\"24\"></circle><path d=\"M184,76.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M184,116.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"48\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M184,76.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M184,108.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"48\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M184,76.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M184,108.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"120\" y1=\"160\" x2=\"120\" y2=\"48\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"120\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M88,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M184,76.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M184,108.4c22.91-14.92,33.09,14.12,56-.8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			return nil
		})
		templ_7745c5c3_Err = svg(props).Render(templ.WithChildren(ctx, templ_7745c5c3_Var6), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func ThermometerSimple(props Props) templ.Component {
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
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M160,146.08V40a32,32,0,0,0-64,0V146.08a56,56,0,1,0,64,0ZM128,24a16,16,0,0,1,16,16v64H112V40A16,16,0,0,1,128,24Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M160,138V48a32,32,0,0,0-64,0v90a56,56,0,1,0,64,0Zm-32,70a24,24,0,1,1,24-24A24,24,0,0,1,128,208Z\" opacity=\"0.2\"></path><line x1=\"128\" y1=\"160\" x2=\"128\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"128\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M96,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"188\" x2=\"128\" y2=\"92\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M88,154.78V56a40,40,0,0,1,80,0v98.77h0a52,52,0,1,1-80,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><circle cx=\"128\" cy=\"188\" r=\"24\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"160\" x2=\"128\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><circle cx=\"128\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></circle><path d=\"M96,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"160\" x2=\"128\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><circle cx=\"128\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></circle><path d=\"M96,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><line x1=\"128\" y1=\"160\" x2=\"128\" y2=\"88\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"128\" cy=\"184\" r=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></circle><path d=\"M96,48a32,32,0,0,1,64,0v90a56,56,0,1,1-64,0Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			return nil
		})
		templ_7745c5c3_Err = svg(props).Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
