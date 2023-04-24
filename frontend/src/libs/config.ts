import { writable, readable } from "svelte/store"
import { GetConfig, Version } from "../../wailsjs/go/app/App"

// This controls tabs and other various pieces of the app.
export type StarrApp = 
    "Lidarr" | 
    "Prowlarr" | 
    "Radarr" | 
    "Readarr" | 
    "Sonarr" | 
    "Whisparr"

export const AllStarrs: StarrApp[] = ["Lidarr", "Prowlarr", "Radarr", "Readarr", "Sonarr", "Whisparr"]

// These are the things we tie to to the 'Hide' menu.
export type HideValues = StarrApp | "Settings" | "Dark"

// This matches AppConfig type in starrs Go mmodule.
export type Instance = {
    App:     StarrApp // Radarr, Sonarr, etc
	Name:    string   // Custom name: Radarr2, Radarr4k, etc.
	URL:     string   // url to app.
	User:    string   // username for app.
	Pass:    string   // password for app.
	Key:     string   // api key for app.
	DBPath:  string   // path to database for app.
	SSL:     boolean  // verify ssl cert?
    Form:    boolean  // Use form login? vs basic auth.
    Timeout: number   // How long to wait for the app API.
}

export type AppConfig = {
    Dark:      boolean
    DevMode:   boolean
    Lang:      string
    Instances: {[key: string]: Instance[]}
    Hide:      {[key: string]: boolean}
    Instance:  {[key: string]: number} // Default Instance.
}

export type AppData = {
    IsWindows: boolean
	IsLinux:   boolean
	IsMac:     boolean
	StartTime: number
    Name:      string
	Title:     string
	Version:   string
	Revision:  string
	Branch:    string
	BuildUser: string
	BuildDate: string
	GoVersion: string
	Started:   string
	Exe:       string
	Home:      string
	Username:  string
}

let config: AppConfig = {
	Dark: false,
	DevMode: false,
	Lang: "en",
	Instances: {},
	Hide: {},
	Instance: {},
}

let appData: Partial<AppData> = {}

export const conf = writable(config, set => { GetConfig().then(v => set(v)) })
export const app =  readable(appData, set => { Version().then(v => set(v)) })
