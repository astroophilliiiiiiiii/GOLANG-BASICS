
// relative path -- u give the path by looking at the reference of the current file 
// we r trying to make a multi-file project 
// we need a mechanism to be able to define every single module seperately 
// and we dont want to care about the pathss 


// go mod -- like we define package.json in the nodejs file 
// Go project ka dependency + module manager file.
// It stores:
//      *  project/module name
//      *  Go version
//      *  external packages needed

// import with the help of these my giving the path importing it 
// and then app.___ ese krke use and Real Imag Complex -- capital honaa chahiye 
// small krke use krrooo -- small imag complexx --- ye ni ho skte import 
// capital hona chahiye import krne ke liye 

// go mod init 



// import krnaaa --- diff diff files mein 
//main-folder
  app/  //folder 
    ├── complex.go
├── go.mod
├── lib.go
└── main.go
// diff packages  m h jaise    main-package( main.go )  app-package( complex.go )  -- toh import krke use + Capital letter 
// capital letters only matter in cross packages not in same 

// same packages  m h jaise    main.go    lib.go 		
// humne import krke upr vaala app package chlaa diyaa abhi go run main.go krke ye chlaa diyaa 
// but humne kahinpe bhi lib.go toh chlaayi nhi 
// soo we need to compile all the files of this folder --- go run *.go  -- current folder files also run 


// compile everything and attach to the final executable !! 
