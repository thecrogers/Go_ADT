The Purpose of this project it so make a simple mostly recursive Go Lib that has several ADT that my be used as individual packages, or as one large package.

The goal implementation will look something like one of the following



`import	github.com/thecrogers/Go_ADT/

   type Some struct {
     someDate 	   string
     someOtherDate int
   }

   myList :=  ADT.SingleList{ *Some }

   var myData Some{ "new data", 5 }
   myList.Add(myData)

   result :=  myList.Find(Some.someOtherData, 5)
`

Or

`
import github.com/thecrogers/Go_ADT/ADT/SingleLinkedList

type data struct {

}

SLL := SingleLinkedList.newAdt()
dataVar data{

}
SLL.AddStart(id, dataVar )
`


I have also included a go test file for each type that I will use to verify each method of the ADT functions as expected.
