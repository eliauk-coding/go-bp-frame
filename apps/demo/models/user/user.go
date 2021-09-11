package user

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"index"`
	Nickname string `gorm:"nickname" json:"-"`
	Password string `gorm:"password" json:"-"`
	Gender   int    `gorm:"gender" json:"gender"`
	Status   int    `gorm:"status" json:"status"`
	UpdateAt int64  `gorm:"autoUpdateTime:milli" json:"utime"`
	CreateAt int64  `gorm:"autoCreateTime:milli" json:"ctime"`
}
