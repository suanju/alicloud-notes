export namespace types {
	
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

