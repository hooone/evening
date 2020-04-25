package navigation

import (
	"fmt"
	"sort"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/card"
	"github.com/hooone/evening/pkg/managers/folder"
	"github.com/hooone/evening/pkg/managers/page"
	cardSvr "github.com/hooone/evening/pkg/services/card"
	"github.com/hooone/evening/pkg/services/locale"

	"github.com/hooone/evening/pkg/registry"
)

func init() {
	registry.RegisterService(&NavigationService{})
}

type NavigationService struct {
	Locale      *locale.LocaleService `inject:""`
	CardService *cardSvr.CardService  `inject:""`
}

func (s *NavigationService) Init() error {
	return nil
}

//GetNavigation 获得完整导航栏
func (s *NavigationService) GetNavigation(orgId int64, lang string) ([]*Folder, error) {
	result := make([]*Folder, 0)
	query := folder.GetFolderQuery{
		OrgId: orgId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return result, err
	}
	for _, f := range query.Result {
		fd := Folder{
			Id:       f.Id,
			OrgId:    f.OrgId,
			Name:     f.Name,
			IsFolder: f.IsFolder,
			Seq:      f.Seq,
		}
		s.Locale.GetLocale(&fd, lang)
		pQuery := page.GetPagesQuery{
			FolderId: fd.Id,
			OrgId:    fd.OrgId,
		}
		if err := bus.Dispatch(&pQuery); err != nil {
			return result, err
		}
		pages := make([]*Page, 0)
		for _, p := range pQuery.Result {
			pg := Page{
				Id:       p.Id,
				FolderId: p.FolderId,
				OrgId:    p.OrgId,
				Name:     p.Name,
				Seq:      p.Seq,
			}
			s.Locale.GetLocale(&pg, lang)
			pages = append(pages, &pg)
		}
		sort.Sort(PageSlice(pages))
		//arrange seq
		for idx, pg := range pages {
			if pg.Seq != int32(idx*2+1) {
				pg.Seq = int32(idx*2 + 1)
				seqCmd := page.UpdatePageSeqCommand{
					FolderId: pg.Id,
					Seq:      pg.Seq,
					OrgId:    orgId,
				}
				if err := bus.Dispatch(&seqCmd); err != nil {
					return result, err
				}
			}
		}
		fd.Pages = pages
		result = append(result, &fd)
	}
	sort.Sort(FolderSlice(result))
	return result, nil
}

//CreateFolder 添加文件夹
func (s *NavigationService) CreateFolder(form Folder, orgId int64, lang string) error {
	seq, _ := getFolderEnd(orgId)
	cmd := folder.CreateFolderCommand{
		Name:     form.Name,
		Seq:      seq,
		IsFolder: true,
		OrgId:    orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	form.Id = cmd.Result

	s.Locale.SetLocale(form, lang)
	return nil
}

//UpdateFolder 修改文件夹信息
func (s *NavigationService) UpdateFolder(form Folder, orgId int64, lang string) error {
	cmd := folder.UpdateFolderCommand{
		FolderId: form.Id,
		Name:     form.Name,
		OrgId:    orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	s.Locale.SetLocale(form, lang)
	return nil
}

//DeleteFolder 删除文件夹
func (s *NavigationService) DeleteFolder(form Folder, orgId int64, lang string) error {
	pQuery := page.GetPagesQuery{
		FolderId: form.Id,
		OrgId:    orgId,
	}
	if err := bus.Dispatch(&pQuery); err != nil {
		return err
	}
	for _, p := range pQuery.Result {
		pg := Page{
			Id:       p.Id,
			FolderId: p.FolderId,
			OrgId:    p.OrgId,
			Name:     p.Name,
			Seq:      p.Seq,
		}
		s.DeletePage(pg, orgId, lang)
	}
	cmd := folder.DeleteFolderCommand{
		Id:    form.Id,
		OrgId: orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	//delete locale
	s.Locale.DeleteLocale(form, lang)
	return nil
}

func pageNameCheck(pageName string, orgId int64) error {
	oldPquery := page.GetPageByNameQuery{
		PageName: pageName,
		OrgId:    orgId,
	}
	errP := bus.Dispatch(&oldPquery)
	if errP == nil {
		return ErrPageNameConflict
	}
	if errP != page.ErrPageNotFound {
		return errP
	}
	return nil
}
func getFolderEnd(orgId int64) (int32, error) {
	query := folder.GetFolderQuery{OrgId: orgId}
	if err := bus.Dispatch(&query); err != nil {
		return 1, err
	}
	var seq int32
	seq = 1
	for _, temp := range query.Result {
		if temp.Seq > seq {
			seq = temp.Seq
		}
	}
	seq = seq + 2
	return seq, nil
}
func getPageEnd(pageId int64, orgId int64) (int32, error) {
	query := page.GetPagesQuery{FolderId: pageId, OrgId: orgId}
	if err := bus.Dispatch(&query); err != nil {
		return 1, err
	}
	var seq int32
	seq = 1
	for _, temp := range query.Result {
		if temp.Seq > seq {
			seq = temp.Seq
		}
	}
	seq = seq + 2
	return seq, nil
}

//CreateTreePage 添加一级页面
func (s *NavigationService) CreateTreePage(form Page, orgId int64, lang string) error {
	//重名判断
	if err := pageNameCheck(form.Name, orgId); err != nil {
		return nil
	}
	//顺序计算
	seq, _ := getFolderEnd(orgId)
	//添加文件夹
	cmd := folder.CreateFolderCommand{
		Name:     form.Name,
		Seq:      seq,
		IsFolder: false,
		OrgId:    orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	//添加页面
	pageCmd := page.CreatePageCommand{
		Name:     form.Name,
		Seq:      1,
		FolderId: cmd.Result,
		OrgId:    orgId,
	}
	if err := bus.Dispatch(&pageCmd); err != nil {
		return err
	}
	form.Id = pageCmd.Result
	//设置文本
	s.Locale.SetLocale(form, lang)
	//添加默认Card
	//TBD
	return nil
}

//CreateNodePage 添加子页面
func (s *NavigationService) CreateNodePage(form Page, orgId int64, lang string) (string, error) {
	//重名判断
	if err := pageNameCheck(form.Name, orgId); err != nil {
		return "", nil
	}
	//顺序计算
	seq, _ := getPageEnd(form.FolderId, orgId)
	//查文件夹
	fquery := folder.GetFolderByIDQuery{
		OrgId:    orgId,
		FolderId: form.FolderId,
	}
	if err := bus.Dispatch(&fquery); err != nil {
		return "", err
	}
	//添加页面
	pageCmd := page.CreatePageCommand{
		Name:     form.Name,
		Seq:      seq,
		FolderId: form.FolderId,
		OrgId:    orgId,
	}
	if err := bus.Dispatch(&pageCmd); err != nil {
		return "", err
	}
	form.Id = pageCmd.Result
	//设置文本
	s.Locale.SetLocale(form, lang)
	return "/" + fquery.Result.Name + "/" + form.Name, nil
}

//UpdatePage 修改页面信息
func (s *NavigationService) UpdatePage(form Page, orgId int64, lang string) error {
	cmd := page.UpdatePageCommand{
		PageId: form.Id,
		Name:   form.Name,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	fmt.Println("UpdatePage OK")
	//tree page change folder name
	query := folder.GetFolderByIDQuery{
		OrgId:    orgId,
		FolderId: cmd.Result.FolderId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return err
	}
	if !query.Result.IsFolder {
		fdCmd := folder.UpdateFolderCommand{
			FolderId: query.Result.Id,
			Name:     form.Name,
			OrgId:    orgId,
		}
		if err := bus.Dispatch(&fdCmd); err != nil {
			return err
		}
	}
	//设置文本
	s.Locale.SetLocale(form, lang)
	return nil
}

//DeletePage 删除页面
func (s *NavigationService) DeletePage(form Page, orgId int64, lang string) error {
	//delete card
	cQuery := card.GetCardsQuery{
		PageId: form.Id,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&cQuery); err != nil {
		return err
	}
	for _, c := range cQuery.Result {
		//call delete Card
		s.CardService.DeleteCard(cardSvr.Card{
			Id:     c.Id,
			PageId: c.PageId,
			OrgId:  c.OrgId,
			Name:   c.Name,
			Seq:    c.Seq,
			Pos:    c.Pos,
			Width:  c.Width,
			Style:  c.Style,
		}, orgId, lang)
	}
	//delete page
	cmd := page.DeletePageCommand{
		PageId: form.Id,
		OrgId:  orgId,
	}
	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}
	//delete folder
	query := folder.GetFolderByIDQuery{
		OrgId:    orgId,
		FolderId: cmd.Result.FolderId,
	}
	if err := bus.Dispatch(&query); err != nil {
		return err
	}
	if !query.Result.IsFolder {
		fdCmd := folder.DeleteFolderCommand{
			Id:    query.Result.Id,
			OrgId: orgId,
		}
		if err := bus.Dispatch(&fdCmd); err != nil {
			return err
		}
	}
	//delete locale
	s.Locale.DeleteLocale(form, lang)
	return nil
}

//NavMove 导航栏拖动
func (s *NavigationService) NavMove(SourceFolder int64, SourcePage int64,
	TargetFolder int64, TargetPage int64, Position int, orgId int64) error {
	//get all need data
	sFquery := folder.GetFolderByIDQuery{
		OrgId:    orgId,
		FolderId: SourceFolder,
	}
	if SourceFolder != 0 {
		if err := bus.Dispatch(&sFquery); err != nil {
			return err
		}
	}
	tFquery := folder.GetFolderByIDQuery{
		OrgId:    orgId,
		FolderId: TargetFolder,
	}
	if TargetFolder != 0 {
		if err := bus.Dispatch(&tFquery); err != nil {
			return err
		}
	}
	sPquery := page.GetPageByIDQuery{
		PageId: SourcePage,
		OrgId:  orgId,
	}
	if SourcePage != 0 {
		if err := bus.Dispatch(&sPquery); err != nil {
			return err
		}
		sFquery.FolderId = sPquery.Result.FolderId
		if err := bus.Dispatch(&sFquery); err != nil {
			return err
		}
	}
	tPquery := page.GetPageByIDQuery{
		PageId: TargetPage,
		OrgId:  orgId,
	}
	if TargetPage != 0 {
		if err := bus.Dispatch(&tPquery); err != nil {
			return err
		}
		tFquery.FolderId = tPquery.Result.FolderId
		if err := bus.Dispatch(&tFquery); err != nil {
			return err
		}
	}
	//move
	if SourceFolder != 0 && TargetFolder != 0 {
		//folder to folder
		if sFquery.Result.IsFolder && tFquery.Result.IsFolder {
			moveCmd := folder.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
				OrgId:    orgId,
			}
			if Position == 1 {
				moveCmd.Seq = tFquery.Result.Seq - 1
			}
			if Position == 3 {
				moveCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
		}
	} else if SourceFolder != 0 && TargetPage != 0 {
		if tFquery.Result.IsFolder {
			//folder to node page
			moveCmd := folder.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
				OrgId:    orgId,
			}
			moveCmd.Seq = tFquery.Result.Seq + 1
			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
		} else {
			//folder to tree page
			moveCmd := folder.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
				OrgId:    orgId,
			}
			if Position == 1 {
				moveCmd.Seq = tFquery.Result.Seq - 1
			}
			if Position == 3 {
				moveCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
		}
	} else if SourcePage != 0 && TargetFolder != 0 {
		if sFquery.Result.IsFolder {
			//node page to folder
			if Position == 2 {
				if sFquery.Result.Id != tFquery.Result.Id {
					moveCmd := page.UpdatePageSeqCommand{
						PageId:   sPquery.Result.Id,
						OrgId:    orgId,
						FolderId: tFquery.Result.Id,
						Seq:      99999,
					}
					if err := bus.Dispatch(&moveCmd); err != nil {
						return err
					}
				}
			} else {
				//create fake folder
				createCmd := folder.CreateFolderCommand{
					Name:     sPquery.Result.Name,
					IsFolder: false,
					OrgId:    orgId,
				}
				if Position == 1 {
					createCmd.Seq = tFquery.Result.Seq - 1
				}
				if Position == 3 {
					createCmd.Seq = tFquery.Result.Seq + 1
				}
				if err := bus.Dispatch(&createCmd); err != nil {
					return err
				}
				//move page
				moveCmd := page.UpdatePageSeqCommand{
					PageId:   sPquery.Result.Id,
					OrgId:    orgId,
					FolderId: createCmd.Result,
					Seq:      1,
				}
				if err := bus.Dispatch(&moveCmd); err != nil {
					return err
				}
				//arrange folder
				moveFCmd := folder.UpdateFolderSeqCommand{
					FolderId: createCmd.Result,
					OrgId:    orgId,
					Seq:      createCmd.Seq,
				}
				if err := bus.Dispatch(&moveFCmd); err != nil {
					return err
				}
			}
		} else {
			//tree page to folder
			if Position == 2 {
				if sFquery.Result.Id != tFquery.Result.Id {
					moveCmd := page.UpdatePageSeqCommand{
						PageId:   sPquery.Result.Id,
						FolderId: tFquery.Result.Id,
						OrgId:    orgId,
						Seq:      99999,
					}
					if err := bus.Dispatch(&moveCmd); err != nil {
						return err
					}
					deleteCmd := folder.DeleteFolderCommand{
						Id:    sFquery.Result.Id,
						OrgId: orgId,
					}
					if err := bus.Dispatch(&deleteCmd); err != nil {
						return err
					}
				}
			} else {
				//move folder
				moveCmd := folder.UpdateFolderSeqCommand{
					FolderId: sFquery.Result.Id,
					OrgId:    orgId,
				}
				if Position == 1 {
					moveCmd.Seq = tFquery.Result.Seq - 1
				}
				if Position == 3 {
					moveCmd.Seq = tFquery.Result.Seq + 1
				}
				if err := bus.Dispatch(&moveCmd); err != nil {
					return err
				}
			}
		}
	} else if SourcePage != 0 && TargetPage != 0 {
		if sFquery.Result.IsFolder && tFquery.Result.IsFolder {
			//node page to node page
			moveCmd := page.UpdatePageSeqCommand{
				PageId:   sPquery.Result.Id,
				OrgId:    orgId,
				FolderId: tFquery.Result.Id,
			}
			if Position == 1 {
				moveCmd.Seq = tPquery.Result.Seq - 1
			}
			if Position == 3 {
				moveCmd.Seq = tPquery.Result.Seq + 1
			}

			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
		} else if sFquery.Result.IsFolder && !tFquery.Result.IsFolder {
			//node page to tree page
			//create fake folder
			createCmd := folder.CreateFolderCommand{
				Name:     sPquery.Result.Name,
				IsFolder: false,
				OrgId:    orgId,
			}
			if Position == 1 {
				createCmd.Seq = tFquery.Result.Seq - 1
			}
			if Position == 3 {
				createCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&createCmd); err != nil {
				return err
			}
			//move page
			moveCmd := page.UpdatePageSeqCommand{
				PageId:   sPquery.Result.Id,
				OrgId:    orgId,
				FolderId: createCmd.Result,
				Seq:      1,
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
			//arrange folder
			moveFCmd := folder.UpdateFolderSeqCommand{
				FolderId: createCmd.Result,
				OrgId:    orgId,
				Seq:      createCmd.Seq,
			}
			if err := bus.Dispatch(&moveFCmd); err != nil {
				return err
			}
		} else if !sFquery.Result.IsFolder && tFquery.Result.IsFolder {
			//tree page to node page
			if sFquery.Result.Id != tFquery.Result.Id {
				moveCmd := page.UpdatePageSeqCommand{
					PageId:   sPquery.Result.Id,
					OrgId:    orgId,
					FolderId: tFquery.Result.Id,
				}
				if Position == 1 {
					moveCmd.Seq = tPquery.Result.Seq - 1
				}
				if Position == 3 {
					moveCmd.Seq = tPquery.Result.Seq + 1
				}
				if err := bus.Dispatch(&moveCmd); err != nil {
					return err
				}
				deleteCmd := folder.DeleteFolderCommand{
					Id:    sFquery.Result.Id,
					OrgId: orgId,
				}
				if err := bus.Dispatch(&deleteCmd); err != nil {
					return err
				}
			}
		} else if !sFquery.Result.IsFolder && !tFquery.Result.IsFolder {
			//tree page to tree page
			moveCmd := folder.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
				OrgId:    orgId,
			}
			if Position == 1 {
				moveCmd.Seq = tFquery.Result.Seq - 1
			}
			if Position == 3 {
				moveCmd.Seq = tFquery.Result.Seq + 1
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
		}
	} else if SourceFolder != 0 && TargetFolder == 0 && TargetPage == 0 {
		//folder to empty
		moveCmd := folder.UpdateFolderSeqCommand{
			FolderId: sFquery.Result.Id,
			OrgId:    orgId,
			Seq:      99999,
		}
		if err := bus.Dispatch(&moveCmd); err != nil {
			return err
		}
	} else if SourcePage != 0 && TargetFolder == 0 && TargetPage == 0 {
		if sFquery.Result.IsFolder {
			//node page to empty
			//create fake folder
			createCmd := folder.CreateFolderCommand{
				Name:     sPquery.Result.Name,
				IsFolder: false,
				Seq:      99999,
				OrgId:    orgId,
			}
			if err := bus.Dispatch(&createCmd); err != nil {
				return err
			}
			//move page
			moveCmd := page.UpdatePageSeqCommand{
				PageId:   sPquery.Result.Id,
				OrgId:    orgId,
				FolderId: createCmd.Result,
				Seq:      1,
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
			//arrange folder
			moveFCmd := folder.UpdateFolderSeqCommand{
				FolderId: createCmd.Result,
				OrgId:    orgId,
				Seq:      createCmd.Seq,
			}
			if err := bus.Dispatch(&moveFCmd); err != nil {
				return err
			}
		} else {
			//tree page to empty
			moveCmd := folder.UpdateFolderSeqCommand{
				FolderId: sFquery.Result.Id,
				OrgId:    orgId,
				Seq:      99999,
			}
			if err := bus.Dispatch(&moveCmd); err != nil {
				return err
			}
		}
	}
	return nil
}
