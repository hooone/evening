import { getLocale } from 'umi'
import { message } from 'antd'
import reqwest from 'reqwest'
import JSZip from 'jszip'
import ejs from 'ejs'
import { saveAs } from 'file-saver'
import { IFolder, ICard, ILocale } from './interfaces';
import page from './models/page'

interface cardCfg {
    Id: number,
    Pos: number,
    Width: number,
    Children: ICard[],
}
interface rowCfg {
    Id: number,
    Cols: cardCfg[],
}

export function ExportUmi() {
    let pagecount = 0;
    let comppcount = 0;
    reqwest({
        url: '/api/renderdata',
        type: 'json',
        method: 'post',
        data: { lang: getLocale(), }
    }).then(function (rst: any) {
        if (!rst.Success) {
            message.error(rst.Message);
            return
        }
        let folders: IFolder[] = rst.Data;
        console.log(folders)
        handler(folders)
    })
}

function handler(folders: IFolder[]) {
    //布局处理
    folders.forEach(fd => {
        if (fd.Pages) {
            fd.Pages.forEach(pg => {
                pg.Rows = [{
                    Id: 0,
                    Cols: []
                }];
                pg.Cards.forEach(
                    card => {
                        let cols: cardCfg[] = pg.Rows[pg.Rows.length - 1].Cols
                        let oldCol = false;
                        let newCol = false;
                        cols.forEach(col => {
                            if (col.Pos === card.Pos && col.Width === card.Width) {
                                oldCol = true
                                col.Id = card.Id
                                col.Children.push(card)
                            }
                        });
                        if (!oldCol) {
                            newCol = true
                            cols.forEach(col => {
                                if ((col.Pos <= card.Pos && (col.Pos + col.Width) > card.Pos) ||
                                    (col.Pos < (card.Pos + card.Width) && (col.Pos + col.Width) >= (card.Pos + card.Width)) ||
                                    (col.Pos >= card.Pos && (col.Pos + col.Width) <= (card.Pos + card.Width))) {
                                    newCol = false
                                }
                            });
                            if (newCol) {
                                cols.push({
                                    Id: card.Id,
                                    Pos: card.Pos,
                                    Width: card.Width,
                                    Children: [card]
                                })
                            } else {
                                pg.Rows.push({
                                    Id: card.Id,
                                    Cols: [{
                                        Id: card.Id,
                                        Pos: card.Pos,
                                        Width: card.Width,
                                        Children: [card]
                                    }]
                                })
                            }
                        }
                    }
                )
            })
        }
    })

    //国际化处理
    let locales: ILocale[] = [];
    folders.forEach(fd => {
        fd.Locale.Default = fd.Name
        fd.Locale.Name = fd.Name + "_folder"
        locales.push(fd.Locale)
        fd.Pages.forEach(page => {
            page.Locale.Default = page.Name
            page.Locale.Name = page.Name + "_page"
            locales.push(page.Locale)
            page.Cards.forEach(card => {
                card.Locale.Default = card.Name
                card.Locale.Name = card.Name + "_card"
                locales.push(card.Locale)
                card.Fields.forEach(field => {
                    field.Locale.Default = field.Name
                    field.Locale.Name = field.Name
                    locales.push(field.Locale)
                })
                card.Actions.forEach(action => {
                    action.Locale.Default = action.Name
                    action.Locale.Name = action.Name + card.Name.replace(card.Name[0], card.Name[0].toUpperCase())
                    locales.push(action.Locale)
                    action.Parameters.forEach(param => {
                        param.Field.Locale.Default = param.Field.Name
                        param.Field.Locale.Name = param.Field.Name
                    })
                })
            })
        })
    })
    locales = locales.sort(function (a, b) { return a.Name.localeCompare(b.Name) })
    console.log(folders)
    render(folders, locales)
}

function render(config: IFolder[], localescfg: ILocale[]) {

    var zip = new JSZip();
    let count = 0;
    let compcount = 0;
    var reg1 = /(\n[\s\t]*\r*\n)/g;
    var reg2 = /^[\n\r\n\t]*|[\n\r\n\t]*$/g;

    reqwest({
        url: '/render/list',
        type: 'text',
        method: 'get'
    }).then(function (list: any) {
        let files: string[] = list.responseText.split('\n');
        files = files.map(str => str.replace(String.fromCharCode(13), ""));
        console.log(files)
        let statics: string[] = [];
        let nav: string[] = [];
        let pages: string[] = [];
        let cards: string[] = [];
        let locales: string[] = [];
        files.forEach(file => {
            let sp = file.split(":");
            if (sp[0] === "~") {
                statics.push(sp[1]);
                count++;
            }
            else if (sp[0] === 'nav') {
                nav.push(sp[1]);
                count++;
            }
            else if (sp[0] === 'page') {
                pages.push(sp[1]);
                count++;
            }
            else if (sp[0] === 'card') {
                cards.push(sp[1]);
                count++;
            }
            else if (sp[0] === 'locale') {
                locales.push(sp[1]);
                count++;
            }
        })

        statics.forEach(tempPath => {
            reqwest({
                url: '/render/' + tempPath,
                type: 'text',
                method: 'get'
            }).then(function (temp: any) {
                let zipFolder = zip;
                let tempPathSp = tempPath.split("/")
                for (let index = 0; index < tempPathSp.length - 1; index++) {
                    zipFolder = zipFolder.folder(tempPathSp[index]);
                }
                let filename = tempPath.split("/").pop();
                zipFolder.file(filename, temp.responseText.replace(reg1, "\n").replace(reg2, ""));
                compcount++;
                console.log(filename)
                console.log(compcount + '/' + count)
                if (compcount === count) {
                    save(zip);
                }
            }, function () {
                let filename = tempPath.split("/").pop();
                compcount++;
                console.log(filename + ' error')
                console.log(compcount + '/' + count)
                if (compcount === count) {
                    save(zip);
                }
            });
        })


        nav.forEach(tempPath => {
            reqwest({
                url: '/render/' + tempPath,
                type: 'text',
                method: 'get'
            }).then(function (temp: any) {
                let zipFolder = zip;
                let tempPathSp = tempPath.split("/")
                for (let index = 0; index < tempPathSp.length - 1; index++) {
                    zipFolder = zipFolder.folder(tempPathSp[index]);
                }
                let filename = tempPath.split("/").pop();
                let content = ejs.render(temp.responseText, { "folders": config })
                zipFolder.file(filename, content.replace(reg1, "\n").replace(reg2, ""));
                compcount++;
                console.log(filename)
                console.log(compcount + '/' + count)
                if (compcount === count) {
                    save(zip);
                }
            }, function () {
                let filename = tempPath.split("/").pop();
                compcount++;
                console.log(filename + ' error')
                console.log(compcount + '/' + count)
                if (compcount === count) {
                    save(zip);
                }
            });
        })
        console.log(localescfg)
        locales.forEach(tempPath => {
            reqwest({
                url: '/render/' + tempPath,
                type: 'text',
                method: 'get'
            }).then(function (temp: any) {
                let zipFolder = zip;
                let tempPathSp = tempPath.split("/")
                for (let index = 0; index < tempPathSp.length - 1; index++) {
                    zipFolder = zipFolder.folder(tempPathSp[index]);
                }
                let filename = tempPath.split("/").pop();
                let content = ejs.render(temp.responseText, { "locales": localescfg })
                zipFolder.file(filename, content.replace(reg1, "\n").replace(reg2, ""));
                compcount++;
                console.log(filename)
                console.log(compcount + '/' + count)
                if (compcount === count) {
                    save(zip);
                }
            }, function () {
                let filename = tempPath.split("/").pop();
                compcount++;
                console.log(filename + ' error')
                console.log(compcount + '/' + count)
                if (compcount === count) {
                    save(zip);
                }
            });
        })

        pages.forEach(pagePath => {
            let psp = pagePath.split("/");
            reqwest({
                url: '/render/' + pagePath,
                type: 'text',
                method: 'get'
            }).then(function (pagetemp: any) {
                let pagePathSp = pagePath.split("/");
                let zipFolder = zip;
                for (let index = 0; index < pagePathSp.length - 1; index++) {
                    zipFolder = zipFolder.folder(pagePathSp[index]);
                }
                let filename = pagePathSp.pop() ?? "";
                config.forEach(folder => {
                    if (folder.Pages) {
                        folder.Pages.forEach(page => {
                            if (folder.IsFolder)
                                page.FolderLocale = folder.Locale;
                            let ppcontent = ejs.render(pagetemp.responseText, { "page": page })
                            zipFolder.file(filename.replace("{page}", page.Name), ppcontent.replace(reg1, "\n").replace(reg2, ""));
                        })
                    }
                })

                compcount++;
                if (compcount === count) {
                    save(zip);
                }
            }, function () {
                compcount++;
                if (compcount === count) {
                    save(zip);
                }
            })
        })

        cards.forEach(cardPath => {
            let psp = cardPath.split("/");
            reqwest({
                url: '/render/' + cardPath,
                type: 'text',
                method: 'get'
            }).then(function (cardtemp: any) {
                let cardPathSp = cardPath.split("/");
                let zipFolder = zip;
                for (let index = 0; index < cardPathSp.length - 1; index++) {
                    zipFolder = zipFolder.folder(cardPathSp[index]);
                }
                let filename = cardPathSp.pop() ?? "";
                config.forEach(folder => {
                    if (folder.Pages) {
                        folder.Pages.forEach(page => {
                            page.Cards.forEach(card => {
                                let ppcontent = ejs.render(cardtemp.responseText, { "card": card })
                                zipFolder.file(filename.replace("{card}", card.Name), ppcontent.replace(reg1, "\n").replace(reg2, ""));
                            })
                        })
                    }
                })

                compcount++;
                if (compcount === count) {
                    save(zip);
                }
            }, function () {
                compcount++;
                if (compcount === count) {
                    save(zip);
                }
            })
        })
    })
}

function save(zip: any) {
    console.log("save")
    zip.generateAsync({ type: "blob" })
        .then(function (content: any) {
            // see FileSaver.js
            saveAs(content, "example.zip");
        });
}