package clienthandlers

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	mw "github.com/tsawler/goblender/pkg/middleware"
	"net/http"
)

// ClientRoutes are the client specific routes
func ClientRoutes(mux *pat.PatternServeMux, standardMiddleWare, dynamicMiddleware alice.Chain) (*pat.PatternServeMux, error) {

	// we can override routes in goblender, if we wish, e.g.
	//mux.Get("/", dynamicMiddleware.ThenFunc(pageHandlers.Home))

	mux.Get("/faust/ft-vote/:ID", dynamicMiddleware.ThenFunc(DisplayFTVoteForm))
	mux.Post("/faust/ft-vote", dynamicMiddleware.ThenFunc(PostFT))

	mux.Get("/faust/pt-vote/:ID", dynamicMiddleware.ThenFunc(DisplayPTVoteForm))
	mux.Post("/faust/pt-vote", dynamicMiddleware.ThenFunc(PostPT))

	mux.Get("/admin/votes/send", dynamicMiddleware.Append(mw.Auth).Append(VotesRole).ThenFunc(SendInvitePage))
	mux.Get("/faust/invite/pt", dynamicMiddleware.Append(mw.Auth).Append(VotesRole).ThenFunc(SendInvitationsPT))
	mux.Get("/faust/invite/ft", dynamicMiddleware.Append(mw.Auth).Append(VotesRole).ThenFunc(SendInvitationsFT))
	mux.Get("/admin/votes/results", dynamicMiddleware.Append(mw.Auth).Append(VotesRole).ThenFunc(VoteResults))
	mux.Get("/admin/votes/resend", dynamicMiddleware.Append(mw.Auth).Append(VotesRole).ThenFunc(Resend))
	mux.Post("/admin/votes/resend", dynamicMiddleware.Append(mw.Auth).Append(VotesRole).ThenFunc(PostResend))

	// public folder
	fileServer := http.FileServer(http.Dir("./client/clienthandlers/public/"))
	mux.Get("/client/static/", http.StripPrefix("/client/static", fileServer))

	return mux, nil
}
