function branch() {
    git_repositoy=$(git rev-parse --is-inside-work-tree)

    if [[ $git_repositoy != "true" ]]; then
        echo "Error: Not a git repository"
        return 1
    fi

    branches=$(git branch)

    if [[ -z "$branches" ]]; then
        echo "No branches found"
        return 1
    fi

    while [[ $# -gt 0 ]]; do
        case "$1" in
            --list|-l)
                echo $branches
                return
                ;;
        esac
        shift
    done
}