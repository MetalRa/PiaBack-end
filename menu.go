op := 1
for{
	
	menu := 
`
	Bienvenido al sistema de surtido de material
	[ 1 ] Ver inventario
	[ 2 ] Entrada de material
	[ 3 ] Pedir material para producir un producto terminado
	[ 0 ] Salir
	Presiona un numero
`
	fmt.Print(menu)

	var eleccion int
	fmt.Scanln(&eleccion)

	switch eleccion{
		case 1:
			fmt.Println("")
		case 2:
			fmt.Println("")
		default:
			fmt.Println("")
		case 0:
			eleccion == op

	if op == 0{
		break
	}
}
fmt.Println("Fin del programa")