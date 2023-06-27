# jump directory
jd() {
  dir=$(jumpdir list | fzf)
  if [ $dir != "" ]; then
    jumpdir addweight ${dir}
    cd ${dir}
  else
    # ensure exit code is returned to prevent follow up calls
    return 1
  fi
}

# short for jd tmux
jdt() {
  dir=$(jumpdir list | fzf)
  if [ $dir != "" ]; then
    jumpdir addweight ${dir}
    dir_name=$(basename ${dir})
    tmux new-window -n ${dir_name} -c "${dir}"
  else
    # ensure exit code is returned to prevent follow up calls
    return 1
  fi
}

# jump directory add
jda() {
  jumpdir add $(pwd -P)
}

# list options
jdl() {
  jumpdir list
}

# delete helper
jdd() {
  dir=$(jumpdir list | fzf)
  if [ $dir != "" ]; then
    jumpdir delete ${dir}
  fi
}

# jumpdir vim
jdv() {
  jd && eval ${jump_dir_editor_cmd:-${editor:-vim}}
}
# jumpdir tmux vim
jdtv() {
  jdt && tmux send 'eval ${jump_dir_editor_cmd:-${EDITOR:-vim}}' Enter
}
