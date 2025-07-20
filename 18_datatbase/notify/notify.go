package notify

// type Notifier interface {
// 	// 注册观察者
// 	RegisterObserver(observerID string, readerFunc reader.ReaderNode) error
// 	// 注销观察者
// 	UnregisterObserver(observerID string) error
// 	// 通知变更
// 	NotifyChange(ctx context.Context, change models.TChangeLogs) error
// }

// type MemoryNotifier struct {
// 	// 存储观察者
// 	observers sync.Map // map[string]*ChangeObserver
// 	mu        sync.RWMutex
// }

// func NewMemoryNotifier() *MemoryNotifier {
// 	return &MemoryNotifier{}
// }

// func (n *MemoryNotifier) RegisterObserver(observerID string, readerFunc reader.ReaderNode) error {
// 	n.mu.Lock()
// 	defer n.mu.Unlock()

// 	observer := &reader.ChangeObserver{
// 		ID:         observerID,
// 		ReaderFunc: readerFunc,
// 		NotifyChan: make(chan models.TChangeLogs, 100),
// 	}

// 	n.observers.Store(observerID, observer)
// 	return nil
// }

// func (n *MemoryNotifier) UnregisterObserver(observerID string) error {
// 	n.mu.Lock()
// 	defer n.mu.Unlock()

// 	if observer, ok := n.observers.Load(observerID); ok {
// 		close(observer.(*reader.ChangeObserver).NotifyChan)
// 		n.observers.Delete(observerID)
// 	}
// 	return nil
// }

// func (n *MemoryNotifier) NotifyChange(ctx context.Context, change models.TChangeLogs) error {
// 	n.mu.RLock()
// 	defer n.mu.RUnlock()

// 	var wg sync.WaitGroup
// 	var errs []error

// 	n.observers.Range(func(key, value interface{}) bool {
// 		observer := value.(*reader.ChangeObserver)

// 		// 检查变更是否匹配观察者的查询条件
// 		// if !n.matchesQuery(change, observer.Query) {
// 		// 	return true
// 		// }

// 		wg.Add(1)
// 		go func(obs *reader.ChangeObserver) {
// 			defer wg.Done()
// 			select {
// 			case obs.NotifyChan <- change:
// 				// 通知成功
// 			case <-ctx.Done():
// 				errs = append(errs, ctx.Err())
// 			case <-time.After(1 * time.Second):
// 				errs = append(errs, fmt.Errorf("timeout notifying observer %s", obs.ID))
// 			}
// 		}(observer)

// 		return true
// 	})

// 	wg.Wait()

// 	if len(errs) > 0 {
// 		return fmt.Errorf("notification errors: %v", errs)
// 	}
// 	return nil
// }

// func (n *MemoryNotifier) matchesQuery(change models.TChangeLogs, query reader.ChangeLogQuery) bool {
// 	if query.TableName != "" && change.TableName != query.TableName {
// 		return false
// 	}
// 	if query.TableKey != "" && change.TableKey != query.TableKey {
// 		return false
// 	}
// 	if query.ChangeUser != "" && change.ChangeUser != query.ChangeUser {
// 		return false
// 	}
// 	return true
// }
