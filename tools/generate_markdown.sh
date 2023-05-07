#!/bin/sh
# This is inspired by blog.sh
# https://github.com/karlb/karl.berlin.
set -eu

MARKDOWN="smu"
OUTPUT_DIR="build"
# We use an intermediate file as bash array sucks.
POSTS_TMP_INDEX="/tmp/posts_index.tsv"
PAGES_TMP_INDEX="/tmp/pages_index.tsv"

IFS="	"

build_tsv_index() {
  for f in "$1"/*.md; do
    local title=$(head -n 1 "$f"|sed -e 's/^# //')
    local commit_dates=$(git log --format="%cs" "$f")
    local created_time="draft"
    local updated_time="draft"
    if [ -n "$commit_dates" ]; then
      created_time=$(echo "$commit_dates" | head -n 1)
      updated_time=$(echo "$commit_dates" | tail -n 1)
    fi
    printf "%s\t%s\t%s\t%s\n" "$f" "$title" "$created_time" "$updated_time"
  done
}

build_homepage() {
  local output_path="$2/index.html"
  cat header.html | sed "s/{{TITLE}}/Homepage/" > "$output_path"
  $MARKDOWN "index.md" >> "$output_path"
  while read -r f title create_time updated_time; do
    local relative_path=$(echo "$f" | sed -e 's/.md$/.html/')
    printf "%s&nbsp;<a href=\"/%s\">%s</a>\n" "$create_time" "$relative_path" "$title" >> "$output_path"
  done < "$1"
  cat footer.html >> "$output_path"
}

build_html() {
  local output_dir="$2"
  while read -r f title create_time updated_time; do
    echo "Processing file: $f"
    # The first line is the title of the post
    local html_file=$(echo "$f" | sed -e 's/.md$/.html/')
    local output_path="$output_dir/$html_file"
    echo "Output file: $output_path"

		cat header.html | sed "s/{{TITLE}}/$title/" > "$output_path"
    $MARKDOWN "$f" >> "$output_path"
    cat footer.html >> "$output_path"
  done < "$1"
}

# We store the index in reverse order so the newest posts are first.
echo "Building the indexes of posts and pages"
build_tsv_index "posts" | sort -rt "	" -k 3 > $POSTS_TMP_INDEX
build_tsv_index "pages" > $PAGES_TMP_INDEX

# Start from a clean state.
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR/posts"
mkdir -p "$OUTPUT_DIR/pages"
# Copy anything under html.
cp -R html/* "$OUTPUT_DIR"

# Build the index.html then individual pages.
echo "Building homepage"
build_homepage "$POSTS_TMP_INDEX" "$OUTPUT_DIR"
echo "Building posts"
build_html "$POSTS_TMP_INDEX" "$OUTPUT_DIR"
echo "Building pages"
build_html "$PAGES_TMP_INDEX" "$OUTPUT_DIR"
