# gotestsuite
golang tookie

```bash
go get github.com/azhao1981/gotestsuite
```


```golang
import "github.com/azhao1981/gotestsuite"

func TestHello(t *testing.T) {
}
```

## 发布

```bash
git tag v1.0.0
git push origin v1.0.0
go list -m github.com/azhao1981/gotestsuite@v1.0.0
```
