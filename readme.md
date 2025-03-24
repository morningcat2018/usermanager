
```
echo "# usermanager" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:morningcat2018/usermanager.git
git push -u origin main
```

```
(base) ➜  user-manager pwd
/Users/morningcat/Documents/user-manager
(base) ➜  user-manager export GOPATH=/Users/morningcat/Documents/user-manager
(base) ➜  user-manager echo $GOPATH
/Users/morningcat/Documents/user-manager
(base) ➜  user-manager go get -u github.com/gin-gonic/gin
go: warning: ignoring go.mod in $GOPATH /Users/morningcat/Documents/user-manager
```