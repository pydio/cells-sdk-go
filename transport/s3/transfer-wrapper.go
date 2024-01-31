package s3

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/dustin/go-humanize"
)

// TODO rather use a percentage for the logs
// TODO investigate to insure that we cannot avoid to handle a buffer at this level

// bufferSizeInKB is the human friendly size to define
// an optimum buffer size to wrap effective file transfers.
const bufferSizeInKB = 1024

// bufferSize caches the size in bytes that is effectively used.
const bufferSize = bufferSizeInKB * 1024

// CallbackTransferProvider provide an access to the real transfer process when using the AWS S3 upload manager,
// e.g to get more logs
type CallbackTransferProvider struct {
	targetName string
	fileSize   int64
	partSize   int64
}

// NewCallbackTransferProvider creates a new instance.
func NewCallbackTransferProvider(fileName string, fileSize, partSize int64) *CallbackTransferProvider {
	return &CallbackTransferProvider{
		targetName: fileName,
		fileSize:   fileSize,
		partSize:   partSize,
	}
}

// GetWriteTo is called by the AWS SDK for each part to create a new instance of the wrapper.
func (m *CallbackTransferProvider) GetWriteTo(seeker io.ReadSeeker) (r manager.ReadSeekerWriteTo, cleanup func()) {
	fmt.Printf("... Launching a new part upload of %s for %s\n", humanize.Bytes(uint64(m.partSize)), m.targetName)
	r = newCustomWrapper(seeker)
	cleanup = func() {
		// TODO Perform necessary cleanup
		fmt.Printf("... Cleaning up after transfering part for %s\n", m.targetName)
	}
	return r, cleanup
}

// customWrapper implements ReadSeekerWriteToProvider interface.
type customWrapper struct {
	buffer   [bufferSize]byte
	dataSrc  io.ReadSeeker
	readPos  int64
	writePos int64
	counter  int
}

// newCustomWrapper creates a new instance.
func newCustomWrapper(src io.ReadSeeker) *customWrapper {
	return &customWrapper{dataSrc: src, counter: 0}
}

// Read implements io.Reader interface.
func (p *customWrapper) Read(b []byte) (n int, err error) {
	n, err = p.dataSrc.Read(b)
	p.readPos += int64(n)
	if p.readPos >= bufferSize {
		if p.counter%5 == 0 {
			fmt.Println(humanize.Bytes(uint64(bufferSize*5)), "of data transfered")
		}
		p.readPos -= bufferSize
		p.counter++
	}
	return
}

// Seek implements io.Seeker interface.
func (p *customWrapper) Seek(offset int64, whence int) (int64, error) {
	return p.dataSrc.Seek(offset, whence)
}

// Seek implements io.WriterTo interface.
func (p *customWrapper) WriteTo(w io.Writer) (n int64, err error) {

	for {
		readBytes, readErr := p.dataSrc.Read(p.buffer[:])
		if readBytes > 0 {
			writeBytes, writeErr := w.Write(p.buffer[:readBytes])
			n += int64(writeBytes)
			if writeErr != nil {
				return n, writeErr
			}
			p.writePos += int64(writeBytes)
			if p.writePos >= bufferSize {
				p.writePos -= bufferSize
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return n, readErr
		}
	}
	return n, nil
}
