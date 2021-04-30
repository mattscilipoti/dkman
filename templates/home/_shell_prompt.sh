source $HOME/_shell_colors.sh
source $HOME/git-prompt.sh # provides _git_ps1

git_info() {
  __git_ps1 # required add'l file
}

go_version() {
  # converts "go version go1.16.3 linux/amd64"
  # to "go1.16.3"
  go version | cut -f3 -d" "
}

ruby_version() {
  # converts "ruby 2.4.1p111 (2017-03-22 revision 58053) [x86_64-linux-musl]"
  # to "2.4.1"
  ruby -v | cut -f2 -d" " | cut -f1 -dp
}

# ---------------------
# style the prompt
# ---------------------

style_user="\[${DEFAULT_COLOR}${WHITE}\]"
style_path="\[${DEFAULT_COLOR}${CYAN}\]"
style_chars="\[${DEFAULT_COLOR}${WHITE}\]"

# ---------------------
# Build the prompt
# ---------------------

# Example with committed changes: username ~/documents/GA/wdi on master[+]
# Use backslash before $ so it isn't expanded immediately: https://stackoverflow.com/a/5380073
prompt=""
prompt+="\t"  # curent time, 24hr mode
prompt+=" 🐳:" # grabbed from https://www.fileformat.info/info/unicode/char/1f433/browsertest.htm
# prompt+="\u" # Username
# prompt+="${WHITE}${DOCKER_CONTAINER_NAME}:"
prompt+="${CYAN}\w" # Working directory
if [ -x "$(command -v ruby)" ]; then
  # prompt+="\$(~/.rvm/bin/rvm-prompt)" # Ruby version via rvm
  ruby_icon='💎'
  prompt+=" ${RED}ruby-\$(ruby_version)"
fi
if [ -x "$(command -v go)" ]; then
  prompt+=" ${DEFAULT_COLOR}$(go_version)"
fi
if [ -x "$(command -v git)" ]; then
  # prompt+='$(prompt_git)' 
  git_icon='ee 82 a0'
  git_icon='Ĳ'    
  git_icon='λ' 
  prompt+="${YELLOW}$(git_info)${DEFAULT_COLOR}"
fi

# prompt+="\n"                                # Newline
prompt+="${style_chars}\$ \[${DEFAULT_COLOR}\]"     # $ (and reset color)

export PS1=$prompt
