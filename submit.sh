#!/usr/bin/env sh

if [ "$#" -lt 2 ] || [ "$#" -gt 3 ]; then
    echo "Usage: sh $0 <difficulty> <task_number> [--sql]"
    exit 1
fi

check_difficulty() {
    VALUE=$1
    if [ "$VALUE" = "easy" ] || [ "$VALUE" = "medium" ] || [ "$VALUE" = "hard" ]; then
        return 0
    fi
    return 1
}

difficulty=$1
task_number=$2
sql_flag=false

if [ "$#" -eq 3 ] && [ "$3" = "--sql" ]; then
    sql_flag=true
fi

check_difficulty "$difficulty"
if [ $? -ne 0 ]; then
    echo "Error: Difficulty must be one of: easy, medium, hard"
    exit 1
fi

cd ~/Programs/alGO-adventures/ || exit 1

if [ "$sql_flag" = true ]; then
    git add "sql/$difficulty"
    git commit -m "sln: sql/$difficulty/$task_number"
else
    git add "$difficulty"
    git commit -m "sln: $difficulty/$task_number"
fi

git push origin main

echo "Done!"

exit 0
