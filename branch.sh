function branch() {
    git_repository=$(git rev-parse --is-inside-work-tree)

    if [[ $git_repository != "true" ]]; then
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
                selected_branch=$(echo "$branches" | fzf --prompt "Press Enter to copy the branch selected: " --height 100%) || exit

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
                echo "$branches"
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

    selected_option=$(echo "$branches" | fzf --prompt "Select a branch: " --height 10%) || { echo "Action canceled."; return; }

    if [[ -z "$selected_option" ]]; then
        echo "No branch selected."
        return 1
    fi

    if [[ "$selected_option" == "Create a new branch" ]]; then
        new_branch_name=$(fzf --prompt "Enter the new branch name: " --height 10%) || { echo "Action canceled."; return; }

        if [[ -n "$new_branch_name" ]]; then
            git checkout -b "$new_branch_name" && echo "New branch '$new_branch_name' created."
        else
            echo "No branch name provided. Action canceled."
        fi
    else
        echo "You selected branch: $selected_option"

        action=$(echo -e "Checkout\nDelete\nCancel" | fzf --prompt "Choose an action: " --height 10%) || { echo "Action canceled."; return; }

        if [[ "$action" == "Checkout" ]]; then
            git checkout "$selected_option" && echo "Branch '$selected_option' checked out."
        elif [[ "$action" == "Delete" ]]; then
            git branch -D "$selected_option" && echo "Branch '$selected_option' deleted."
        else
            echo "Action canceled."
        fi
    fi
}
