package actioninfo
import(
	"log"
	
	
		
)

type DataParser interface {
	// TODO: добавить методы

	Parse(string) (error)  
	
	ActionInfo() (string, error)

}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
for q := range dataset{
er:= dp.Parse(dataset[q])
	if er !=nil{
	log.Println("Ошибка парсинга для элемента", er)
			continue
	}
 dp.ActionInfo()
if er !=nil{
	log.Println("Ошибка парсинга для элемента", er)
			}
return 
}
}
