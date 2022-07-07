//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

type PublicPhotoinc struct {
	PhotoincID int64  `db:"photoinc_id,pk" json:"photoincId" form:"photoincId" validate:"required"`
	RefTable   string `db:"ref_table,use_zero" json:"refTable" form:"refTable" validate:"required,lte=50"`
	FolderInc  int64  `db:"folder_inc,use_zero" json:"folderInc" form:"folderInc" validate:"required"`
	FolderName string `db:"folder_name,use_zero" json:"folderName" form:"folderName" validate:"required,lte=50"`
	Running    int64  `db:"running,use_zero" json:"running" form:"running" validate:"required"`
}
