package utils

type NsqWriter struct {
	writer *nsq.Writer
	topic  string
}

// Implement Writer interface
// TODO: Should handle errors
func (nsqw *NsqWriter) Write(body []byte) (int, error) {
	nsqw.writer.Publish(nsqw.topic, body)
	return len(body), nil
}

// Returns a new nsq writer connected to passed host and
// publishing to passed topic
func NewNsqWriter(host, topic string) *NsqWriter {
	return &NsqWriter{
		nsq.NewWriter(host),
		topic,
	}
}
