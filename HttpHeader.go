package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
//Argümanlar.
func main() {

	if len(os.Args) != 2 {
		usage(os.Args[0])
	}

	url := os.Args[1]

//Response 
	response, err := http.Head(url)
	
	//Olası Hata : Yanlış Url Girilmesi
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
	//Başarılı : Geri dönen istekleri yaz. 
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

	//Olası Hata : 1'den fazla string girilmesi
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
