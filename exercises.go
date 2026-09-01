package main

import (
	"fmt"
	"time"
	"errors"
	"sync"
)

/*
	Функция, принимающая 2 слайса.
	Возвращает их пересечение.
	-----------------------------
	a = [0, 1, 4, 9], b = [0, 2, 3, 9]
	> c = [0, 9]
*/
func intersection(a, b []int) []int {
	result := make([]int, 0)

	bMap := make(map[int]bool, len(b))
	for _, el := range b {
		bMap[el] = true
	}

	for _, el := range a {
		if _, ok := bMap[el]; ok {
			result = append(result, el)
		}
	}

	return result
}

/*
	Функция принимает на вход строку.
	Возвращает мапу, где ключи - символы строки, а значения - их частоты.
	-----------------------------------------
	s = "Hello, I love U, Гошка!"
	> map[ :4 !:1 ,:2 H:1 I:1 U:1 e:2 l:3 o:2 v:1 Г:1 а:1 к:1 о:1 ш:1]
*/
func calculateFrequency(s string) map[string]int {
	freq := make(map[string]int, 0)
	for _, r := range s {
		if _, ok := freq[string(r)]; ok {
			freq[string(r)] += 1
		} else {
			freq[string(r)] = 1
		}
	}
	return freq
}

/*
	Функция принимает слайс int.
	Модифицирует его in-place, удаляя чётные числа.
	-----------------------------------
	a = [0, 1, 4, 5, 6, 9]
	> [1, 5, 9, 0, 0, 0]
*/
func filterOdd(nums []int) []int {
	result := nums[:0]

	for _, el := range nums {
		if el % 2 != 0 {
			result = append(result, el)
		}
	}

	for i := len(result); i < len(nums); i++ {
		nums[i] = 0
	}

	return result
}

type Messenger interface {
	SendMessage(msg string) error
}

type EmailService struct{
	IsAvailable bool
}

func NewEmailService(isAvailable bool) *EmailService {
	return &EmailService{IsAvailable: isAvailable}
}

func (e *EmailService) SendMessage(msg string) error {
	if !e.IsAvailable {
		return errors.New("email service is currently unreachable")
	}
	fmt.Printf("[EMAIL]: %s", msg)
	return nil
}

type SMSService struct{
	IsAvailable bool
}

func NewSMSService(isAvailable bool) *SMSService {
	return &SMSService{IsAvailable: isAvailable}
}

func (s *SMSService) SendMessage(msg string) error {
	if !s.IsAvailable {
		return errors.New("sms service is currently unreachable")
	}
	fmt.Printf("[SMS]: %s", msg)
	return nil
}

func Notify(m Messenger, text string) error {
	if err := m.SendMessage(text); err != nil {
		return err
	}
	return nil
}

type Order struct {
	OrderID int `json:"order_id"`
	Status string `json:"status"`
}

/*
	Функция принимает строку.
	Возвращает её перевёрнутую версию
	-----------------------------
	s = "Самокат"
	> "такомаС"
*/
func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

const query = `
	SELECT u.id, u.name, SUM(o.amount) as total_amount
	FROM users u
	JOIN orders o ON u.id = o.user_id
	GROUP BY u.id, u.name
	ORDER BY total_amount DESC
	LIMIT 3;
`

const queryDublicate = `
	SELECT sku
	FROM products
	GROUP BY sku 
	HAVING COUNT(*) > 1;
`

const queryJoin = `
	SELECT c.name
	FROM customers c
	LEFT JOIN orders o ON c.id = o.customer_id
	WHERE o.id IS NULL;
`

const queryProducts = `
	SELECT p.name
	FROM products p
	LEFT JOIN order_items o ON p.id = o.product_id
	WHERE o.id IS NULL;
`

const queryOrders = `
	SELECT user_id, SUM(amount) AS total_amount
	FROM orders
	WHERE status = 'completed'
	GROUP BY user_id
	HAVING total_amount > 5000
	ORDER BY total_amount DESC;
`

const queryEmployees = `
	SELECT name
	FROM employees e1
	INNER JOIN employees e2 ON e1.manager_id = e2.id
	WHERE e1.salary > e2.salary;
`
func Produce(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func workerPool() {
    const (
		numJobs = 15
		numWorkers = 5
	)

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All jobs are completed.")
}

func Merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup{}

	output := func(c <- chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main(){
	ch1 := make(chan int)
	ch := make(chan int)

	go func() {
		ch1 <- 2
		ch1 <- 3
	}
}