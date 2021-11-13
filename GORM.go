	db, err := gorm.Open(sqlite.Open("PCB.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Fallo la conexion a la base de datos")
	}
	
	// Migrate the schema
	db.AutoMigrate(&ProductoTerminado{})
	db.AutoMigrate(&Soldadura{})
	db.AutoMigrate(&Componente{})

	// Create
	db.Create(&ProductoTerminado{Modelo: "Honda", NumeroParte: 10024, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Honda", NumeroParte: 9721, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Ford", NumeroParte: 2340, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Ford", NumeroParte: 2350, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Subaru", NumeroParte: 10362, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Subaru", NumeroParte: 10372, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Fiat", NumeroParte: 6720, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Fiat", NumeroParte: 6320, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Kia", NumeroParte: 8120, Stock: 0})
	db.Create(&ProductoTerminado{Modelo: "Kia", NumeroParte: 5740, Stock: 0})

	db.Create(&Soldadura{Tipo: "Libre de plomo", Cantidad: 500})
	db.Create(&Soldadura{Tipo: "Plomo", Cantidad: 500})

	db.Create(&Componente{TipoComponente: "Resistencia", NumeroParte: 31255, Cantidad: 5000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Resistencia", NumeroParte: 32655, Cantidad: 5000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Resistencia", NumeroParte: 31255, Cantidad: 5000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Resistencia", NumeroParte: 3330, Cantidad: 1000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Resistencia", NumeroParte: 7640, Cantidad: 5000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Capacitor", NumeroParte: 2280, Cantidad: 500, Stock: 10})
	db.Create(&Componente{TipoComponente: "Capacitor", NumeroParte: 2270, Cantidad: 500, Stock: 10})
	db.Create(&Componente{TipoComponente: "Capacitor", NumeroParte: 1220, Cantidad: 800, Stock: 10})
	db.Create(&Componente{TipoComponente: "Transistor", NumeroParte: 3280, Cantidad: 1000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Transistor", NumeroParte: 3290, Cantidad: 1000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Transistor", NumeroParte: 7640, Cantidad: 500, Stock: 10})
	db.Create(&Componente{TipoComponente: "Diodo", NumeroParte: 1470, Cantidad: 700, Stock: 10})
	db.Create(&Componente{TipoComponente: "Diodo", NumeroParte: 1520, Cantidad: 750, Stock: 10})
	db.Create(&Componente{TipoComponente: "Diodo", NumeroParte: 1880, Cantidad: 700, Stock: 10})
	db.Create(&Componente{TipoComponente: "Led", NumeroParte: 9200, Cantidad: 2000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Led", NumeroParte: 9210, Cantidad: 2000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Led", NumeroParte: 1160, Cantidad: 3000, Stock: 10})
	db.Create(&Componente{TipoComponente: "Inductor", NumeroParte: 8800, Cantidad: 500, Stock: 10})

	// Read
	var product ProductoTerminado
	var solder Soldadura
	var component Componente

	db.First(&product, 1) // find product with integer primary key
	fmt.Println(product.ID, product.Model, product.NumeroParte, product.Stock)
	db.First(&product, 2) // find product with integer primary key
	fmt.Println(product.ID, product.Model, product.NumeroParte, product.Stock)
	db.First(&product, 3) // find product with integer primary key
	fmt.Println(product.ID, product.Model, product.NumeroParte, product.Stock)*/

	db.First(&solder, 1) // find product with integer primary key
	fmt.Println(solder.Tipo, solder.Cantidad)

	db.First(&component, 10) // find product with integer primary key
	fmt.Println(component.TipoComponente, component.NumeroParte, component.Cantidad, component.Stock)

	// Update
	db.Model(&product).Update("Stock", +1)
	fmt.Println(product.ID, product.Model, product.NumeroParte, product.Stock)

	// Delete
	db.Delete(&product, 1)