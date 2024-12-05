function branch() {
    branches=$(git branch)

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