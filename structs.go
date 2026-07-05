// in go there is no class -- there r structs 
// whenever we pass somehing usually the object tha passes is the copy not the actual object 

// struct defination 
type product struct{
	name string
	price int
	company string
}

// func fun( copyOfp Product ){
// 	// this p is the copy of the Product p 
// }

// func fun( p *Product ){
// 	// this is actual p   --- pointer passed that points to the actual p product 
// }

// constructor function for the above struct  -- pass by value 
func newProduct( name string , price int , company string ) : Product{
	copyOfpp := Product{
		name : name , 
		price : price , 
		company : company 
	}
	return copyOfp ; // but its returning the copy of the product p ( not the actual )
}

// public -- if Name ( capital N )
// private -- name ( small n )

// constructor function for the above struct  -- pass by reference 
func newProduct2( name string , price int , company string ) : *Product{
	p := Product{
		name : name , // private 
		Salary : salary , // public -------------------------------------------
		price : price , 
		company : company 
	}
	return &p ; // address of p returning -- returning the pointer 
}

// member functions 
// func ( object & strcut ) display return-type ( void return type genrally not written )
func ( p *Product ) display(){
	 fmt.Println( "name:- ", p.name , " price:- ", p.price , " company:- ", p.company ) ; 
}

func main(){

	// object creation 
	p := Product{
		name : "Iphone", 
		price : 1000 , 
		company : "Apple Inc",
	}

	fmt.Println( p.name ) 

	// constructor where   p2 is the pointer 
	p2 := newProduct2( "Iphone" , 1000 , "Apple" ) 
	fmt.Println( p2.name ) // pointer hone ke bawjood . krke access ,  usually * se pointer value accessed 

	
}