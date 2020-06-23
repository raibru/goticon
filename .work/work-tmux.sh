#!/bin/bash
#
#


session="work"

tmux new-session -d -s $session -n shell \; \
  split-window -h \; \
  new-window -n 'develop' \; \
    split-window -h \; \
    split-window -v \; \
    select-pane -t 0 \; \
  new-window -n 'runtime-1' \; \
    send-keys 'cd runtime/develop && clear' C-m \; \
    select-pane -t 0 \;

tmux select-window -t $session:1
tmux attach -t $session

