package clienthandlers

import (
	"fmt"
	channel_data "github.com/tsawler/goblender/pkg/channel-data"
	"github.com/tsawler/goblender/pkg/helpers"
	"github.com/tsawler/goblender/pkg/templates"
	"github.com/tsawler/goblender/pkg/urlsigner"
	"html/template"
	"net/http"
	"strconv"
)

// JSONResponse is a generic struct to hold json responses
type JSONResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// DisplayFTVoteForm displays ft form
func DisplayFTVoteForm(w http.ResponseWriter, r *http.Request) {
	// validate the link
	url := r.RequestURI
	testURL := fmt.Sprintf("%s%s", app.ServerURL, url)
	urlsigner.NewURLSigner(app.URLSignerKey)
	okay := urlsigner.VerifyToken(testURL)

	if !okay {
		session.Put(r.Context(), "error", "Invalid URL!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// make sure they have not voted already
	id, _ := strconv.Atoi(r.URL.Query().Get(":ID"))
	member, err := dbModel.GetFTMember(id)
	if err != nil {
		session.Put(r.Context(), "error", "Invalid URL!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if member.Voted == 1 {
		session.Put(r.Context(), "error", "You have already voted")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	intMap := make(map[string]int)
	intMap["id"] = id

	helpers.Render(w, r, "ft-vote.page.tmpl", &templates.TemplateData{
		IntMap: intMap,
	})
}

// DisplayPTVoteForm displays pt form
func DisplayPTVoteForm(w http.ResponseWriter, r *http.Request) {
	session.Put(r.Context(), "lang", "en")

	// validate the link
	url := r.RequestURI
	testURL := fmt.Sprintf("%s%s", app.ServerURL, url)
	urlsigner.NewURLSigner(app.URLSignerKey)
	okay := urlsigner.VerifyToken(testURL)

	if !okay {
		infoLog.Println("checking url failed")
		session.Put(r.Context(), "error", "Invalid URL!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// make sure they have not voted already
	id, _ := strconv.Atoi(r.URL.Query().Get(":ID"))
	member, err := dbModel.GetPTMember(id)
	if err != nil {
		// invalid url signature, so just throw a generic error page at the user
		session.Put(r.Context(), "error", "Invalid URL!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if member.Voted == 1 {
		session.Put(r.Context(), "error", "You have already voted")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	intMap := make(map[string]int)
	intMap["id"] = id

	helpers.Render(w, r, "pt-vote.page.tmpl", &templates.TemplateData{
		IntMap: intMap,
	})
}

// PostPT handle voting of FT member
func PostPT(w http.ResponseWriter, r *http.Request) {

}

// PostFT handle voting of PT member
func PostFT(w http.ResponseWriter, r *http.Request) {

}

// SendInvitations sends the invitations
func SendInvitationsPT(w http.ResponseWriter, r *http.Request) {
	pt, err := dbModel.GetAllPTMembers()
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	serverURL := app.ServerURL

	for _, x := range pt {
		linkEn := fmt.Sprintf("%s/faust/pt-vote/%d", serverURL, x.ID)
		urlsigner.NewURLSigner(app.URLSignerKey)
		signedLinkEn := urlsigner.GenerateTokenFromString(linkEn)

		html := fmt.Sprintf(`
<p>Dear %s:</p>

<p>Please use the link below to cast your vote to ratify the collective agreement for your unit. Note that voting is anonymous, and that the link below will only work once.</p>

<p>You have until Thursday, June 4th at 5:00PM to cast your vote.</p>

<p>Thank you.</p>

<p><a class="btn btn-primary" href="%s">Click here to cast your vote</a"></p>
`, x.FirstName, signedLinkEn)

		mailMessage := channel_data.MailData{
			ToName:      "",
			ToAddress:   x.Email,
			FromName:    "FAUST",
			FromAddress: "faust@stu.ca",
			Subject:     "Online vote to ratify agreement",
			Content:     template.HTML(html),
			Template:    "generic-email.mail.tmpl",
			CC:          nil,
			UseHermes:   false,
			Attachments: nil,
			StringMap:   nil,
			IntMap:      nil,
			FloatMap:    nil,
			RowSets:     nil,
		}

		helpers.SendEmail(mailMessage)
	}

	w.Write([]byte("Sent"))
}

func SendInvitationsFT(w http.ResponseWriter, r *http.Request) {
	pt, err := dbModel.GetAllFTMembers()
	if err != nil {
		errorLog.Println(err)
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	serverURL := app.ServerURL

	for _, x := range pt {
		linkEn := fmt.Sprintf("%s/faust/ft-vote/%d", serverURL, x.ID)
		urlsigner.NewURLSigner(app.URLSignerKey)
		signedLinkEn := urlsigner.GenerateTokenFromString(linkEn)

		html := fmt.Sprintf(`
<p>Dear %s:</p>

<p>Please use the link below to cast your vote to ratify the collective agreement for your unit. Note that voting is anonymous, and that the link below will only work once.</p>

<p>You have until Thursday, June 4th at 5:00PM to cast your vote.</p>

<p>Thank you.</p>

<p><a class="btn btn-primary" href="%s">Click here to cast your vote</a"></p>
`, x.FirstName, signedLinkEn)

		mailMessage := channel_data.MailData{
			ToName:      "",
			ToAddress:   x.Email,
			FromName:    "FAUST",
			FromAddress: "faust@stu.ca",
			Subject:     "Online vote to ratify agreement",
			Content:     template.HTML(html),
			Template:    "generic-email.mail.tmpl",
			CC:          nil,
			UseHermes:   false,
			Attachments: nil,
			StringMap:   nil,
			IntMap:      nil,
			FloatMap:    nil,
			RowSets:     nil,
		}

		helpers.SendEmail(mailMessage)
	}

	w.Write([]byte("Sent"))
}
