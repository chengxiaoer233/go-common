/**
* @Description：golang 无限制的channel
* @Author: cdx
* @Date: 2022/12/4 2:50 下午
 */

package _chan

type UnboundedChan struct {
	In     chan<- interface{} // channel for write
	Out    <-chan interface{} // channel for read
	buffer []interface{}      // buffer
}

// Len returns len of Out plus len of buffer.
func (c UnboundedChan) Len() int {
	return len(c.buffer) + len(c.Out)
}
// BufLen returns len of the buffer.
func (c UnboundedChan) BufLen() int {
	return len(c.buffer)
}

func NewUnboundedChan(initCapacity int) UnboundedChan {

	// Create a chan type with infinite existence and three fields
	in := make(chan interface{}, initCapacity)
	out := make(chan interface{}, initCapacity)
	ch := UnboundedChan{In: in, Out: out, buffer: make([]interface{}, 0, initCapacity)}

	// Through a goroutine, continuously read data from in and put it into out or buffer
	go func() {
		defer close(out) // in is closed, and out is also closed after the data is read

	loop:
		for {
			val, ok := <-in
			if !ok { // If in has been closed, exit the loop
				break loop
			}

			// try to put the data read from in into out
			select {
			case out <- val:
				continue
			// If the input is successful, it means that the out is not full,
			// and there is no additional data to be processed in the buffer,
			// so return to the loop to start
			default:
			}

			// If out is full, you need to put the data into the cache
			ch.buffer = append(ch.buffer, val)

			// Handle the cache, keep trying to put the data in the cache into out,
			// until there is no data in the cache,

			// In order to avoid blocking the in channel, you need to try to read data from in,
			// because out is full at this time, so put the data directly into the cache
			for len(ch.buffer) > 0 {
				select {
				case val, ok := <-in: // read data from in, put it into the cache, if in is closed, exit the loop
					if !ok {
						break loop
					}
					ch.buffer = append(ch.buffer, val)

				case out <- ch.buffer[0]: // Put the oldest data in the cache into out and remove the first element
					ch.buffer = ch.buffer[1:]
					if len(ch.buffer) == 0 { // Avoid memory leaks. If the cache is processed, restore to the original state
						ch.buffer = make([]interface{}, 0, initCapacity)
					}
				}
			}
		}

		// in is closed, after exiting the loop, there may still be unprocessed data in the buffer,
		// which needs to be stuffed into out

		// This logic is called "drain".

		// After this logic is processed, the out can be closed.
		for len(ch.buffer) > 0 {
			out <- ch.buffer[0]
			ch.buffer = ch.buffer[1:]
		}

		if len(ch.buffer) == 0 { // Avoid memory leaks. If the cache is processed, restore to the original state
			ch.buffer = make([]interface{}, 0, initCapacity)
		}

	}()
	return ch
}