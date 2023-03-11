// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {sync} from '../models';
import {app} from '../models';

export function CheckUpdate():Promise<any>;

export function CreateShortcut():Promise<string>;

export function DownloadUpdate():Promise<string>;

export function GetConfig():Promise<any>;

export function IsWindows():Promise<boolean>;

export function LaunchInstaller():Promise<string>;

export function Lock():Promise<void>;

export function PickFolder(arg1:string):Promise<string>;

export function Quit():Promise<void>;

export function RLock():Promise<void>;

export function RLocker():Promise<sync.Locker>;

export function RUnlock():Promise<void>;

export function SaveConfigItem(arg1:string,arg2:string,arg3:boolean):Promise<string>;

export function TryLock():Promise<boolean>;

export function TryRLock():Promise<boolean>;

export function Unlock():Promise<void>;

export function Version():Promise<app.Version>;
