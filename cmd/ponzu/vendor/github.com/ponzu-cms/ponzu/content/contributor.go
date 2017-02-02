package content

import (
	"fmt"

	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Contributor struct {
	item.Item

	Name    string `json:"name"`
	Picture string `json:"picture"`
	Bio     string `json:"bio"`
}

// MarshalEditor writes a buffer of html to edit a Contributor within the CMS
// and implements editor.Editable
func (c *Contributor) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(c,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Contributor field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", c, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.File("Picture", c, map[string]string{
				"label":       "Picture",
				"placeholder": "Enter the Picture here",
			}),
		},
		editor.Field{
			View: editor.Textarea("Bio", c, map[string]string{
				"label":       "Bio",
				"placeholder": "Enter the Bio here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Contributor editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Contributor"] = func() interface{} { return new(Contributor) }
}

func (c *Contributor) String() string {
	return c.Name
}

func (c *Contributor) Accept(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (c *Contributor) Approve(res http.ResponseWriter, req *http.Request) error {
	return nil
}
