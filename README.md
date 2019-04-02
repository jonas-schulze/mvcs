# Meta/Multi Version Control System (MVCS)

Did you ever have to clone or pull from several repositories at once? E.g.

* all the different parts of your app
* several lecture notes as well as all the shared `LaTeX` macros
* all the projects of the students you are supervising or your peers

This tool is a small wrapper around the VCS of your choice to do just that.
It has been tested with [mercurial] and [git].

# Usage

```bash
$ mvcs clone
$ mvcs pull
```

# Configuration

`mvcs` looks for a `Repositories.toml` in the current directory:

```toml
vcs = "git"

[[remotes]]

baseUrl = "git@github.com:jonas-schulze/"
repositories = [
  "dotfiles",
]
```

If you want to use `git`, the repositories must skip the trailing `.git`. It
will be added internally. This is so that configurations for different VCS look
more alike.

[git]: https://git-scm.com
[mercurial]: https://www.mercurial-scm.org
