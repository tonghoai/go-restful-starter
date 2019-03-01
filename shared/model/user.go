package shared

type User struct {
	ID         int    `gorm:"type:int(11) AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	Username   string `gorm:"type:varchar(50)" json:"username" validate:"required"`
	Password   string `gorm:"type:varchar(100)" json:"password" validate:"required"`
	Permission int    `gorm:"type:tinyint(1) DEFAULT 1" json:"permission"`
	Name       string `gorm:"type:varchar(100)" json:"name"`
	Age        int    `gorm:"type:int(11)" json:"age"`
	Address    string `gorm:"type:varchar(200)" json:"address"`
}

type UserService struct {
}
