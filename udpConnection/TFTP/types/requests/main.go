package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strings"
)

const(
	DatagramSize = 516 // maximum supported datagram DatagramSize
	BlockSize = DatagramSize - 4 // the DatagramSize minus a 4-byte header
)

type OpCode uint16
const(
	OpRRQ OpCode = iota + 1
	_    //no WRQ support 
	OpData 
	OpAck
	OpErr
)

type ErrCode uint16

const(
	ErrUnknown ErrCode = iota + 500
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID 
	ErrFileExists
	ErrNoUser
)

type ReadReq struct{
	FileName string
	Mode string
}

func (q ReadReq) marshalBinary()([]byte, error){
	mode := "octet"
	if q.Mode != ""{
		mode = q.Mode
	}

	// operaton code + filename + 0 byte + mode + 0 byte
	cap := 2 + 2 + len(q.FileName) + 1 + len(q.Mode) + 1

	b := new(bytes.Buffer)
	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OpRRQ) //write operation code
	if err != nil{
		return nil, err
	}

	_, err = b.WriteString(q.FileName)//write FileName
	if err != nil{
		return nil, err
	}

	err = b.WriteByte(0)//write 0 byte
	if err != nil{
		return nil, err
	}
	_,err = b.WriteString(mode)//write Mode
	if err != nil{
		return nil, err
	}
err = b.WriteByte(0)// write 0 byte
if err != nil{
	return nil, err
}

	return b.Bytes(), nil
}

func (q *ReadReq) unmarshalBinary(p []byte)error{
	r := bytes.NewBuffer(p)
	var code OpCode

	err := binary.Read(r, binary.BigEndian, &code)//read operation code
	if err != nil{
		return err
	}

	if code != OpRRQ{
		return errors.New("invalid RRQ")
	}

	q.FileName, err = r.ReadString(0) // read FileName
	if err != nil{
		return errors.New("invalid RRQ")
	}

	q.FileName = strings.TrimRight(q.FileName, "\x00")//remove the 0-byte
	if len(q.FileName) == 0{
		return errors.New("invalid RRQ")
	}

	q.Mode, err = r.ReadString(0) // read mode 
	if err != nil{
		return errors.New("invalid RRQ")
	}

	q.Mode = strings.TrimRight(q.Mode, "\x00")// remove the 0-byte 
	if len(q.Mode) == 0{
		return errors.New("invalid RRQ")
	}

	actual := strings.ToLower(q.Mode)//enforce octet Mode
	if actual != "octet"{
		return errors.New("only binary transfers supported")
	}
	return nil
}