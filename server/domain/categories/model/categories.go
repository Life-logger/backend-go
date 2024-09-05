package categories

type Category interface {
	Id() int
	UserEmail() string
	Title() string
	Color() string
}

type categoryImpl struct {
	id        int
	userEmail string
	title     string
	color     string
}

func NewCategory(userEmail, title, color string) Category {
	i := new(categoryImpl)
	i.userEmail = userEmail
	i.title = title
	i.color = color
	return i
}

func (c categoryImpl) Id() int {
	return c.id
}

func (c categoryImpl) UserEmail() string {
	return c.userEmail
}

func (c categoryImpl) Title() string {
	return c.title
}

func (c categoryImpl) Color() string {
	return c.color
}
