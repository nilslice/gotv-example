package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Channel struct {
	item.Item

	Title  string   `json:"title"`
	Videos []string `json:"videos"`
	Rating string   `json:"rating"`
	Tags   []string `json:"tags"`
}

// MarshalEditor writes a buffer of html to edit a Channel within the CMS
// and implements editor.Editable
func (c *Channel) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(c,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Channel field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", c, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Videos", c, map[string]string{
				"label": "Videos",
			}, "Video", `{{.title}} ({{.rating}})`),
		},
		editor.Field{
			View: editor.Select("Rating", c, map[string]string{
				"label": "Rating",
			}, map[string]string{
				"G":     "G",
				"PG":    "PG",
				"PG-13": "PG-13",
				"R":     "R",
			}),
		},
		editor.Field{
			View: editor.Tags("Tags", c, map[string]string{
				"label": "Tags",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Channel editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Channel"] = func() interface{} { return new(Channel) }
}

func (c *Channel) String() string {
	return c.Title
}

func (c *Channel) Push() []string {
	return []string{
		"videos",
	}
}
