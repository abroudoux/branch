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

        local function run() {
        selected_option=$(echo "$branches" | fzf --prompt "Select a branch: " --no-info --height 10% --pointer ">") || { echo "Action canceled."; return; }

        if [[ -z "$selected_option" ]]; then
            echo "No branch selected."
            return 1
        fi

        echo "You selected branch: $selected_option"

        action=$(echo -e "Copy\nCheckout\nNew\nMerge\nDelete\nCancel" | fzf --prompt "Choose an action: " --no-info --height 10% --pointer ">") || { echo "Action canceled."; return; }

        if [[ -z "$action" ]]; then
            echo "No action selected."
            return 1
        fi

        if [[ "$action" == "Copy" ]]; then
            echo -n "$selected_option" | pbcopy && git push && echo "Branch '$selected_option' copied to clipboard."
        elif [[ "$action" == "Checkout" ]]; then
            git checkout "$selected_option" && echo "Branch '$selected_option' checked out."
        elif [[ "$action" == "New" ]]; then
            read "name?Enter a name for your new branch: "

            if [[ -z "$name" ]]; then
                echo "No branch name provided. Action canceled."
                return 1
            fi

            git checkout -b "$name" && echo "Branch '$name' based on '$selected_option' created."
        elif [[ "$action" == "Merge" ]]; then
            git merge "$selected_option" && echo "Branch '$selected_option' merged."
        elif [[ "$action" == "Delete" ]]; then
            git branch -D "$selected_option" && echo "Branch '$selected_option' deleted."
        else
            echo "Action canceled."
        fi
    }

    while [[ $# -gt 0 ]]; do
        case "$1" in
            --run|-r)
                run
                return
                ;;
            --list|-l)
                echo "$branches"
                return
                ;;
            --version|-v)
                cat ./ascii.txt
                echo ""
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

    run
}
