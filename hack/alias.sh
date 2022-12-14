# jump directory
jd() {
  dir=$(jumpdir list | fzf)
  if [ $dir != "" ]; then
    jumpdir addweight ${dir}
    cd ${dir}
  fi
}

# short for jd tmux
jdt() {
  dir=$(jumpdir list | fzf)
  if [ $dir != "" ]; then
    jumpdir addweight ${dir}
    dir_name=$(basename ${dir})
    tmux new-window -n ${dir_name} -c "${dir}"
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
  jd
  ${EDITOR} .
}

