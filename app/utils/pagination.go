package utils

type Paginator struct {
  Page int `form:"page,default=1" binding:"min=1"`
  Size int `form:"size,default=25" binding:"min=1,max=100"`
}
