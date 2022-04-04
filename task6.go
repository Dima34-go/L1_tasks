package main

//остановка по сигналу
//func main()  {
//	quit := make(chan bool)
//	go func(){
//		for {
//			select {
//			case <- quit:
//				return
//			default:
//				time.Sleep(100*time.Millisecond)
//				fmt.Println("gorutine working")
//			}
//		}
//	}()
//	b:=make([]byte,1)
//	os.Stdin.Read(b)
//	quit <- true
//}

//остановка после n  операций
//func main() {
//	wg:=&sync.WaitGroup{}
//	n:=10
//	wg.Add(n)
//	go func() {
//		for i:=0;i<n;i++{
//			fmt.Println(i,"iter")
//			wg.Done()
//		}
//	}()
//	wg.Wait()
//}

//остановка при закрытии контекста
//func main()  {
//	ctx,cl:=context.WithCancel(context.Background())
//	go func(ctx context.Context) {
//		for true {
//			select {
//			case <-ctx.Done():
//				return
//			default:
//				time.Sleep(300*time.Millisecond)
//				fmt.Println("gorutine work")
//			}
//		}
//	}(ctx)
//	b:=make([]byte,1)
//	os.Stdin.Read(b)
//	cl()
//	os.Stdin.Read(b)
//}