// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Cheers(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M213.93,213.94l-17.65,4.73-10.42-38.89a40.06,40.06,0,0,0,20.77-46.14c-12.6-47-38.78-88.22-39.89-89.95a8,8,0,0,0-8.68-3.45L136.2,45.71c0-8.25-.18-13.43-.21-14.08a8,8,0,0,0-6.05-7.39l-32-8a8,8,0,0,0-8.68,3.45c-1.11,1.73-27.29,42.93-39.89,90a40.06,40.06,0,0,0,20.77,46.14L59.72,194.67l-17.65-4.73a8,8,0,0,0-4.14,15.46l48,12.86a8.23,8.23,0,0,0,2.07.27,8,8,0,0,0,2.07-15.73l-14.9-4,10.42-38.89c.81.05,1.61.08,2.41.08a40.12,40.12,0,0,0,37-24.88c1.18,6.37,2.6,12.82,4.31,19.22A40.08,40.08,0,0,0,168,184c.8,0,1.6,0,2.41-.08l10.42,38.89-14.9,4A8,8,0,0,0,168,242.53a8.23,8.23,0,0,0,2.07-.27l48-12.86a8,8,0,0,0-4.14-15.46ZM156.22,57.19c2.78,4.7,7.23,12.54,12.2,22.46L136,87.77c-.42-10-.38-18.25-.25-23.79,0-.56.05-1.12.08-1.68Zm-56.44-24,20.37,5.09c.06,4.28,0,10.67-.21,18.47-.06,1.21-.16,3.19-.23,5.84,0,1-.1,2-.16,3L86.69,57.43C92,46.67,96.84,38.16,99.78,33.19Zm85.06,10.39a8,8,0,0,1,3.58-10.74l16-8a8,8,0,1,1,7.16,14.32l-16,8a8,8,0,0,1-10.74-3.58ZM232,72a8,8,0,0,1-8,8H208a8,8,0,0,1,0-16h16A8,8,0,0,1,232,72ZM32.84,20.42a8,8,0,0,1,10.74-3.58l16,8a8,8,0,0,1-7.16,14.32l-16-8A8,8,0,0,1,32.84,20.42ZM40,72H24a8,8,0,0,1,0-16H40a8,8,0,0,1,0,16Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M75.23,62.81a307.67,307.67,0,0,0-18.13,48.9,32,32,0,1,0,61.82,16.56,332.07,332.07,0,0,0,8-52.54Z\" opacity=\"0.2\"></path><path d=\"M179.94,85a310.76,310.76,0,0,1,19,50.69,32,32,0,1,1-61.82,16.56c-5-18.72-7.46-38-8.59-54.39Z\" opacity=\"0.2\"></path><path d=\"M128,56l32-8s26.48,41.35,38.9,87.71a32,32,0,1,1-61.82,16.56C124.66,105.91,128,56,128,56Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"176.27\" y1=\"174.9\" x2=\"190.63\" y2=\"228.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"216\" y1=\"221.67\" x2=\"168\" y2=\"234.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M128,32,96,24S69.52,65.35,57.1,111.71a32,32,0,1,0,61.82,16.56C130.29,81.8,128,32,128,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"79.73\" y1=\"150.9\" x2=\"65.37\" y2=\"204.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"40\" y1=\"197.67\" x2=\"88\" y2=\"210.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M128.49,97.88,179.94,85\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M126.92,75.73,75.23,62.81\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"192\" y1=\"40\" x2=\"208\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"208\" y1=\"72\" x2=\"224\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"56\" y1=\"32\" x2=\"40\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"40\" y1=\"64\" x2=\"24\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M128,56l32-8s26.48,41.35,38.9,87.71a32,32,0,1,1-61.82,16.56C124.66,105.91,128,56,128,56Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"176.27\" y1=\"174.9\" x2=\"190.63\" y2=\"228.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"216\" y1=\"221.67\" x2=\"168\" y2=\"234.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M128,32,96,24S69.52,65.35,57.1,111.71a32,32,0,1,0,61.82,16.56C130.29,81.8,128,32,128,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"79.73\" y1=\"150.9\" x2=\"65.37\" y2=\"204.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"40\" y1=\"197.67\" x2=\"88\" y2=\"210.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><path d=\"M128.49,97.88,179.94,85\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M126.92,75.73,75.23,62.81\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><line x1=\"196\" y1=\"36\" x2=\"208\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"212\" y1=\"72\" x2=\"224\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"52\" y1=\"28\" x2=\"40\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line><line x1=\"36\" y1=\"64\" x2=\"24\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M128,56l32-8s26.48,41.35,38.9,87.71a32,32,0,1,1-61.82,16.56C124.66,105.91,128,56,128,56Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"176.27\" y1=\"174.9\" x2=\"190.63\" y2=\"228.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"216\" y1=\"221.67\" x2=\"168\" y2=\"234.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M128,32,96,24S69.52,65.35,57.1,111.71a32,32,0,1,0,61.82,16.56C130.29,81.8,128,32,128,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"79.73\" y1=\"150.9\" x2=\"65.37\" y2=\"204.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"40\" y1=\"197.67\" x2=\"88\" y2=\"210.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><path d=\"M128.49,97.88,179.94,85\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M126.92,75.73,75.23,62.81\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><line x1=\"192\" y1=\"40\" x2=\"208\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"208\" y1=\"72\" x2=\"224\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"56\" y1=\"32\" x2=\"40\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"40\" y1=\"64\" x2=\"24\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M128,56l32-8s26.48,41.35,38.9,87.71a32,32,0,1,1-61.82,16.56C124.66,105.91,128,56,128,56Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"176.27\" y1=\"174.9\" x2=\"190.63\" y2=\"228.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"216\" y1=\"221.67\" x2=\"168\" y2=\"234.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M128,32,96,24S69.52,65.35,57.1,111.71a32,32,0,1,0,61.82,16.56C130.29,81.8,128,32,128,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"79.73\" y1=\"150.9\" x2=\"65.37\" y2=\"204.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"40\" y1=\"197.67\" x2=\"88\" y2=\"210.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><path d=\"M128.49,97.88,179.94,85\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M126.92,75.73,75.23,62.81\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><line x1=\"192\" y1=\"40\" x2=\"208\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"208\" y1=\"72\" x2=\"224\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"56\" y1=\"32\" x2=\"40\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"40\" y1=\"64\" x2=\"24\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M128,56l32-8s26.48,41.35,38.9,87.71a32,32,0,1,1-61.82,16.56C124.66,105.91,128,56,128,56Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"176.27\" y1=\"174.9\" x2=\"190.63\" y2=\"228.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"216\" y1=\"221.67\" x2=\"168\" y2=\"234.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M128,32,96,24S69.52,65.35,57.1,111.71a32,32,0,1,0,61.82,16.56C130.29,81.8,128,32,128,32Z\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"79.73\" y1=\"150.9\" x2=\"65.37\" y2=\"204.47\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"40\" y1=\"197.67\" x2=\"88\" y2=\"210.53\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><path d=\"M128.49,97.88,179.94,85\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M126.92,75.73,75.23,62.81\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><line x1=\"192\" y1=\"40\" x2=\"208\" y2=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"208\" y1=\"72\" x2=\"224\" y2=\"72\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"56\" y1=\"32\" x2=\"40\" y2=\"24\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"40\" y1=\"64\" x2=\"24\" y2=\"64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line>")
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
