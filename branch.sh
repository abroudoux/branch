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

    default_branch=$(git symbolic-ref --short HEAD)

    branches_with_default=""

    while IFS= read -r branch; do
        if [[ "$branch" == "$default_branch" ]]; then
            branches_with_default+="* $branch"$'\n'
        else
            branches_with_default+="  $branch"$'\n'
        fi
    done <<< "$branches"

    local function run() {
        selected_option=$(echo "$branches_with_default" | fzf --prompt "Select a branch: " --no-info --height 10% --pointer ">") || { echo "Action canceled."; return; }

        if [[ -z "$selected_option" ]]; then
            echo "No branch selected."
            return 1
        fi

        selected_option=$(echo "$selected_option" | sed 's/^[ *]*//;s/[ *]*$//')

        echo "You selected branch: $selected_option"

        action=$(echo -e "(N)ame\n(C)heckout\n(B)ranch\n(M)erge\n(D)elete\n(E)xit" | fzf --prompt "Choose an action: " --no-info --height 10% --pointer ">") || { echo "Action canceled."; return; }

        if [[ -z "$action" ]]; then
            echo "No action selected."
            return 1
        fi

        if [[ "$action" == "(N)ame" ]]; then
            echo -n "$selected_option" | pbcopy && git push && echo "Branch '$selected_option' copied to clipboard."
        elif [[ "$action" == "(C)heckout" ]]; then
            git checkout "$selected_option" && echo "Branch '$selected_option' checked out."
        elif [[ "$action" == "(B)ranch" ]]; then
            read "name?Enter a name for your new branch: "

            if [[ -z "$name" ]]; then
                echo "No branch name provided. Action canceled."
                return 1
            fi

            read "checkout?Do you want to checkout on $name? (y/n): "

            if [[ "$checkout" =~ ^(yes|y|Y)$ ]]; then
                git checkout -b "$name" && echo "Branch '$name' created and checked out."
            else
                git branch "$name" && echo "Branch '$name' based on '$selected_option' created."
            fi
        elif [[ "$action" == "(M)erge" ]]; then
            git merge "$selected_option" && echo "Branch '$selected_option' merged."
        elif [[ "$action" == "(D)elete" ]]; then
            git branch -D "$selected_option" && echo "Branch '$selected_option' deleted."
        else
            echo "Action canceled."
        fi
    }

    while [[ $# -gt 0 ]]; do
        case "$1" in
            run|-r)
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
                echo "  run, -r               Start the interactive mode"
                echo "  --list, -l            List all branches"
                echo "  --help, -h            Show this help message"
                return
                ;;
        esac
        shift
    done

    run
}
