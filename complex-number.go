// Create a Complex Number class without using any built-in complex number data type.

// Your class should support the following operations:

// Create a new complex number with real and imaginary parts.
// Add two complex numbers.
// Subtract two complex numbers.
// Multiply two complex numbers.
// Provide a method to display a complex number in proper format.

// struct -- defination 

// like this is Complex naa -- toh if i use this -- its kindoff global type !! 
type Complex struct{
	real int 
	imag int 
}

// constructor -- isme direct likhte hain name , imag 
func newcom( real int , imag int ) *Complex{
     p := Complex{
		real : real,  // p.real : real   XXX
		imag : imag   // p.imag : imag   XXX
	 }
	 return &p; 
}

// member functions 
// p1 jispe call kiyyaa gyee and new p2 give to function to add to it 
func ( p1 *Complex ) add(p2 *Complex ) *Complex{
	p := Complex{ // object defination hai : aayegaa na
		real : p1.real + p2.real , 
		imag : p1.imag + p2.imag 
	}
	return &p ; 
}

// objects as parameters also pass by reference -- avoid creating too many copies 
func ( p1 *Complex ) subtract(p2 *Complex) *Complex{
	p := Complex{
		real : p1.real - p2.real , 
		imag : p1.imag - p2.imag 
	}
	return &p ; 
}

func ( p1 *Complex ) multiply(p2 *Complex) *Complex{
	p := Complex{
		real : p1.real*p2.real - p1.imag*p2.imag , 
		imag : p1.real*p2.imag + p1.imag*p2.real 
	}
	return &p ; 
}

func ( p *Complex ) display(){
	fmt.Println( p.real , "+" , p.imag , "i" ); 
}

func main(){
	p1 := Complex{ 2 , 3 }
	p2 := Complex{ 2 , 2 }

	p3 := p1.add( p2 ) ;   // first time -- defining -- give := 
	p3.display() // 4+5i

	p3 = p1.subtract( p2 ) ; // redefining -- only = give !! 
	p3.display() // 0+1i

	p3 = p1.multiply( p2 ) ; 
	p3.display() // 4+6i
}
