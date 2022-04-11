package Animal

type animal struct{
	age,name,raze string
	
}
func animalDescription(a animal) string{
return ("The animal "+a.name+" is a "+a.raze+" and it is "+a.age+" years old")
}