for d in pkg/* ; do
    mockery --dir="$d" --all --output=./internal/mocks
done
