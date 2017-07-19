# drone-hg2

Very simple plugins supporting cloning of mercurial repositories. It makes use of netrc infrastructure and injects https authentication into .hgrc for cloning.

Pull requests or any other fancy features are not supported.

## Usage

Plugin parameters are defined in the yaml file, note that some fake parameters has to be supplied (in this case, the depth):

```
clone:
  hg:
  	image: krakonos/drone-hg2
	depth: 50
```

## Bugs

Probably a lot! Please report them as github issues.
