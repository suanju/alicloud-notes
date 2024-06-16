export namespace aliyunpan {
	
	export class UserInfo {
	    domainId: string;
	    fileDriveId: string;
	    safeBoxDriveId: string;
	    albumDriveId: string;
	    resourceDriveId: string;
	    userId: string;
	    userName: string;
	    createdAt: string;
	    email: string;
	    phone: string;
	    role: string;
	    status: string;
	    nickname: string;
	    totalSize: number;
	    usedSize: number;
	
	    static createFrom(source: any = {}) {
	        return new UserInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.domainId = source["domainId"];
	        this.fileDriveId = source["fileDriveId"];
	        this.safeBoxDriveId = source["safeBoxDriveId"];
	        this.albumDriveId = source["albumDriveId"];
	        this.resourceDriveId = source["resourceDriveId"];
	        this.userId = source["userId"];
	        this.userName = source["userName"];
	        this.createdAt = source["createdAt"];
	        this.email = source["email"];
	        this.phone = source["phone"];
	        this.role = source["role"];
	        this.status = source["status"];
	        this.nickname = source["nickname"];
	        this.totalSize = source["totalSize"];
	        this.usedSize = source["usedSize"];
	    }
	}

}

export namespace types {
	
	export class CreatInstanceReq {
	    refreshToken: string;
	
	    static createFrom(source: any = {}) {
	        return new CreatInstanceReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.refreshToken = source["refreshToken"];
	    }
	}
	export class GenerateQrcodeResp {
	    qrCode: string;
	    t: number;
	    ck: string;
	
	    static createFrom(source: any = {}) {
	        return new GenerateQrcodeResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.qrCode = source["qrCode"];
	        this.t = source["t"];
	        this.ck = source["ck"];
	    }
	}
	export class QrcodeStateReq {
	    t: number;
	    ck: string;
	
	    static createFrom(source: any = {}) {
	        return new QrcodeStateReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.t = source["t"];
	        this.ck = source["ck"];
	    }
	}
	export class QrcodeStateResp {
	    userName: string;
	    userId: string;
	    tokenType: string;
	    refreshToken: string;
	
	    static createFrom(source: any = {}) {
	        return new QrcodeStateResp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.userName = source["userName"];
	        this.userId = source["userId"];
	        this.tokenType = source["tokenType"];
	        this.refreshToken = source["refreshToken"];
	    }
	}

}

