export I_PACKAGE="git.ucode.space/Phil/goshorly/utils"
export I_GitCommitShort=$(git rev-parse --short HEAD)
export I_GitBranch=$(git rev-parse --abbrev-ref HEAD)

go mod download

go build -a -installsuffix -ldflags="-w -s -X $I_PACKAGE.GitCommitShort=$I_GitCommitShort -X $I_PACKAGE.GitBranch=$I_GitBranch" -o app .