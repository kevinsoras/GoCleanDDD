module.exports = {
  '*.go': ['gofumpt -w', 'goimports -w'],
  '*.{js,json,yml,yaml,md}': ['prettier --write']
};
