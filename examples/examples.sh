# sort file by size
alias lt='du -sh * | sort -h'

# succint mnt
alias mnt='mount | grep -E ^/dev | column -t'

# my function
hello() {
 echo "World"
}

# see all bash history
alias bash-history='history|grep bash'

# returns file count
alias count-files='find . -type f | wc -l'

# move to trash
alias trash='mv --force -t ~/.Trash'

# find with preview
alias preview="find . | fzf --preview 'bat --theme={} --color=always {}'"

alias fsearch='fzf'
alias lt2='du * abc'
