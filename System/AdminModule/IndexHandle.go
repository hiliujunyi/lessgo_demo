package AdminModule

import (
	"github.com/lessgo/lessgo"
)

func IndexHandle(ctx lessgo.Context) error {
	ctx.String(200, "这里是后台首页")
	return nil
}
