set -ex
# get current branch
BRANCH=$(git rev-parse --abbrev-ref HEAD)
git checkout test
git pull origin test
git merge $BRANCH --no-ff --no-edit
git push origin test
git checkout $BRANCH
