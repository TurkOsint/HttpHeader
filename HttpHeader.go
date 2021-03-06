package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		usage(os.Args[0])
	}

	url := os.Args[1]


	response, err := http.Head(url)
	if err != nil {
		fmt.Println(`
###      ###  _____         _    ___     _     _   
#	 o # |_   _|  _ _ _| |__/ _ \ __(_)_ _| |_ 
	       | || || | '_| / / (_) (_-< | ' \  _|
               |_| \_,_|_| |_\_\ \__//__/_|_||_\__|
#	   # Proje Adı: https://github.com/TurkOsint
###      ### Kodlayan : https://github.com/EyupErgin`)	
		log.Fatal("[TurkOsint] URL Alınırken Hata Oluştu Lütfen Doğru URL'yi Yazınız.")
	}

	fmt.Println(`
###      ###  _____         _    ___     _     _   
#	 o # |_   _|  _ _ _| |__/ _ \ __(_)_ _| |_ 
	       | || || | '_| / / (_) (_-< | ' \  _|
               |_| \_,_|_| |_\_\ \__//__/_|_||_\__|
#	   # Proje Adı: https://github.com/TurkOsint
###      ### Kodlayan : https://github.com/EyupErgin`)	
	for key, value := range response.Header {
		fmt.Printf("%s: %s\n", key, value[0])
	}
}

func usage(name string) {
	fmt.Println(`
###      ###  _____         _    ___     _     _   
#	 o # |_   _|  _ _ _| |__/ _ \ __(_)_ _| |_ 
	       | || || | '_| / / (_) (_-< | ' \  _|
               |_| \_,_|_| |_\_\ \__//__/_|_||_\__|
#	   # Proje Adı: https://github.com/TurkOsint
###      ### Kodlayan : https://github.com/EyupErgin`)
	fmt.Printf("[TurkOsint] Lütfen Sadece URL girin.\n")
	fmt.Printf("[TurkOsint] Kullanım : ./HttpHeader https:example.com\n")
	os.Exit(1)
}
