package utils

var (
	CI_COMMIT_SHORT_SHA string
	CI_COMMIT_BRANCH    string
	// TODO: check if Tag build
	CI_COMMIT_TAG string
	CI_BUILD      bool
)

func Init_build_vars() {
	if CI_COMMIT_SHORT_SHA == "" && CI_COMMIT_BRANCH == "" {
		CI_BUILD = false
	} else {
		CI_BUILD = true
	}
}
