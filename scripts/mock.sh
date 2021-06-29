for d in pkg/* ; do
    mockery --dir="$d" --all --output=./internal/mocks
done
for d in third_party/* ; do
    mockery --dir="$d" --all --output=./internal/mocks
done