#!/bin/bash 

# an example for bash cmdline tab autocompletion

function _solactl() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    opts_first_stage="service server task job mission"
    opts_service_follow="up down start stop restart"
    opts_server_follow="prepare start stop"
    opts_task_follow="pause resume cancel redo info"
    opts_job_follow="create"
    opts_mission_follow="create"

    if [[ ${prev} == *solactl ]]; then 
      COMPREPLY=( $(compgen -W "${opts_first_stage}" -- ${cur}) )
    fi
    
    case ${prev} in 
      service)
        COMPREPLY=( $(compgen -W "${opts_service_follow}" -- ${cur}) )
        return 0
        ;;
      server)
        COMPREPLY=( $(compgen -W "${opts_server_follow}" -- ${cur}) )
        return 0
        ;;
      task)
        COMPREPLY=( $(compgen -W "${opts_task_follow}" -- ${cur}) )
        return 0
        ;;
      job)
        COMPREPLY=( $(compgen -W "${opts_job_follow}" -- ${cur}) )
        return 0
        ;;
      mission)
        COMPREPLY=( $(compgen -W "${opts_mission_follow}" -- ${cur}) )
        return 0
    esac
}
complete -F _solactl solactl
