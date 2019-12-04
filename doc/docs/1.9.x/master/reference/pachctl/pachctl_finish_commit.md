## pachctl finish commit

Finish a started commit.

### Synopsis

Finish a started commit. Commit-id must be a writeable commit.

```
pachctl finish commit <repo>@<branch-or-commit> [flags]
```

### Options

```
      --description string   A description of this commit's contents (synonym for --message)
  -h, --help                 help for commit
  -m, --message string       A description of this commit's contents (overwrites any existing commit description)
```

### Options inherited from parent commands

```
      --no-color   Turn off colors.
  -v, --verbose    Output verbose logs
```

### SEE ALSO

* [pachctl finish](pachctl_finish.md)	 - Finish a Pachyderm resource.

###### Auto generated by spf13/cobra on 4-Dec-2019