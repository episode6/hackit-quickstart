# hackit-quickstart
A gradle project generator written in go. Generates gradle projects that use episode6's hackit open source tools.

## Why
I wanted a faster way to create new libraries and apps using the episode6/hackit toolkit. I also wanted to learn go.

## Usage
Install with
```bash
# You'll also need a relatively modern version of gradle to run the initial wrapper task
go get -u github.com/episode6/hackit-quickstart/...
```
`cd` to the root directory of your project (or a new directory) and execute the `hackit-quickstart` command (if you exclude any required flags you'll be asked for them). Some features will require your working dir to be a git repo with a valid remote.
```bash
Usage of hackit-quickstart:

-androidCompileSdkVersion string
    For android apps/libs, the value of compileSdkVersion (default "26")
-androidNdkDir string
    Android ndk directory (default "/android/sdk/ndk-bundle")
-androidSdkDir string
    Android sdk directory (default "/android/sdk")
-config string
    path to config file (default "/Users/ghackett/.hackit-quickstart")
-deployable
    Make a deployable library (has no effect on apps)
-gdmc string
    Url of a shared gdmc repo to add as a sub-module
-gradleVersion string
    Gradle version to apply to the project (root project only) (default "4.4")
-group string
    GroupId (aka package name) of library to generate
-lang string
    Language of project to create. Valid values are
  java: A deployable java library
  groovy: A deployable groovy library
  gradle: A deployable groovy library with the gradle api and an empty gradle plugin.
  android: An android library
  androidApp: An android application
  androidAppDagger: An android application with a default dagger 2 implementation
-licenseName string
    The name of the license you want to use (for deployable libraries) (default "The MIT License (MIT)")
-name string
    The name of the new module to generate (for a multi-module project, this will be the sub-modules name)
-type string
    Type of project to create. Valid values are
  single: A single-module project
  multi: A multi-module project with a single sub-module to start
  sub: A new submodule in an existing multi-module project
-v	Display hackit-quickstart version
-version string
    Initial version name to use (default "0.0.1-SNAPSHOT")
```

`hackit-quickstart` can be configured via a file or environment variables. By default the app checks for a file at `~/.hackit-quickstart`, but that can be overridden using the `-config` flag. For episode6 projects, I use the following config...
```bash
# hackit-quickstart default config

gdmc git@github.com:episode6/hackit-gdmc.git
androidCompileSdkVersion gdmcVersion('android.compilesdk') as Integer
```

## Generated Projects
Generated projects include some or all of the following classpath dependencies by default...
- gradle wrapper v4.4 - All projects
- [gdmc](https://github.com/episode6/gdmc) - All projects - dependency manager
- [deployable](https://github.com/episode6/deployable) - All projects except androidApp (when `-deployable` flag is used)- simplified publishing to mavenCentral
- [android gradle build tools](https://developer.android.com/studio/releases/gradle-plugin.html) - android and androidApp

## License
MIT: https://github.com/episode6/hackit-quickstart/blob/master/LICENSE
