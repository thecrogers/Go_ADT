The Purpose of this project it so make a simple recursive Go Lib that has most ADT that my be used.

The goal implementation will look something like the following


``` import {
	github.com/thecrogers/Go_ADT ADT

    }
    
   type Some struct {
     someDate 	   string
     someOtherDate int
   }
	
   myList :=  ADT.SingleList{ *Some }

   var myData Some{ "new data", 5 }
   myList.Add(myData)

   result :=  myList.Find(Some.someOtherData, 5)
```

