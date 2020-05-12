import { getLocale } from 'umi'
import { message } from 'antd'
import reqwest from 'reqwest'
import JSZip from 'jszip'
import ejs from 'ejs'
import { saveAs } from 'file-saver'
import { IFolder, ICard, ILocale, IViewAction, IField, IStyle } from './interfaces';
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
interface rowCfg {
    Id: number,
    Cols: cardCfg[],
}
interface CardRenderInfo {
    Name: string,
    Locale: string,
    Style: any,
    Pos: number,
    Width: number,
    Seq: number,
    Fields: IField[],
    Actions: IViewAction[],
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
    locales = locales.sort(function (a: any, b: any) { return a.Name.localeCompare(b.Name) })
    console.log(folders)
    render(folders, locales)
}

interface setting {
    xAxis?: IStyle,
    y1Axis?: IStyle,
    y2Axis?: IStyle,
}
interface FolderRender {
    Name: string,
    Locale: string,
    Children: string,
}
interface PageRender {
    Name: string,
    Locale: string,
}
function render(config: IFolder[], localescfg: ILocale[]) {

    var zip = new JSZip();
    let count = 0;
    let compcount = 0;
    var reg1 = /(\n[\s\t]*\r*\n)/g;
    var reg2 = /^[\n\r\n\t]*|[\n\r\n\t]*$/g;
    //source导出
    let sourceFolder = zip.folder("source");
    sourceFolder.file("Navigation.json", JSON.stringify(config.map((folder) => {
        if (folder.IsFolder) {
            return {
                "Name": folder.Name,
                "Locale": folder.Locale.Name,
                "Type": "Folder",
                "Children": folder.Pages.map((pg) => {
                    return {
                        "Name": pg.Name,
                        "Locale": pg.Locale.Name,
                    }
                })
            }
        }
        else {
            if (folder.Pages && folder.Pages.length == 1) {
                return {
                    "Name": folder.Pages[0].Name,
                    "Locale": folder.Pages[0].Locale.Name,
                    "Type": "Page",
                }
            }
            else {
                return {
                    "Name": folder.Name,
                    "Locale": folder.Locale.Name,
                    "Type": "Page",
                }
            }
        }
    }), (k: any, v: any) => { return v; }, '    '))
    config.forEach(folder => {
        if (folder.Pages) {
            folder.Pages.forEach(page => {
                page.Cards.forEach(card => {
                    let cardRenderInfo: CardRenderInfo = {
                        Name: card.Name,
                        Locale: card.Locale.Name ?? card.Name,
                        Style: {
                            Type: card.Style,
                        },
                        Seq: card.Seq,
                        Width: card.Width,
                        Pos: card.Pos,
                        Fields: card.Fields,
                        Actions: card.Actions
                    }
                    if (card.Style === "POINT") {
                        card.Styles.forEach(st => {
                            if (st.Property === "XAXIS" && st.Field) {
                                cardRenderInfo.Style.xAxis = st.Field.Name
                            }
                            else if (st.Property === "Y1AXIS" && st.Field) {
                                cardRenderInfo.Style.y1Axis = st.Field.Name
                                cardRenderInfo.Style.y1Color = st.Value
                            }
                            else if (st.Property === "Y2AXIS" && st.Field) {
                                cardRenderInfo.Style.y2Axis = st.Field.Name
                                cardRenderInfo.Style.y2Color = st.Value
                            }
                        })
                    }
                    else if (card.Style === "RECT") {
                        card.Styles.forEach(st => {
                            if (st.Property === "XAXIS" && st.Field) {
                                cardRenderInfo.Style.xAxis = st.Field.Name
                            }
                            else if (st.Property === "Y1AXIS" && st.Field) {
                                cardRenderInfo.Style.y1Axis = st.Field.Name
                                cardRenderInfo.Style.y1Type = st.Value
                            }
                            else if (st.Property === "Y2AXIS" && st.Field) {
                                cardRenderInfo.Style.y2Axis = st.Field.Name
                                cardRenderInfo.Style.y2Type = st.Value
                            }
                        })
                        card.Styles.forEach(st => {
                            if (st.Property === "Y1COLOR" && cardRenderInfo.Style.y1Axis) {
                                cardRenderInfo.Style.y1Color = st.Value
                            }
                            else if (st.Property === "Y2COLOR" && cardRenderInfo.Style.y2Axis) {
                                cardRenderInfo.Style.y2Color = st.Value
                            }
                        })

                    }
                    sourceFolder.file(card.Name + ".json", JSON.stringify(cardRenderInfo, (k: any, v: any) => {
                        switch (k) {
                            case 'Fields':
                                {
                                    return v.map((f: IField) => {
                                        return ({
                                            "Name": f.Name,
                                            "Locale": f.Locale.Name,
                                            "Seq": f.Seq,
                                            "IsVisible": f.IsVisible,
                                            "Type": f.Type,
                                            "Filter": f.Filter,
                                            "Default": f.Default,
                                        })
                                    })
                                }
                            case 'Actions':
                                {
                                    return v.map((a: IViewAction) => {
                                        return {
                                            "Name": a.Name,
                                            "Locale": a.Locale.Name,
                                            "Type": a.Type,
                                            "Seq": a.Seq,
                                            "DoubleCheck": a.DoubleCheck,
                                            "Parameters": a.Parameters.map((p) => {
                                                return {
                                                    "Name": p.Field.Name,
                                                    "Locale": p.Field.Locale.Name,
                                                    "Seq": p.Field.Seq,
                                                    "Type": p.Field.Type,
                                                    "IsVisible": p.IsVisible,
                                                    "IsEditable": p.IsEditable,
                                                    "Default": p.Default,
                                                    "Compare": p.Compare,
                                                }
                                            })
                                        }
                                    })
                                }
                            default:
                                return v;
                        }
                    }, '    '))
                })
            })
        }
    })
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