// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package phosphor

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Numpad(props Props) templ.Component {
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
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"64\" cy=\"48\" r=\"20\"></circle><circle cx=\"128\" cy=\"48\" r=\"20\"></circle><circle cx=\"192\" cy=\"48\" r=\"20\"></circle><circle cx=\"64\" cy=\"104\" r=\"20\"></circle><circle cx=\"128\" cy=\"104\" r=\"20\"></circle><circle cx=\"192\" cy=\"104\" r=\"20\"></circle><circle cx=\"64\" cy=\"160\" r=\"20\"></circle><circle cx=\"128\" cy=\"160\" r=\"20\"></circle><circle cx=\"128\" cy=\"216\" r=\"20\"></circle><circle cx=\"192\" cy=\"160\" r=\"20\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Thin" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"64\" cy=\"48\" r=\"12\"></circle><circle cx=\"128\" cy=\"48\" r=\"12\"></circle><circle cx=\"192\" cy=\"48\" r=\"12\"></circle><circle cx=\"64\" cy=\"104\" r=\"12\"></circle><circle cx=\"128\" cy=\"104\" r=\"12\"></circle><circle cx=\"192\" cy=\"104\" r=\"12\"></circle><circle cx=\"64\" cy=\"160\" r=\"12\"></circle><circle cx=\"128\" cy=\"160\" r=\"12\"></circle><circle cx=\"128\" cy=\"216\" r=\"12\"></circle><circle cx=\"192\" cy=\"160\" r=\"12\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Light" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"64\" cy=\"48\" r=\"14\"></circle><circle cx=\"128\" cy=\"48\" r=\"14\"></circle><circle cx=\"192\" cy=\"48\" r=\"14\"></circle><circle cx=\"64\" cy=\"104\" r=\"14\"></circle><circle cx=\"128\" cy=\"104\" r=\"14\"></circle><circle cx=\"192\" cy=\"104\" r=\"14\"></circle><circle cx=\"64\" cy=\"160\" r=\"14\"></circle><circle cx=\"128\" cy=\"160\" r=\"14\"></circle><circle cx=\"128\" cy=\"216\" r=\"14\"></circle><circle cx=\"192\" cy=\"160\" r=\"14\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "Filled" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M200,24H56A16,16,0,0,0,40,40V216a16,16,0,0,0,16,16H200a16,16,0,0,0,16-16V40A16,16,0,0,0,200,24ZM80,164a12,12,0,1,1,12-12A12,12,0,0,1,80,164Zm0-40a12,12,0,1,1,12-12A12,12,0,0,1,80,124Zm0-40A12,12,0,1,1,92,72,12,12,0,0,1,80,84Zm48,120a12,12,0,1,1,12-12A12,12,0,0,1,128,204Zm0-40a12,12,0,1,1,12-12A12,12,0,0,1,128,164Zm0-40a12,12,0,1,1,12-12A12,12,0,0,1,128,124Zm0-40a12,12,0,1,1,12-12A12,12,0,0,1,128,84Zm48,80a12,12,0,1,1,12-12A12,12,0,0,1,176,164Zm0-40a12,12,0,1,1,12-12A12,12,0,0,1,176,124Zm0-40a12,12,0,1,1,12-12A12,12,0,0,1,176,84Z\"></path>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if props.Variant == "DuoTone" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><path d=\"M64,48H192a0,0,0,0,1,0,0V200a16,16,0,0,1-16,16H80a16,16,0,0,1-16-16V48A0,0,0,0,1,64,48Z\" opacity=\"0.2\"></path><circle cx=\"64\" cy=\"48\" r=\"16\"></circle><circle cx=\"128\" cy=\"48\" r=\"16\"></circle><circle cx=\"192\" cy=\"48\" r=\"16\"></circle><circle cx=\"64\" cy=\"104\" r=\"16\"></circle><circle cx=\"128\" cy=\"104\" r=\"16\"></circle><circle cx=\"192\" cy=\"104\" r=\"16\"></circle><circle cx=\"64\" cy=\"160\" r=\"16\"></circle><circle cx=\"128\" cy=\"160\" r=\"16\"></circle><circle cx=\"128\" cy=\"216\" r=\"16\"></circle><circle cx=\"192\" cy=\"160\" r=\"16\"></circle>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<rect width=\"256\" height=\"256\" fill=\"none\"></rect><circle cx=\"64\" cy=\"48\" r=\"16\"></circle><circle cx=\"128\" cy=\"48\" r=\"16\"></circle><circle cx=\"192\" cy=\"48\" r=\"16\"></circle><circle cx=\"64\" cy=\"104\" r=\"16\"></circle><circle cx=\"128\" cy=\"104\" r=\"16\"></circle><circle cx=\"192\" cy=\"104\" r=\"16\"></circle><circle cx=\"64\" cy=\"160\" r=\"16\"></circle><circle cx=\"128\" cy=\"160\" r=\"16\"></circle><circle cx=\"128\" cy=\"216\" r=\"16\"></circle><circle cx=\"192\" cy=\"160\" r=\"16\"></circle>")
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
