#!/bin/bash
# Script to update module path from GitLab to GitHub

OLD_PATH="gitlab-master.nvidia.com/urg/go-inflect"
NEW_PATH="github.com/cv/go-inflect"

echo "Updating module path:"
echo "  From: $OLD_PATH"
echo "  To:   $NEW_PATH"
echo

# Find and update all relevant files
files=$(grep -rl "$OLD_PATH" --include="*.go" --include="*.mod" --include="*.md" .)

if [ -z "$files" ]; then
    echo "No files found containing the old path."
    exit 0
fi

for file in $files; do
    echo "Updating: $file"
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS requires '' after -i
        sed -i '' "s|$OLD_PATH|$NEW_PATH|g" "$file"
    else
        # Linux
        sed -i "s|$OLD_PATH|$NEW_PATH|g" "$file"
    fi
done

echo
echo "Done! Updated files:"
grep -rl "$NEW_PATH" --include="*.go" --include="*.mod" --include="*.md" .

echo
echo "Verification:"
grep -r "$NEW_PATH" --include="*.go" --include="*.mod" --include="*.md" .
