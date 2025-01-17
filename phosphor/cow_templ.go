// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Cow(props Props) templ.Component {
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
			if props.Variant == Filled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M104,192a8,8,0,0,1-8,8H80a8,8,0,0,1,0-16H96A8,8,0,0,1,104,192Zm72-8H160a8,8,0,0,0,0,16h16a8,8,0,0,0,0-16Zm68.39-61.88A16,16,0,0,1,232,128H200v32a40,40,0,0,1-24,72H80a40,40,0,0,1-24-72V128H24A16,16,0,0,1,8.31,109,56.13,56.13,0,0,1,63.22,64h1.64A55.83,55.83,0,0,1,48,24a8,8,0,0,1,16,0,40,40,0,0,0,40,40h48a40,40,0,0,0,40-40,8,8,0,0,1,16,0,55.83,55.83,0,0,1-16.86,40h1.64a56.13,56.13,0,0,1,54.91,45A15.82,15.82,0,0,1,244.39,122.12ZM144,124a12,12,0,1,0,12-12A12,12,0,0,0,144,124Zm-56,0a12,12,0,1,0,12-12A12,12,0,0,0,88,124ZM56,112v-8a39.81,39.81,0,0,1,8-24h-.8A40.09,40.09,0,0,0,24,112Zm144,80a24,24,0,0,0-24-24H80a24,24,0,0,0,0,48h96A24,24,0,0,0,200,192Zm32-80a40.08,40.08,0,0,0-39.2-32H192a39.81,39.81,0,0,1,8,24v8Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><rect x=\"48\" y=\"160\" width=\"160\" height=\"64\" rx=\"32\" opacity=\"0.2\"></rect><path d=\"M64,104A32,32,0,0,1,96,72H63.22a48,48,0,0,0-47.07,38.53A8,8,0,0,0,24,120H64Z\" opacity=\"0.2\"></path><path d=\"M239.85,110.53A48,48,0,0,0,192.78,72H160a32,32,0,0,1,32,32v16h40A8,8,0,0,0,239.85,110.53Z\" opacity=\"0.2\"></path><path d=\"M56,24h0a48,48,0,0,0,48,48h48a48,48,0,0,0,48-48h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><rect x=\"48\" y=\"160\" width=\"160\" height=\"64\" rx=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"80\" y1=\"192\" x2=\"96\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"160\" y1=\"192\" x2=\"176\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"100\" cy=\"124\" r=\"12\"></circle><circle cx=\"156\" cy=\"124\" r=\"12\"></circle><path d=\"M160,72h32.78a48,48,0,0,1,47.07,38.53A8,8,0,0,1,232,120H192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M96,72H63.22a48,48,0,0,0-47.07,38.53A8,8,0,0,0,24,120H64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M64,164.28V104A32,32,0,0,1,96,72h64a32,32,0,0,1,32,32v60.28\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Bold" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M56,24h0a48,48,0,0,0,48,48h48a48,48,0,0,0,48-48h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><rect x=\"48\" y=\"172\" width=\"160\" height=\"52\" rx=\"26\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></rect><circle cx=\"104\" cy=\"128\" r=\"16\"></circle><circle cx=\"152\" cy=\"128\" r=\"16\"></circle><path d=\"M160,72h32.78a48,48,0,0,1,47.07,38.53A8,8,0,0,1,232,120H192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M96,72H63.22a48,48,0,0,0-47.07,38.53A8,8,0,0,0,24,120H64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path><path d=\"M64,174V104A32,32,0,0,1,96,72h64a32,32,0,0,1,32,32v70\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"24\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M56,24h0a48,48,0,0,0,48,48h48a48,48,0,0,0,48-48h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><rect x=\"48\" y=\"160\" width=\"160\" height=\"64\" rx=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></rect><line x1=\"80\" y1=\"192\" x2=\"96\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><line x1=\"160\" y1=\"192\" x2=\"176\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></line><circle cx=\"100\" cy=\"124\" r=\"8\"></circle><circle cx=\"156\" cy=\"124\" r=\"8\"></circle><path d=\"M160,72h32.78a48,48,0,0,1,47.07,38.53A8,8,0,0,1,232,120H192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M96,72H63.22a48,48,0,0,0-47.07,38.53A8,8,0,0,0,24,120H64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path><path d=\"M64,164.28V104A32,32,0,0,1,96,72h64a32,32,0,0,1,32,32v60.28\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"8\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M56,24h0a48,48,0,0,0,48,48h48a48,48,0,0,0,48-48h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><rect x=\"48\" y=\"160\" width=\"160\" height=\"64\" rx=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></rect><line x1=\"80\" y1=\"192\" x2=\"96\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><line x1=\"160\" y1=\"192\" x2=\"176\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></line><circle cx=\"100\" cy=\"124\" r=\"10\"></circle><circle cx=\"156\" cy=\"124\" r=\"10\"></circle><path d=\"M160,72h32.78a48,48,0,0,1,47.07,38.53A8,8,0,0,1,232,120H192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M96,72H63.22a48,48,0,0,0-47.07,38.53A8,8,0,0,0,24,120H64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path><path d=\"M64,164.28V104A32,32,0,0,1,96,72h64a32,32,0,0,1,32,32v60.28\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"12\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M56,24h0a48,48,0,0,0,48,48h48a48,48,0,0,0,48-48h0\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><rect x=\"48\" y=\"160\" width=\"160\" height=\"64\" rx=\"32\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></rect><line x1=\"80\" y1=\"192\" x2=\"96\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><line x1=\"160\" y1=\"192\" x2=\"176\" y2=\"192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></line><circle cx=\"100\" cy=\"124\" r=\"12\"></circle><circle cx=\"156\" cy=\"124\" r=\"12\"></circle><path d=\"M160,72h32.78a48,48,0,0,1,47.07,38.53A8,8,0,0,1,232,120H192\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M96,72H63.22a48,48,0,0,0-47.07,38.53A8,8,0,0,0,24,120H64\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path><path d=\"M64,164.28V104A32,32,0,0,1,96,72h64a32,32,0,0,1,32,32v60.28\" fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"16\"></path>")
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
