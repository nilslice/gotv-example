package content

import (
	"fmt"

	"net/http"
	"net/url"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Video struct {
	item.Item

	Title       string `json:"title"`
	VideoUrl    string `json:"video_url"`
	VideoId     string `json:"video_id"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	Contributor string `json:"contributor"`
}

// MarshalEditor writes a buffer of html to edit a Video within the CMS
// and implements editor.Editable
func (v *Video) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(v,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Video field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", v, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("VideoUrl", v, map[string]string{
				"label":       "VideoUrl",
				"type":        "text",
				"placeholder": "Enter the VideoUrl here",
			}),
		},
		editor.Field{
			View: v.makeYoutubeThumbnail(),
		},
		editor.Field{
			View: editor.Input("VideoId", v, map[string]string{
				"type": "hidden",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", v, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Select("Rating", v, map[string]string{
				"label": "Rating",
			}, map[string]string{
				"G":     "G",
				"PG":    "PG",
				"PG-13": "PG-13",
				"R":     "R",
			}),
		},
		editor.Field{
			View: reference.Select("Contributor", v, map[string]string{
				"label": "Contributor",
			}, "Contributor", `{{.name}}`),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Video editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Video"] = func() interface{} { return new(Video) }
}

func (v *Video) String() string {
	return v.Title
}

func (v *Video) BeforeSave(res http.ResponseWriter, req *http.Request) error {
	vid := req.FormValue("video_url")

	u, _ := url.Parse(vid)
	req.PostForm.Set("video_id", u.Query().Get("v"))

	return nil
}

func (v *Video) makeYoutubeThumbnail() []byte {
	// https://img.youtube.com/vi/$VIDID/0.jpg
	if v.VideoId == "" {
		return nil
	}

	html := []byte(`
	<div class="input-field col s12">
		<img src="https://img.youtube.com/vi/` + v.VideoId + `/0.jpg" style="width:300px"/>
	</div>
	`)

	return html
}
