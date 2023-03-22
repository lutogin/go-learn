package interfaces

import (
	"fmt"
)

type IVehicle interface {
	move(distance uint)
}

type Car struct{}

type Aircraft struct{}

func (c Car) move(d uint) {
	fmt.Println("Car is moving. Distance: ", d)
}

func (a Aircraft) move(d uint) {
	fmt.Println("Aircraft is flying. Distance: ", d)
}

func drive(v IVehicle, distance uint) {
	v.move(distance)
}

// ---- stream example

type Stream interface {
	read() string
	write(data string)
	close()
}

type Reader interface {
	read() string
}

type Writer interface {
	write(data string)
}

func readFromStream(r Reader) {
	r.read()
}

func writeToStream(s Stream, data string) {
	s.write(data)
}

func streamClose(s Stream) {
	s.close()
}

type File struct {
	data string
}

type Folder struct{}

func (f *File) read() string {
	return f.data
}

func (f *File) write(data string) {
	f.data = data
	fmt.Println("Data was writed in file.")
}

func (f *File) close() {
	fmt.Println("File is close.")
}

func (f *Folder) close() {
	fmt.Println("Folder is close.")
}

func Run() {
	var c1 IVehicle = Car{}
	var a1 IVehicle = Aircraft{}
	c1.move(3)
	a1.move(10)

	drive(c1, 3)
	drive(c1, 10)

	// ---- stream example

	file1 := File{""}
	//	folder1 := Folder{}
	//   writeToStream(folder1)  error, not realized methods write and reads
	writeToStream(&file1, "test writing")
	readFromStream(&file1)
	fmt.Println(file1.read())
	streamClose(&file1)

}
