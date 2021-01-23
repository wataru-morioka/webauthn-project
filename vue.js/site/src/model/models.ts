export class PhotoInfo {
    public id: number;
    public subTitle: string;
    public title: string;
    public mimetype: string;
    public data: Buffer;
    public createdDatetime: string;

    constructor(id: number, subTitle: string, title: string, mimetype: string, fileName: string,
                size: number, data: Buffer, createdDatetime: string, modifiedDatetime: string) {
        this.id = id;
        this.subTitle = subTitle;
        this.title = title;
        this.mimetype = mimetype;
        this.data = data;
        this.createdDatetime = createdDatetime;
    }
}

export class VideoInfo {
  public mimetype: string;
  public data: Buffer;

  constructor(mimetype: string, data: Buffer) {
        this.mimetype = mimetype;
        this.data = data;
  }
}

export class Article {
    public id: number;
    public contributorName: string;
    public body: string;
    public thumbnail: Uint8Array;
    public createdDatetime: string;
    public modifiedDatetime: string;

    constructor(id: number, contributorName: string, body: string, thumbnail: Uint8Array,
                createdDatetime: string, modifiedDatetime: string) {
        this.id = id;
        this.contributorName = contributorName;
        this.body = body;
        this.thumbnail = thumbnail;
        this.createdDatetime = createdDatetime;
        this.modifiedDatetime = modifiedDatetime;
    }
}

export class ArticleInfo {
    public id: number;
    public contributorName: string;
    public body: string;
    public thumbnail: Uint8Array;
    public createdDatetime: string;
    public modifiedDatetime: string;
    public commentatorName: string;
    public commentatorThumbnail: Uint8Array;
    public commentBody: string;
    public commentCreatedDatetime: string;

    constructor(id: number, contributorName: string, body: string, thumbnail: Uint8Array,
                createdDatetime: string, modifiedDatetime: string, commentatorName: string,
                commentatorThumbnail: Uint8Array, commentBody: string, commentCreatedDatetime: string) {
        this.id = id;
        this.contributorName = contributorName;
        this.body = body;
        this.thumbnail = thumbnail;
        this.createdDatetime = createdDatetime;
        this.modifiedDatetime = modifiedDatetime;
        this.commentatorName = commentatorName;
        this.commentatorThumbnail = commentatorThumbnail;
        this.commentBody = commentBody;
        this.commentCreatedDatetime = commentCreatedDatetime;
    }
}

export class Comment {
    public name: string;
    public body: string;
    public thumbnail: Uint8Array;
    public createdDatetime: string;

    constructor(name: string, body: string, thumbnail: Uint8Array, createdDatetime: string) {
        this.name = name;
        this.body = body;
        this.thumbnail = thumbnail;
        this.createdDatetime = createdDatetime;
    }
}

export class AccountInfo {
    public uid: string;
    public deleteFlag: boolean;
    public webrtcFlag: boolean;
    public adminFlag: boolean;
    public account: string;
    public name: string;
    public state: string;
    public loginCount: number;
    public latestLogin: string;
    public createdDatetime: string;
    public modifiedDatetime: string;

    constructor(uid: string, deleteFlag: boolean, webrtcFlag: boolean, adminFlag: boolean, account: string,
                name: string, state: string, loginCount: number, latestLogin: string,
                createdDatetime: string, modifiedDatetime: string) {
        this.uid = uid;
        this.deleteFlag = deleteFlag;
        this.webrtcFlag = webrtcFlag;
        this.adminFlag = adminFlag;
        this.account = account;
        this.name = name;
        this.state = state;
        this.loginCount = loginCount;
        this.latestLogin = latestLogin;
        this.createdDatetime = createdDatetime;
        this.modifiedDatetime = modifiedDatetime;
    }
}

export class ContactInfo {
    public id: number;
    public createdDatetime: string;
    public account: string;
    public name: string;
    public organization: string;
    public state: string;
    public email: string;
    public phone: string;
    public message: string;

    constructor(id: number, createdDatetime: string, account: string, name: string, organization: string, state: string,
                email: string, phone: string, message: string) {
        this.id = id;
        this.createdDatetime = createdDatetime,
        this.account = account;
        this.name = name;
        this.organization = organization;
        this.state = state;
        this.email = email;
        this.phone = phone;
        this.message = message;
    }
}

export default PhotoInfo;

