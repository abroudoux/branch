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

    if [[ $# -eq 0 ]]; then
        echo "You must provide an option"
        ecjo "Use --help or -h for usage information."
        return
    fi

    while [[ $# -gt 0 ]]; do
        case "$1" in
            --list|-l)
                echo $branches
                return
                ;;
            --version|-v)
                version=$(jq -r '.version' ./package.json) 
                echo "$version"
                return
                ;;
            --help|-h)
                echo "Usage: branch [options]"
                echo "Options:"
                echo "  --list, -l            List all branches"
                echo "  --help, -h            Show this help message"
                return
                ;;
        esac
        shift
    done
}