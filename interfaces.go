// defining an interface 
type sellableProduct interface{
	// function -- return type -- string 
	buy() 
	getDiscount() int 
}


// to implement the interface we dont have a keyword to write with the struct 
// like in ts or cpp class implements interface -- something like this we used to do 

// to implement the interface -- 
// 		struct has to implement all the methods mentioned in the interface 
// if it doesnt then it will not be considered to implementing that interface 

// it accepts the object that implements the interface( sellableProduct )
func check_discount_and_buy( p sellableProduct ){
	discount := p.getDiscount() 
	if( discount>30 ){
		fmt.Println( "Discount is good !! buying the product !! "); 
		p.buy()
		return 
	}else{
		fmt.Println("Discount is not good !! not buying the product !! "); 
		return 
	}
}

type Product struct{
	name string , 
	price int 
}

func ( p *Product ) buy(){
	fmt.Println("buying the product:- " , p.name , "-- " , p.price )
}

func ( p *Product ) getDiscount() int{
	discount = 40 ; 
	fmt.Println("Discount for product ", p.name , " is- ", discount ) ; 
	return discount ;
}

func main(){
	p := Product{
		name: "Iphone", 
		price : 150000 , 
	}
	check_discount_and_buy( p ); // accpeted -- as it belongs to the class -- that accepts all methods
	

}
