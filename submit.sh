#!/usr/bin/env sh

if [ "$#" -lt 2 ] || [ "$#" -gt 3 ]; then
    echo "Usage: sh $0 <difficulty> <task_number> [--sql]"
    exit 1
fi

check_difficulty() {
    case "$1" in
        easy|medium|hard) return 0 ;;
        *) return 1 ;;
    esac
}

difficulty=$1
task_number=$2
base_dir="algo"

if [ "$#" -eq 3 ] && [ "$3" = "--sql" ]; then
    base_dir="sql"
fi

check_difficulty "$difficulty" || {
    echo "Error: Difficulty must be one of: easy, medium, hard"
    exit 1
}

cd ~/Programs/alGO-adventures/ || exit 1

path="$base_dir/$difficulty"

git add "$path"
git commit -m "sln: $path/$task_number"
git push origin main

echo "Done!"
