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
	export class Exclusion {
	    foreignId: string;
	    artistName: string;
	    id?: number;
	
	    static createFrom(source: any = {}) {
	        return new Exclusion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.foreignId = source["foreignId"];
	        this.artistName = source["artistName"];
	        this.id = source["id"];
	    }
	}
	export class ImportListInput {
	    enableAutomaticAdd: boolean;
	    shouldMonitorExisting: boolean;
	    shouldSearch: boolean;
	    listOrder: number;
	    id?: number;
	    qualityProfileId?: number;
	    metadataProfileId?: number;
	    configContract?: string;
	    implementation?: string;
	    listType?: string;
	    monitorNewItems?: string;
	    name?: string;
	    rootFolderPath?: string;
	    shouldMonitor?: string;
	    tags?: number[];
	    fields?: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new ImportListInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAutomaticAdd = source["enableAutomaticAdd"];
	        this.shouldMonitorExisting = source["shouldMonitorExisting"];
	        this.shouldSearch = source["shouldSearch"];
	        this.listOrder = source["listOrder"];
	        this.id = source["id"];
	        this.qualityProfileId = source["qualityProfileId"];
	        this.metadataProfileId = source["metadataProfileId"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.listType = source["listType"];
	        this.monitorNewItems = source["monitorNewItems"];
	        this.name = source["name"];
	        this.rootFolderPath = source["rootFolderPath"];
	        this.shouldMonitor = source["shouldMonitor"];
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
	export class QualityProfile {
	    id: number;
	    name: string;
	    upgradeAllowed: boolean;
	    cutoff: number;
	    items: starr.Quality[];
	    minFormatScore: number;
	    cutoffFormatScore: number;
	    formatItems: starr.FormatItem[];
	
	    static createFrom(source: any = {}) {
	        return new QualityProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.upgradeAllowed = source["upgradeAllowed"];
	        this.cutoff = source["cutoff"];
	        this.items = this.convertValues(source["items"], starr.Quality);
	        this.minFormatScore = source["minFormatScore"];
	        this.cutoffFormatScore = source["cutoffFormatScore"];
	        this.formatItems = this.convertValues(source["formatItems"], starr.FormatItem);
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
	export class Exclusion {
	    tmdbId: number;
	    movieTitle: string;
	    movieYear: number;
	    id?: number;
	
	    static createFrom(source: any = {}) {
	        return new Exclusion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tmdbId = source["tmdbId"];
	        this.movieTitle = source["movieTitle"];
	        this.movieYear = source["movieYear"];
	        this.id = source["id"];
	    }
	}
	export class ImportListInput {
	    enableAuto: boolean;
	    enabled: boolean;
	    searchOnAdd: boolean;
	    listOrder: number;
	    id?: number;
	    qualityProfileId?: number;
	    configContract?: string;
	    implementation?: string;
	    implementationName?: string;
	    infoLink?: string;
	    listType?: string;
	    monitor?: string;
	    name?: string;
	    rootFolderPath?: string;
	    minimumAvailability?: string;
	    tags?: number[];
	    fields?: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new ImportListInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAuto = source["enableAuto"];
	        this.enabled = source["enabled"];
	        this.searchOnAdd = source["searchOnAdd"];
	        this.listOrder = source["listOrder"];
	        this.id = source["id"];
	        this.qualityProfileId = source["qualityProfileId"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.implementationName = source["implementationName"];
	        this.infoLink = source["infoLink"];
	        this.listType = source["listType"];
	        this.monitor = source["monitor"];
	        this.name = source["name"];
	        this.rootFolderPath = source["rootFolderPath"];
	        this.minimumAvailability = source["minimumAvailability"];
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
	export class QualityProfile {
	    id?: number;
	    name?: string;
	    upgradeAllowed: boolean;
	    cutoff: number;
	    items?: starr.Quality[];
	    minFormatScore: number;
	    cutoffFormatScore: number;
	    formatItems: starr.FormatItem[];
	    language?: starr.Value;
	
	    static createFrom(source: any = {}) {
	        return new QualityProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.upgradeAllowed = source["upgradeAllowed"];
	        this.cutoff = source["cutoff"];
	        this.items = this.convertValues(source["items"], starr.Quality);
	        this.minFormatScore = source["minFormatScore"];
	        this.cutoffFormatScore = source["cutoffFormatScore"];
	        this.formatItems = this.convertValues(source["formatItems"], starr.FormatItem);
	        this.language = this.convertValues(source["language"], starr.Value);
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
	    implementationName: string;
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
	        this.implementationName = source["implementationName"];
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
	export class Exclusion {
	    foreignId: string;
	    authorName: string;
	    id?: number;
	
	    static createFrom(source: any = {}) {
	        return new Exclusion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.foreignId = source["foreignId"];
	        this.authorName = source["authorName"];
	        this.id = source["id"];
	    }
	}
	export class ImportListInput {
	    enableAutomaticAdd: boolean;
	    shouldMonitorExisting: boolean;
	    shouldSearch: boolean;
	    listOrder: number;
	    id?: number;
	    metadataProfileId?: number;
	    qualityProfileId?: number;
	    listType?: string;
	    configContract?: string;
	    implementation?: string;
	    name?: string;
	    rootFolderPath?: string;
	    shouldMonitor?: string;
	    monitorNewItems?: string;
	    tags?: number[];
	    fields?: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new ImportListInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAutomaticAdd = source["enableAutomaticAdd"];
	        this.shouldMonitorExisting = source["shouldMonitorExisting"];
	        this.shouldSearch = source["shouldSearch"];
	        this.listOrder = source["listOrder"];
	        this.id = source["id"];
	        this.metadataProfileId = source["metadataProfileId"];
	        this.qualityProfileId = source["qualityProfileId"];
	        this.listType = source["listType"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.name = source["name"];
	        this.rootFolderPath = source["rootFolderPath"];
	        this.shouldMonitor = source["shouldMonitor"];
	        this.monitorNewItems = source["monitorNewItems"];
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
	export class QualityProfile {
	    id: number;
	    name: string;
	    upgradeAllowed: boolean;
	    cutoff: number;
	    items: starr.Quality[];
	    minFormatScore: number;
	    cutoffFormatScore: number;
	    formatItems: starr.FormatItem[];
	
	    static createFrom(source: any = {}) {
	        return new QualityProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.upgradeAllowed = source["upgradeAllowed"];
	        this.cutoff = source["cutoff"];
	        this.items = this.convertValues(source["items"], starr.Quality);
	        this.minFormatScore = source["minFormatScore"];
	        this.cutoffFormatScore = source["cutoffFormatScore"];
	        this.formatItems = this.convertValues(source["formatItems"], starr.FormatItem);
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
	export class Exclusion {
	    tvdbId: number;
	    title: string;
	    id?: number;
	
	    static createFrom(source: any = {}) {
	        return new Exclusion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tvdbId = source["tvdbId"];
	        this.title = source["title"];
	        this.id = source["id"];
	    }
	}
	export class ImportListInput {
	    enableAutomaticAdd: boolean;
	    seasonFolder: boolean;
	    listOrder: number;
	    qualityProfileId?: number;
	    id?: number;
	    configContract?: string;
	    implementation?: string;
	    implementationName?: string;
	    infoLink?: string;
	    listType?: string;
	    minRefreshInterval?: string;
	    name?: string;
	    rootFolderPath?: string;
	    seriesType?: string;
	    shouldMonitor?: string;
	    tags?: number[];
	    fields?: starr.FieldInput[];
	
	    static createFrom(source: any = {}) {
	        return new ImportListInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.enableAutomaticAdd = source["enableAutomaticAdd"];
	        this.seasonFolder = source["seasonFolder"];
	        this.listOrder = source["listOrder"];
	        this.qualityProfileId = source["qualityProfileId"];
	        this.id = source["id"];
	        this.configContract = source["configContract"];
	        this.implementation = source["implementation"];
	        this.implementationName = source["implementationName"];
	        this.infoLink = source["infoLink"];
	        this.listType = source["listType"];
	        this.minRefreshInterval = source["minRefreshInterval"];
	        this.name = source["name"];
	        this.rootFolderPath = source["rootFolderPath"];
	        this.seriesType = source["seriesType"];
	        this.shouldMonitor = source["shouldMonitor"];
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
	export class QualityProfile {
	    upgradeAllowed: boolean;
	    id: number;
	    cutoff: number;
	    name: string;
	    items: starr.Quality[];
	    minFormatScore: number;
	    cutoffFormatScore: number;
	    formatItems?: starr.FormatItem[];
	    language?: starr.Value;
	
	    static createFrom(source: any = {}) {
	        return new QualityProfile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.upgradeAllowed = source["upgradeAllowed"];
	        this.id = source["id"];
	        this.cutoff = source["cutoff"];
	        this.name = source["name"];
	        this.items = this.convertValues(source["items"], starr.Quality);
	        this.minFormatScore = source["minFormatScore"];
	        this.cutoffFormatScore = source["cutoffFormatScore"];
	        this.formatItems = this.convertValues(source["formatItems"], starr.FormatItem);
	        this.language = this.convertValues(source["language"], starr.Value);
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

export namespace starr {
	
	export class Value {
	    id: number;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new Value(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	    }
	}

}

