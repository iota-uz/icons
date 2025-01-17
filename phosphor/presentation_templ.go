// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func PresentationChart(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"96\" y1=\"120\" x2=\"96\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"128\" y1=\"104\" x2=\"128\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"160\" y1=\"88\" x2=\"160\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M216,40H136V24a8,8,0,0,0-16,0V40H40A16,16,0,0,0,24,56V176a16,16,0,0,0,16,16H79.36L57.75,219a8,8,0,0,0,12.5,10l29.59-37h56.32l29.59,37a8,8,0,1,0,12.5-10l-21.61-27H216a16,16,0,0,0,16-16V56A16,16,0,0,0,216,40ZM104,144a8,8,0,0,1-16,0V120a8,8,0,0,1,16,0Zm32,0a8,8,0,0,1-16,0V104a8,8,0,0,1,16,0Zm32,0a8,8,0,0,1-16,0V88a8,8,0,0,1,16,0Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" opacity=\"0.2\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"120\" x2=\"96\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"104\" x2=\"128\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"160\" y1=\"88\" x2=\"160\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"92\" y1=\"120\" x2=\"92\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"128\" y1=\"104\" x2=\"128\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"164\" y1=\"88\" x2=\"164\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"96\" y1=\"120\" x2=\"96\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"128\" y1=\"104\" x2=\"128\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"160\" y1=\"88\" x2=\"160\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"120\" x2=\"96\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"104\" x2=\"128\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"160\" y1=\"88\" x2=\"160\" y2=\"144\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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

func Presentation(props Props) templ.Component {
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
			if props.Variant == Light {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M216,40H136V24a8,8,0,0,0-16,0V40H40A16,16,0,0,0,24,56V176a16,16,0,0,0,16,16H79.36L57.75,219a8,8,0,0,0,12.5,10l29.59-37h56.32l29.59,37a8,8,0,1,0,12.5-10l-21.61-27H216a16,16,0,0,0,16-16V56A16,16,0,0,0,216,40Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" opacity=\"0.2\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"32\" y=\"48\" width=\"192\" height=\"136\" rx=\"8\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"160\" y1=\"184\" x2=\"192\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"96\" y1=\"184\" x2=\"64\" y2=\"224\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"128\" y1=\"48\" x2=\"128\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
