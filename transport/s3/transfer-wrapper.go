package s3

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

// TODO investigate to insure that we cannot avoid to handle a buffer at this level

// bufferSizeInKB is the human friendly size to define
// an optimum buffer size to wrap effective file transfers.
const bufferSizeInKB = 1024

// bufferSize caches the size in bytes that is effectively used.
const bufferSize = bufferSizeInKB * 1024

// CallbackTransferProvider provides access to the real transfer process when using the AWS S3 upload manager,
// e.g. to get more logs
type CallbackTransferProvider struct {
	targetName   string
	fileSize     int64
	partSize     int64
	partNumber   int
	partLaunched int
}

// NewCallbackTransferProvider creates a new instance.
func NewCallbackTransferProvider(fileName string, fileSize, partSize int64) *CallbackTransferProvider {
	return &CallbackTransferProvider{
		targetName:   fileName,
		fileSize:     fileSize,
		partSize:     partSize,
		partNumber:   int(float32(fileSize)/float32(partSize)) + 1,
		partLaunched: 0,
	}
}

// GetWriteTo is called by the AWS SDK for each part to create a new instance of the wrapper.
func (m *CallbackTransferProvider) GetWriteTo(seeker io.ReadSeeker) (r manager.ReadSeekerWriteTo, cleanup func()) {

	// Note that it does not need to be wrapped in a mutex because GetWriteTo is only called from the main thread from the AWS SDK.
	m.partLaunched++
	currIndex := m.partLaunched

	// fmt.Printf("... Launching upload of part %d of %s for %s\n", m.partLaunched, humanize.Bytes(uint64(m.partSize)), m.targetName)
	r = newCustomWrapper(seeker, m.partSize, currIndex)
	cleanup = func() {
		// TODO Perform necessary cleanup
		fmt.Printf("Part #%d: 100%% transferred\n", currIndex)
		fmt.Printf("... Cleaning up after transferring part %d/%d for %s\n", currIndex, m.partNumber, m.targetName)
	}
	return r, cleanup
}

// customWrapper implements the ReadSeekerWriteToProvider interface.
type customWrapper struct {
	partId   int
	buffer   [bufferSize]byte
	dataSrc  io.ReadSeeker
	readPos  int64
	writePos int64
	counter  int
	partSize int64
}

// newCustomWrapper creates a new instance.
func newCustomWrapper(src io.ReadSeeker, partSize int64, partId int) *customWrapper {
	return &customWrapper{dataSrc: src, counter: 0, partId: partId, partSize: partSize}
}

// Read implements io.Reader interface.
func (p *customWrapper) Read(b []byte) (n int, err error) {
	n, err = p.dataSrc.Read(b)
	p.readPos += int64(n)
	if p.readPos >= bufferSize {
		if p.counter%10 == 0 {
			ratio := float32(p.counter*bufferSize) / float32(p.partSize) * 100
			fmt.Printf("Part #%d: %d%% transferred\n", p.partId, int(ratio))
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

// WriteTo implements io.WriterTo interface.
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
