package notifier

import (
	"context"
	"log"
	"time"
)

// Notifier: Quản lý queue và strategies
type Notifier struct {
	strategies []NotificationStrategy
	queue      chan NotificationMessage
}

// NewNotifier: Khởi tạo với list strategies và buffer size
func NewNotifier(strategies []NotificationStrategy, bufferSize int) *Notifier {
	n := &Notifier{
		strategies: strategies,
		queue:      make(chan NotificationMessage, bufferSize),
	}
	go n.startWorker()
	return n
}

// startWorker: Goroutine xử lý queue background
func (n *Notifier) startWorker() {
	for msg := range n.queue {
		n.sendWithAllStrategies(context.Background(), msg)
	}
}

// sendWithAllStrategies: Gửi qua tất cả strategies (email + sms nếu config)
func (n *Notifier) sendWithAllStrategies(ctx context.Context, msg NotificationMessage) {
	for _, strategy := range n.strategies {
		if err := n.sendWithRetry(ctx, msg, strategy); err != nil {
			log.Printf("[%s] Failed to send %s to %s after retries: %v", strategy.Name(), msg.Type, msg.To, err)
			// Có thể thêm metrics Prometheus: counter.Inc("notification_failed")
		} else {
			log.Printf("[%s] Sent %s successfully to %s", strategy.Name(), msg.Type, msg.To)
			// metrics: counter.Inc("notification_sent")
		}
	}
}

// sendWithRetry: Retry exponential backoff (1s, 2s, 4s,...)
func (n *Notifier) sendWithRetry(ctx context.Context, msg NotificationMessage, strategy NotificationStrategy) error {
	maxRetries := 3
	backoff := time.Second

	for i := 0; i <= maxRetries; i++ {
		err := strategy.Send(ctx, msg)
		if err == nil {
			return nil
		}
		log.Printf("[%s] Failed to send %s (attempt %d): %v", strategy.Name(), msg.Type, i+1, err)
		if i == maxRetries {
			return err
		}
		time.Sleep(backoff)
		backoff *= 2
	}
	return nil
}

// Send: Public method để đẩy message vào queue (non-blocking nếu full)
func (n *Notifier) Send(msg NotificationMessage) {
	select {
	case n.queue <- msg:
	default:
		log.Printf("Notification queue full, dropping %s for %s", msg.Type, msg.To)
	}
}
