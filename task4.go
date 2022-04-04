package main


//func main(){
//	N:=10
//	ch:=make(chan int)
//	go func() {
//		for ; ; {
//			ch<-rand.Int()
//		}
//	}()
//	wg:=sync.WaitGroup{}
//	for i:=0;i < N;i++{
//		k:=i
//		go func() {
//			wg.Add(1)
//			for val:=range ch{
//				fmt.Println("отработал воркер",k,val)
//			}
//			wg.Done()
//		}()
//	}
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
//	<-quit
//	// Закрыть канал ch  и ждать пока все воркеры закончат работа
//	wg.Wait()
//}