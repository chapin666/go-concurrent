package v1

import (
	"errors"
	"io"
	"os"
	"sync"
)

type Data []byte

type DataFile interface {

	// Read data from file
	Read() (rsn int64, d Data, err error)

	// Write data from file
	Write(d Data) (wsn int64, err error)

	// read data block NO
	RSN() int64

	// write data block NO
	WSN() int64

	// data length
	DataLen() uint32

	// close file
	Close() error
}

type myDataFile struct {
	f       *os.File     //文件
	fmutex  sync.RWMutex //文件读写锁
	woffset int64        //写操作偏移
	roffset int64        //读操作偏移
	wmutex  sync.Mutex   //写操作互斥锁
	rmutex  sync.Mutex   //读操作互斥锁
	dataLen uint32       //数据块长度
}

// 新建数据文件实例
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}

// 读操作
/*
func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取数据块
	rsn = offset / int64(df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()
	// 缓存数组
	bytes := make([]byte, df.dataLen)
	_, err = df.f.ReadAt(bytes, offset)
	if err != nil {
		return
	}
	d = bytes
	return
}*/

// 读操作
func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取数据块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)

	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		df.fmutex.RUnlock()
		return
	}
}

// Write file
func (df *myDataFile) Write(d Data) (wsn int64, err error) {

	// 更新offset
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	// 写入数据块
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return
}

// RSN
func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

// RSN
func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

// DataLen
func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

// Close
func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}
