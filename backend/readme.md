**开发环境热启动**

```
# 安装
go get github.com/silenceper/gowatch

# 执行
gowatch
```

**去掉未注释警告**

项目根目录下创建文件夹`.vscode`，新建文件`settings.json`，内容如下：

```
{
    "go.lintFlags": ["--disable=all", "--enable=errcheck"]
}
```


/* type UserMenu sruct {
	ID  string `json:"id"`
	Url string `json:"url"`
} */




/*
func (u Login) TableName() string {
	return "user"
}

func (u User) TableName() string {
	turn "user"


unc (u *LoginUr )UserByAccountAndPassword() {
	:= global.MYSQL.Where("account  ? and password = ?", u.Account, u.Password).First(&u)



func (u *User) UserInfo( {
	db := global.MYSQL.Where("i= ?", u.ID).First(&u)
	rmissionList : []UserPermission{}
b.Related(&permissionList)

menuList := []Usenu{}
	d.Related(&menuList)
}

unc (u User) List([]User, error) {
	t := []User{}
f err := globl.YSQL.Find(&list).Error; err != nil {
	eturn list, err

	return list, nil


un(u User) Pag(page *tools.Pagination) ([]User, error) {
	list := []User{}
	 := global.MYSQL.Find(&list).Error
 err != nil {
	eturn list, err
}
	turn list,nil


c (u *User) Create() (string, error) {
	turn "", nil


unc (u *User) Update() {



unc (u *User)elete() {
	 global.MYSQL.Where("id = ?" u.ID).Delete(User{})
}

func (u *User) IsExist() bool {
	return false
}
*/
