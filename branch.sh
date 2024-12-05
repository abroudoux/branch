function branch() {
    git_repositoy=$(git rev-parse --is-inside-work-tree)

    if [[ $git_repositoy != "true" ]]; then
        echo "Error: Not a git repository"
        return 1
    fi

    branches=$(git branch --format="%(refname:short)" 2>/dev/null)

    if [[ -z "$branches" ]]; then
        echo "No branches found"
        return 1
    fi

    while [[ $# -gt 0 ]]; do
        case "$1" in
            --run|-r)
                selected_branch=$(echo "$branches" | fzf --prompt "Press Enter to copy the branch selected: " --height 100%)

                if [[ -n "$selected_branch" ]]; then
                    clean_branch=$(echo "$selected_branch" | xargs)
                    echo -n "$clean_branch" | pbcopy
                    echo "Branch '$clean_branch' copied to clipboard."
                else
                    echo "No branch selected."
                fi

                return
                ;;
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
                echo "  --run, -r             Start the interactive mode"
                echo "  --list, -l            List all branches"
                echo "  --help, -h            Show this help message"
                return
                ;;
        esac
        shift
    done

    selected_branch=$(echo "$branches" | fzf --prompt "Press Enter to copy the branch selected: " --height 100%)

    if [[ -n "$selected_branch" ]]; then
        clean_branch=$(echo "$selected_branch" | xargs)
        echo -n "$clean_branch" | pbcopy
        echo "Branch '$clean_branch' copied to clipboard."
    else
        echo "No branch selected."
    fi
}

function run() {}