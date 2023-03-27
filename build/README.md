# Build Directory

The build directory is used to house all the build files and assets for this application. 

The structure is:

* bin - Output directory (not committed)
* darwin - macOS specific build files
* windows - Windows specific build files

## Mac

- `icon.icns` - App bundle icon file.

The `darwin` directory holds files specific to Mac builds.
These may be customised and used as part of the build. 
To return these files to the default state, simply delete them and build with `wails build`.

The directory contains the following files:

- `Info.plist` - the main plist file used for Mac builds. It is used when building using `wails build`.
- `Info.dev.plist` - same as the main plist file but used when building using `wails dev`.

## Windows

The `windows` directory contains the manifest and rc files used when building with `wails build`.
These may be customised for your application. To return these files to the default state, simply delete them and
build with `wails build`.

- `icon.ico` - The icon used for the application. This is used when building using `wails build`.
- `installer/*` - The files used to create the Windows installer. These are used when building using `wails build`.
- `info.json` - Application details used for Windows builds. The data here will be used by the Windows installer,
  as well as the application itself (right click the exe -> properties -> details)
- `wails.exe.manifest` - The main application manifest file.