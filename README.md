# hackit-quickstart
A gradle project generator written in go. Generates gradle projects that use episode6's hackit open source tools.

## Why
I wanted a faster way to create new libraries and apps using the episode6/hackit toolkit. I also wanted to learn go.

## Usage
Install with
```bash
# You'll also need a relatively modern version of gradle to run the initial wrapper task
go get -u github.com/episode6/hackit-quickstart
```
First, ensure you `cd` to a git repo with a remote named origin (url is used to pre-populate some of [deployable](https://github.com/episode6/deployable)'s properties)
```bash
git clone git@github.com:<myusername>/<mynewproject>.git
cd <mynewproject>

# OR

mkdir <mynewproject>
cd <mynewproject>
git init
git remote add origin git@github.com:<myusername>/<mynewproject>.git
```
Then execute the `hackit-quickstart` command (if you exclude any required flags you'll be asked for them)
```bash
Usage of hackit-quickstart:
-config string
    path to config file (default "~/.hackit-quickstart")
-gdmc string
    Url of a shared gdmc repo to add as a sub-module
-group string
    GroupId (aka package name) of library to generate
-lang string
    Language of project to create. Valid values are
      androidApp: An android application
      androidAppDagger: An android application with a default dagger 2 implementation
      java: A deployable java library
      groovy: A deployable groovy library
      gradle: A deployable groovy library with the gradle api and an empty gradle plugin.
      android: A deployable android library
-licenseName string
    The name of the license you want to use (for deployable libraries) (default "The MIT License (MIT)")
-name string
    The name of the new module to generate (for a multi-module project, this will be the sub-modules name)
-noGdmcRepo
    Dont use a gdmc repo, equivilent to gdmc=""
-type string
    Type of project to create. Valid values are
      sub: A new submodule in an existing multi-module project
      single: A single-module project
      multi: A multi-module project with a single sub-module to start
-version string
    Initial version name to use (default "0.0.1-SNAPSHOT")
```

If you want to use a gdmc repository by default, set it as an environment variable named `GDMC`
```bash
# this repo will be included as a sub-module by default in /gdmc
export GDMC="git@github.com:episode6/hackit-gdmc.git"
```

## Generated Projects
Generated projects include some or all of the following classpath dependencies by default...
- gradle wrapper v3.3 - All projects
- [gdmc](https://github.com/episode6/gdmc) - All projects - dependency manager
- [deployable](https://github.com/episode6/deployable) - All projects except androidApp - simplified publishing to mavenCentral
- [android gradle build tools](https://developer.android.com/studio/releases/gradle-plugin.html) - android and androidApp

## License
MIT: https://github.com/episode6/hackit-quickstart/blob/master/LICENSE
