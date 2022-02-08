NAME="jrxy_login"
OUTPUT_DIR="output"
RELEASE_DIR="releases"

rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

PLATFORMS="$PLATFORMS windows/amd64 windows/386" # arm compilation not available for Windows
PLATFORMS="$PLATFORMS linux/amd64 linux/386"


for PLATFORM in $PLATFORMS; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_FILENAME="${OUTPUT_DIR}/${NAME}-${GOOS}-${GOARCH}"
  if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
  CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -v -ldflags \"-s -w -extldflags '-static'\" -o ${BIN_FILENAME} "
  echo $CMD
  eval $CMD || FAILURES="${FAILURES} ${PLATFORM}"
done