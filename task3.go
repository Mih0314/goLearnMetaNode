package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// 1
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

func insertData(db *gorm.DB) {
	// 自动创建 students 表（若不存在）
	err := db.AutoMigrate(&Student{})
	if err != nil {
		panic("failed to migrate database")
	}
	db.Create(&Student{ID: 1, Name: "张三", Age: 18, Grade: "A"})
}

func findGtAge18(db *gorm.DB) []Student {
	var stu []Student
	db.Where("age > ?", 18).Find(&stu)
	return stu
}

func updateGrade(db *gorm.DB) {
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
}

func delAgelt15(db *gorm.DB) {
	db.Where("age < ?", 15).Delete(&Student{})
}

// 2
type Accounts struct {
	ID      int
	Balance float64
}
type Transcations struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

func transferMoney(db *gorm.DB, amount float64, fromAccountId int, toAccountId int) error {
	err := db.AutoMigrate(&Accounts{}, &Transcations{})
	if err != nil {
		return err
	}
	db.Create(&Accounts{ID: 1, Balance: 1000.0})
	db.Create(&Accounts{ID: 2, Balance: 1000.0})
	err = db.Transaction(func(tx *gorm.DB) error {
		var A Accounts
		if err := tx.First(&A, fromAccountId).Error; err != nil {
			return err
		}
		if A.Balance < amount {
			return fmt.Errorf("余额不足")
		}
		var B Accounts
		if err := tx.First(&B, toAccountId).Error; err != nil {
			return err
		}
		tx.Model(&Accounts{}).Where("id = ?", fromAccountId).Update("Balance", A.Balance-amount)
		tx.Model(&Accounts{}).Where("id = ?", toAccountId).Update("Balance", B.Balance+amount)
		tx.Create(&Transcations{ID: 1, FromAccountId: fromAccountId, ToAccountId: toAccountId, Amount: amount})
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

type Employee2 struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 查询所有部门为"技术部"的员工，返回结构体切片
func getTechDepartmentEmployees(db *sqlx.DB) ([]Employee2, error) {
	var employees []Employee2
	// 使用 Select 方法直接将查询结果映射到切片
	// SQL 条件：department = "技术部"
	err := db.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// 查询工资最高的员工，返回单个结构体
func getHighestSalaryEmployee(db *sqlx.DB) (Employee2, error) {
	var emp Employee2
	// 使用 Get 方法将单条查询结果映射到结构体
	// SQL 逻辑：按 salary 降序排序，取第一条（即最高工资）
	err := db.Get(&emp, "SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return Employee2{}, err
	}
	return emp, nil
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Posts    []Post
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	User     User
	Comments []Comment
}
type Comment struct {
	gorm.Model
	PostID  uint
	UserID  uint
	User    User
	Post    Post
	Content string
}

func query(db *gorm.DB, userId int) {
	var user User
	targetUserID := uint(1) // 目标用户 ID（可替换为用户名查询）

	// 关键：通过 Preload 多级预加载，避免 N+1 查询
	// Preload("Posts")：预加载用户的所有文章
	// Preload("Posts.Comments")：预加载每篇文章的所有评论
	err := db.Preload("Posts").Preload("Posts.Comments").
		First(&user, "id = ?", targetUserID).Error // 根据用户 ID 查询

	if err != nil {
		log.Fatalf("查询用户文章及评论失败: %v", err)
	}

	// 3. 打印查询结果
	fmt.Printf("用户：%s（ID：%d）的文章及评论如下：\n", user.Username, user.ID)
	for _, post := range user.Posts {
		fmt.Printf("\n文章 ID：%d，标题：%s\n", post.ID, post.Title)
		fmt.Printf("文章内容：%s\n", post.Content)
		fmt.Printf("该文章共有 %d 条评论：\n", len(post.Comments))
		for _, comment := range post.Comments {
			fmt.Printf("  - 评论 ID：%d，内容：%s（发布时间：%v）\n",
				comment.ID, comment.Content, comment.CreatedAt)
		}
	}
}

type PostWithCommentCount struct {
	Post         Post  `gorm:"embedded"`             // 嵌入 Post 结构体，接收文章完整信息
	CommentCount int64 `gorm:"column:comment_count"` // 评论数（别名映射）
}

func queryTopPostsWithMaxComments(db *gorm.DB) {

	// 2. 步骤1：查询最大评论数（先统计所有文章的评论数，找到最大值）
	var maxCommentCount int64
	err := db.Model(&Comment{}).
		Select("COUNT(*) AS comment_count").
		Group("post_id").
		Having("comment_count = MAX(comment_count)"). // 筛选最大评论数
		Limit(1).
		Pluck("comment_count", &maxCommentCount).Error

	if err != nil {
		// 若没有任何评论，maxCommentCount 为 0，需特殊处理
		if err == gorm.ErrRecordNotFound {
			log.Println("暂无任何评论，最大评论数为 0")
			return
		}
		log.Fatalf("查询最大评论数失败: %v", err)
	}

	// 3. 步骤2：查询所有评论数等于最大值的文章（含作者信息）
	var topPosts []PostWithCommentCount
	err = db.Table("posts p").
		// 关联 comments 表，统计每篇文章的评论数
		Joins("LEFT JOIN comments c ON p.id = c.post_id").
		// 关联 users 表，获取文章作者信息（嵌入到 Post.User 中）
		Joins("LEFT JOIN users u ON p.user_id = u.id").
		// 分组统计，筛选评论数等于最大值的文章
		Group("p.id, u.id"). // 需包含关联表的主键（避免分组错误）
		Having("COUNT(c.id) = ?", maxCommentCount).
		// 选择文章字段和评论数（p.* 表示所有文章字段，COUNT(c.id) 为评论数）
		Select("p.*, u.username, u.email, COUNT(c.id) AS comment_count").
		// 将结果映射到自定义结构体（Post 嵌入，CommentCount 单独接收）
		Scan(&topPosts).Error

	if err != nil {
		log.Fatalf("查询评论数最多的文章失败: %v", err)
	}

	// 4. 打印查询结果
	fmt.Printf("评论数最多的文章（共 %d 篇，每篇 %d 条评论）：\n", len(topPosts), maxCommentCount)
	for _, item := range topPosts {
		post := item.Post
		fmt.Printf("\n文章 ID：%d\n", post.ID)
		fmt.Printf("标题：%s\n", post.Title)
		fmt.Printf("内容：%s\n", post.Content)
		fmt.Printf("作者：%s（邮箱：%s）\n", post.User.Username, post.User.Email)
		fmt.Printf("发布时间：%v\n", post.CreatedAt)
	}
}
