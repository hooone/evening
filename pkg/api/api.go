package api

import (
	"github.com/go-macaron/binding"
	"github.com/hooone/evening/pkg/api/dtos"
	"github.com/hooone/evening/pkg/api/routing"
	"github.com/hooone/evening/pkg/middleware"
)

func (hs *HTTPServer) registerRoutes() {
	reqSignedIn := middleware.ReqSignedIn
	bind := binding.Bind

	r := hs.RouteRegister
	r.Get("/logout", hs.Logout)
	r.Post("/login", bind(dtos.LoginCommand{}), Wrap(hs.LoginPost))
	r.Get("/login", hs.LoginView)

	r.Group("/api/navigation", func(navRoute routing.RouteRegister) {
		navRoute.Get("/", Wrap(GetNavigation))
		navRoute.Post("/", Wrap(GetNavigation))
		navRoute.Post("/createFolder", bind(dtos.CreateFolderForm{}), bind(dtos.LocaleForm{}), Wrap(CreateFolder))
		navRoute.Post("/updateFolder", bind(dtos.UpdateFolderForm{}), bind(dtos.LocaleForm{}), Wrap(UpdateFolder))
		navRoute.Post("/deleteFolder", bind(dtos.DeleteFolderForm{}), bind(dtos.LocaleForm{}), Wrap(DeleteFolder))
		navRoute.Post("/createTreePage", bind(dtos.CreateTreePageForm{}), bind(dtos.LocaleForm{}), Wrap(CreateTreePage))
		navRoute.Post("/createNodePage", bind(dtos.CreateNodePageForm{}), bind(dtos.LocaleForm{}), Wrap(CreateNodePage))
		navRoute.Post("/updatePage", bind(dtos.UpdatePageForm{}), bind(dtos.LocaleForm{}), Wrap(UpdatePage))
		navRoute.Post("/deletePage", bind(dtos.DeletePageForm{}), bind(dtos.LocaleForm{}), Wrap(DeletePage))
		navRoute.Post("/move", bind(dtos.NavMoveForm{}), bind(dtos.LocaleForm{}), Wrap(NavMove))
	}, reqSignedIn)
	r.Group("/api/card", func(cardRoute routing.RouteRegister) {
		cardRoute.Get("/", bind(dtos.GetCardsForm{}), Wrap(GetCards))
		cardRoute.Post("/", bind(dtos.GetCardsForm{}), Wrap(GetCards))
		cardRoute.Post("/getById", bind(dtos.GetCardByIdForm{}), Wrap(GetCardByID))
		cardRoute.Post("/create", bind(dtos.CreateCardForm{}), bind(dtos.LocaleForm{}), Wrap(CreateCard))
		cardRoute.Post("/update", bind(dtos.UpdateCardForm{}), bind(dtos.LocaleForm{}), Wrap(UpdateCard))
		cardRoute.Post("/updateSeq", bind(dtos.UpdateCardSeqForm{}), Wrap(UpdateCardSeq))
		cardRoute.Post("/delete", bind(dtos.DeleteCardForm{}), Wrap(DeleteCard))
	}, reqSignedIn)
	r.Group("/api/field", func(fieldRoute routing.RouteRegister) {
		fieldRoute.Get("/", bind(dtos.GetFieldsForm{}), Wrap(GetFields))
		fieldRoute.Post("/", bind(dtos.GetFieldsForm{}), Wrap(GetFields))
		fieldRoute.Post("/getById", bind(dtos.GetFieldByIdForm{}), Wrap(GetFieldByID))
		fieldRoute.Post("/create", bind(dtos.CreateFieldForm{}), bind(dtos.LocaleForm{}), Wrap(CreateField))
		fieldRoute.Post("/update", bind(dtos.UpdateFieldForm{}), bind(dtos.LocaleForm{}), Wrap(UpdateField))
		fieldRoute.Post("/updateSeq", bind(dtos.UpdateFieldSeqForm{}), bind(dtos.LocaleForm{}), Wrap(UpdateFieldSeq))
		fieldRoute.Post("/delete", bind(dtos.DeleteFieldForm{}), Wrap(DeleteField))
	}, reqSignedIn)
	r.Group("/api/action", func(actionRoute routing.RouteRegister) {
		actionRoute.Get("/", bind(dtos.GetViewActionsForm{}), Wrap(GetViewActions))
		actionRoute.Post("/", bind(dtos.GetViewActionsForm{}), Wrap(GetViewActions))
		actionRoute.Post("/getById", bind(dtos.GetViewActionByIdForm{}), Wrap(GetViewActionByID))
		actionRoute.Post("/create", bind(dtos.CreateViewActionForm{}), bind(dtos.LocaleForm{}), Wrap(CreateViewAction))
		actionRoute.Post("/update", bind(dtos.UpdateViewActionForm{}), bind(dtos.LocaleForm{}), Wrap(UpdateViewAction))
		actionRoute.Post("/updateSeq", bind(dtos.UpdateViewActionSeqForm{}), bind(dtos.LocaleForm{}), Wrap(UpdateViewActionSeq))
		actionRoute.Post("/delete", bind(dtos.DeleteViewActionForm{}), Wrap(DeleteViewAction))
	}, reqSignedIn)
	r.Group("/api/parameter", func(actionRoute routing.RouteRegister) {
		actionRoute.Post("/update", bind(dtos.UpdateParameterForm{}), Wrap(UpdateParameter))
	}, reqSignedIn)
	r.Group("/api/style", func(actionRoute routing.RouteRegister) {
		actionRoute.Post("/update", bind(dtos.UpdateStylesForm{}), Wrap(UpdateStyles))
	}, reqSignedIn)
	r.Group("/data", func(actionRoute routing.RouteRegister) {
		actionRoute.Post("/read", bind(dtos.ReadTestDataForm{}), Wrap(ReadTestData))
		actionRoute.Post("/create", bind(dtos.CreateTestDataForm{}), Wrap(CreateTestData))
		actionRoute.Post("/update", bind(dtos.UpdateTestDataForm{}), Wrap(UpdateTestData))
		actionRoute.Post("/delete", bind(dtos.DeleteTestDataForm{}), Wrap(DeleteTestData))
	}, reqSignedIn)

	r.Get("/*", reqSignedIn, hs.Index)
}
