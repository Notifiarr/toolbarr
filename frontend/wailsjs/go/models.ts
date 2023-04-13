export namespace lidarr {
	
	export class DownloadClientInput {
	    enable: boolean;
	    removeCompletedDownloads: boolean;
	    removeFailedDownloads: boolean;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new DownloadClientInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enable = source["enable"];
	        this.removeCompletedDownloads = source["removeCompletedDownloads"];
	        this.removeFailedDownloads = source["removeFailedDownloads"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class IndexerInput {
	    enableAutomaticSearch: boolean;
	    enableInteractiveSearch: boolean;
	    enableRss: boolean;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new IndexerInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAutomaticSearch = source["enableAutomaticSearch"];
	        this.enableInteractiveSearch = source["enableInteractiveSearch"];
	        this.enableRss = source["enableRss"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace prowlarr {
	
	export class DownloadClientInput {
	    enable: boolean;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new DownloadClientInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enable = source["enable"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class IndexerInput {
	    enable: boolean;
	    redirect: boolean;
	    priority: number;
	    id?: number;
	    appProfileId: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags?: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new IndexerInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enable = source["enable"];
	        this.redirect = source["redirect"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.appProfileId = source["appProfileId"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace radarr {
	
	export class DownloadClientInput {
	    enable: boolean;
	    removeCompletedDownloads: boolean;
	    removeFailedDownloads: boolean;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new DownloadClientInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enable = source["enable"];
	        this.removeCompletedDownloads = source["removeCompletedDownloads"];
	        this.removeFailedDownloads = source["removeFailedDownloads"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class IndexerInput {
	    enableAutomaticSearch: boolean;
	    enableInteractiveSearch: boolean;
	    enableRss: boolean;
	    downloadClientId: number;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new IndexerInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAutomaticSearch = source["enableAutomaticSearch"];
	        this.enableInteractiveSearch = source["enableInteractiveSearch"];
	        this.enableRss = source["enableRss"];
	        this.downloadClientId = source["downloadClientId"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace readarr {
	
	export class DownloadClientInput {
	    enable: boolean;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new DownloadClientInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enable = source["enable"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class IndexerInput {
	    enableAutomaticSearch: boolean;
	    enableInteractiveSearch: boolean;
	    enableRss: boolean;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new IndexerInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAutomaticSearch = source["enableAutomaticSearch"];
	        this.enableInteractiveSearch = source["enableInteractiveSearch"];
	        this.enableRss = source["enableRss"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace sonarr {
	
	export class DownloadClientInput {
	    enable: boolean;
	    removeCompletedDownloads: boolean;
	    removeFailedDownloads: boolean;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new DownloadClientInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enable = source["enable"];
	        this.removeCompletedDownloads = source["removeCompletedDownloads"];
	        this.removeFailedDownloads = source["removeFailedDownloads"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class IndexerInput {
	    enableAutomaticSearch: boolean;
	    enableInteractiveSearch: boolean;
	    enableRss: boolean;
	    downloadClientId: number;
	    priority: number;
	    id?: number;
	    configContract: string;
	    implementation: string;
	    name: string;
	    protocol: string;
	    tags: number[];
	    fields: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new IndexerInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAutomaticSearch = source["enableAutomaticSearch"];
	        this.enableInteractiveSearch = source["enableInteractiveSearch"];
	        this.enableRss = source["enableRss"];
	        this.downloadClientId = source["downloadClientId"];
	        this.priority = source["priority"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.protocol = source["protocol"];
	        this.tags = source["tags"];
	        this.fields = this.convertValues(source["fields"], starr.FieldInput);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

