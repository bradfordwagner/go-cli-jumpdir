# go-cli-jumpdir
Easy jumping to previously visited dirs

## Commands

### `jumpdir add <directory>`
Add directories to the jumpdir database.

### `jumpdir delete <directory>`
Remove directories from the jumpdir database.

### `jumpdir list`
List all directories sorted by weight (descending order). Used for fzf integration.

### `jumpdir leaderboard`
Display directories with their weights in a leaderboard format.

**Flags:**
- `-r, --reverse`: Reverse sort order to ascending (lowest weight first)
- `-n, --number <N>`: Limit output to top N directories

**Examples:**
```bash
# Show all directories with weights (highest first)
jumpdir leaderboard

# Show in ascending order (lowest weight first)
jumpdir leaderboard -r

# Show top 10 directories
jumpdir leaderboard -n 10

# Show top 5 in ascending order
jumpdir leaderboard -r -n 5
```

### `jumpdir addweight <directory> <weight>`
Increment the weight of a specific directory.

