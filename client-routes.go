package clienthandlers

import (
	"github.com/go-chi/chi/v5"
	mw "github.com/tsawler/goblender/pkg/middleware"
	"net/http"
)

// ClientRoutes are the client specific routes
func ClientRoutes(mux *chi.Mux) {
	fileServer := http.FileServer(http.Dir("./client/clienthandlers/public/"))
	mux.Handle("/client/static/*", http.StripPrefix("/client/static", fileServer))
	fileServer = http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.Get("/faust/ft-vote/{id}", DisplayFTVoteForm)
	mux.Post("/faust/ft-vote", PostFT)

	mux.Get("/faust/pt-vote/{id}", DisplayPTVoteForm)
	mux.Post("/faust/pt-vote", PostPT)

	mux.With(mw.Auth, VotesRole).Get("/admin/votes/send", SendInvitePage)
	mux.With(mw.Auth, VotesRole).Get("/faust/invite/pt", SendInvitationsPT)
	mux.With(mw.Auth, VotesRole).Get("/faust/invite/ft", SendInvitationsFT)
	mux.With(mw.Auth, VotesRole).Get("/admin/votes/results", VoteResults)
	mux.With(mw.Auth, VotesRole).Get("/admin/votes/resend", Resend)
	mux.With(mw.Auth, VotesRole).Post("/admin/votes/resend", PostResend)
}
