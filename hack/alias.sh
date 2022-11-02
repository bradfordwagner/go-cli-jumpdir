jd() {
  dir=$(jumpdir list | fzf)
  if [ $dir != "" ]; then
    jumpdir addweight ${dir}
    cd ${dir}
  fi
}