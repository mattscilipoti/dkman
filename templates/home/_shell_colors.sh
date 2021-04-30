
# --------------------
# Provides colors as Environment Variables
# e.g. To make git info yellow then reset to default: ${YELLOW}$(__git_ps1)${DEFAULT_COLOR}
#
# This is a shell scripting "partial"
# Designed to be sourced by other shell scripts.
# This is indicated by the underscore prefix (a Rails convention)
# --------------------

# Set the TERM var to xterm-256color
if ( set +u && [[ $COLORTERM == gnome-* && $TERM == xterm ]]) && infocmp gnome-256color >/dev/null 2>&1; then
  export TERM=gnome-256color
elif infocmp xterm-256color >/dev/null 2>&1; then
  export TERM=xterm-256color
fi
if tput setaf 1 &> /dev/null; then
  tput sgr0
  if [[ $(tput colors) -ge 256 ]] 2>/dev/null; then
    # this is for xterm-256color
    BLACK=$(tput setaf 0)
    RED=$(tput setaf 1)
    GREEN=$(tput setaf 2)
    YELLOW=$(tput setaf 226)
    BLUE=$(tput setaf 4)
    MAGENTA=$(tput setaf 5)
    CYAN=$(tput setaf 6)
    WHITE=$(tput setaf 7)
    ORANGE=$(tput setaf 172)
    PURPLE=$(tput setaf 141)
    BG_BLACK=$(tput setab 0)
    BG_RED=$(tput setab 1)
    BG_GREEN=$(tput setab 2)
    BG_BLUE=$(tput setab 4)
    BG_MAGENTA=$(tput setab 5)
    BG_CYAN=$(tput setab 6)
    BG_YELLOW=$(tput setab 226)
    BG_ORANGE=$(tput setab 172)
    BG_WHITE=$(tput setab 7)
  else
    MAGENTA=$(tput setaf 5)
    ORANGE=$(tput setaf 4)
    GREEN=$(tput setaf 2)
    PURPLE=$(tput setaf 1)
    WHITE=$(tput setaf 7)
  fi
  BOLD=$(tput bold)
  DEFAULT_COLOR=$(tput sgr0)
  UNDERLINE=$(tput sgr 0 1)
else
  BLACK="\[\e[0;30m\]"
  RED="\033[1;31m"
  ORANGE="\033[1;33m"
  GREEN="\033[1;32m"
  PURPLE="\033[1;35m"
  WHITE="\033[1;37m"
  YELLOW="\[\e[0;33m\]"
  CYAN="\[\e[0;36m\]"
  BLUE="\[\e[0;34m\]"
  BOLD=""
  DEFAULT_COLOR="\033[m"
fi

# Returns colorized text
# - text is whatever text you want colorized
# - text_color is the color name as a Capitalized String (e.g. RED, ORANGE, YELLOW)
#   -- available colors are listed above
#  Use Case: used as label for messages
colorize_text() {
  # WORKAROUND: wrap in "set +x" to avoid debugging noise
  ( set +x;
    local text=$1;
    local text_color=$2;
    printf "${text_color}$text${DEFAULT_COLOR}"
  )
}
