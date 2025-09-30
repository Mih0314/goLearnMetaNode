package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func main() {
	//task1
	//singleNumber([]int{2, 2, 1})
	//println(isValid("()[]{}"))
	//fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(plusOne([]int{1, 2, 3}))
	//fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	//task2
	//指针
	//num := 10
	//recePointerAdd(&num)
	//fmt.Println(num)

	//a, b, c := 1, 2, 3
	//nums := []*int{&a, &b, &c}
	//recePointerMul(nums)
	//fmt.Println(a, b, c)

	//go print1_10()
	//go print2_10()
	//time.Sleep(time.Second * 1)

	//r := Rectangle{}
	//r.Area()
	//r.Perimeter()
	//c := Circle{}
	//c.Area()
	//c.Perimeter()

	//p := Person{"12", 12}
	//e := Employee{p, 54}
	//fmt.Println(e)

	//ch := make(chan int)
	//go sendNum(ch,10)
	//go receNum(ch,10)
	//time.Sleep(time.Second * 2)

	//ch2 := make(chan int, 100)
	//go sendNum(ch2, 100)
	//go receNum(ch2, 100)
	//time.Sleep(time.Second * 2)

	//var mu = sync.Mutex{}
	//count :=int32(0)
	//s := safeCounter{mu, count}
	//for i := 0; i < 10; i++ {
	//	go s.inc1000()
	//}
	//time.Sleep(time.Second * 5)
	//fmt.Println(s.count)

	//var wg sync.WaitGroup
	//a := atomicCounter{0}
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go a.atomicInc1000(&wg)
	//}
	//wg.Wait()
	//fmt.Println(a.count)

	//task3
	dsn := "root:root@tcp(127.0.0.1:3306)/grom?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}

	//insertData(db)

	//str := findGtAge18(db)
	//fmt.Println(str)

	//updateGrade(db)

	//delAgelt15(db)

	//err = transferMoney(db, 100.0, 1, 2)
	//if err != nil {
	//	panic(err)
	//}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close() // 程序退出时关闭连接

	// 2. 查询所有部门为"技术部"的员工
	//techEmployees, err := getTechDepartmentEmployees(db)
	//if err != nil {
	//	log.Printf("查询技术部员工失败: %v", err)
	//} else {
	//	fmt.Println("----- 技术部员工列表 -----")
	//	for _, emp := range techEmployees {
	//		fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
	//			emp.ID, emp.Name, emp.Department, emp.Salary)
	//	}
	//}

	// 3. 查询工资最高的员工
	//topSalaryEmp, err := getHighestSalaryEmployee(db)
	//if err != nil {
	//	log.Printf("查询最高工资员工失败: %v", err)
	//} else {
	//	fmt.Println("\n----- 工资最高的员工 -----")
	//	fmt.Printf("ID: %d, 姓名: %s, 部门: %s, 工资: %.2f\n",
	//		topSalaryEmp.ID, topSalaryEmp.Name, topSalaryEmp.Department, topSalaryEmp.Salary)
	//}
}
