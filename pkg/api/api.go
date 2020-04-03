package api

import (
	"github.com/go-macaron/binding"
	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/api/middleware"
	"github.com/hooone/evening/pkg/api/routing"
)

func (hs *HTTPServer) registerRoutes() {
	reqSignedIn := middleware.ReqSignedIn
	bind := binding.Bind

	r := hs.RouteRegister
	r.Get("/logout", hs.Logout)
	r.Post("/login", bind(dtos.LoginCommand{}), Wrap(hs.LoginPost))
	r.Get("/login", hs.LoginView)

	r.Group("/api/navigation", func(navRoute routing.RouteRegister) {
		navRoute.Get("/", bind(dtos.LocaleForm{}), Wrap(hs.GetNavigation))
		navRoute.Post("/", bind(dtos.LocaleForm{}), Wrap(hs.GetNavigation))
		navRoute.Post("/createFolder", bind(dtos.CreateFolderForm{}), bind(dtos.LocaleForm{}), Wrap(hs.CreateFolder))
		navRoute.Post("/updateFolder", bind(dtos.UpdateFolderForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateFolder))
		navRoute.Post("/deleteFolder", bind(dtos.DeleteFolderForm{}), bind(dtos.LocaleForm{}), Wrap(hs.DeleteFolder))
		navRoute.Post("/createTreePage", bind(dtos.CreateTreePageForm{}), bind(dtos.LocaleForm{}), Wrap(hs.CreateTreePage))
		navRoute.Post("/createNodePage", bind(dtos.CreateNodePageForm{}), bind(dtos.LocaleForm{}), Wrap(hs.CreateNodePage))
		navRoute.Post("/updatePage", bind(dtos.UpdatePageForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdatePage))
		navRoute.Post("/deletePage", bind(dtos.DeletePageForm{}), bind(dtos.LocaleForm{}), Wrap(hs.DeletePage))
		navRoute.Post("/move", bind(dtos.NavMoveForm{}), bind(dtos.LocaleForm{}), Wrap(hs.NavMove))
	}, reqSignedIn)

	r.Group("/api/card", func(cardRoute routing.RouteRegister) {
		cardRoute.Get("/", bind(dtos.GetCardsForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetCards))
		cardRoute.Post("/", bind(dtos.GetCardsForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetCards))
		cardRoute.Post("/getById", bind(dtos.GetCardByIdForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetCardByID))
		cardRoute.Post("/create", bind(dtos.CreateCardForm{}), bind(dtos.LocaleForm{}), Wrap(hs.CreateCard))
		cardRoute.Post("/update", bind(dtos.UpdateCardForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateCard))
		cardRoute.Post("/updateSeq", bind(dtos.UpdateCardSeqForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateCardSeq))
		cardRoute.Post("/delete", bind(dtos.DeleteCardForm{}), bind(dtos.LocaleForm{}), Wrap(hs.DeleteCard))
	}, reqSignedIn)

	r.Group("/api/field", func(fieldRoute routing.RouteRegister) {
		fieldRoute.Get("/", bind(dtos.GetFieldsForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetFields))
		fieldRoute.Post("/", bind(dtos.GetFieldsForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetFields))
		fieldRoute.Post("/getById", bind(dtos.GetFieldByIdForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetFieldByID))
		fieldRoute.Post("/create", bind(dtos.CreateFieldForm{}), bind(dtos.LocaleForm{}), Wrap(hs.CreateField))
		fieldRoute.Post("/update", bind(dtos.UpdateFieldForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateField))
		fieldRoute.Post("/updateSeq", bind(dtos.UpdateFieldSeqForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateFieldSeq))
		fieldRoute.Post("/delete", bind(dtos.DeleteFieldForm{}), bind(dtos.LocaleForm{}), Wrap(hs.DeleteField))
	}, reqSignedIn)

	r.Group("/api/action", func(actionRoute routing.RouteRegister) {
		actionRoute.Get("/", bind(dtos.GetViewActionsForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetViewActions))
		actionRoute.Post("/", bind(dtos.GetViewActionsForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetViewActions))
		actionRoute.Post("/getById", bind(dtos.GetViewActionByIdForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetViewActionByID))
		actionRoute.Post("/create", bind(dtos.CreateViewActionForm{}), bind(dtos.LocaleForm{}), Wrap(hs.CreateViewAction))
		actionRoute.Post("/update", bind(dtos.UpdateViewActionForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateViewAction))
		actionRoute.Post("/updateSeq", bind(dtos.UpdateViewActionSeqForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateViewActionSeq))
		actionRoute.Post("/delete", bind(dtos.DeleteViewActionForm{}), bind(dtos.LocaleForm{}), Wrap(hs.DeleteViewAction))
	}, reqSignedIn)

	r.Group("/api/parameter", func(actionRoute routing.RouteRegister) {
		actionRoute.Post("/", bind(dtos.GetParametersForm{}), bind(dtos.LocaleForm{}), Wrap(hs.GetParameters))
		actionRoute.Post("/update", bind(dtos.UpdateParameterForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateParameter))
	}, reqSignedIn)

	r.Group("/api/style", func(actionRoute routing.RouteRegister) {
		actionRoute.Post("/update", bind(dtos.UpdateStylesForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateStyles))
	}, reqSignedIn)

	r.Group("/data", func(actionRoute routing.RouteRegister) {
		actionRoute.Post("/read", bind(dtos.ReadTestDataForm{}), bind(dtos.LocaleForm{}), Wrap(hs.ReadTestData))
		actionRoute.Post("/create", bind(dtos.CreateTestDataForm{}), bind(dtos.LocaleForm{}), Wrap(hs.CreateTestData))
		actionRoute.Post("/update", bind(dtos.UpdateTestDataForm{}), bind(dtos.LocaleForm{}), Wrap(hs.UpdateTestData))
		actionRoute.Post("/delete", bind(dtos.DeleteTestDataForm{}), bind(dtos.LocaleForm{}), Wrap(DeleteTestData))
	}, reqSignedIn)

	r.Get("/*", reqSignedIn, hs.Index)
}
