# PIA Back-end
Equipo:

* MORENO GARCIA JOSE RAUL 1558818
* SEBASTIAN HERNANDEZ SANCHEZ 1746694
* KEVIN CRUZ VEGA SIERRA 1821382

## Introducción
La problemática de nuestro proyecto se encuentra localizada en la empresa Denso, específicamente en el departamento SMD en donde se producen los PCB que después serán utilizados para el ensamble de un cluster automotriz. 

El proceso de la línea SMD en general consiste en 3 estaciones, impresión de soldadura, montaje de componentes y solidificación de soldadura.

La problemática detectada es la falta de material que se ve reflejada en paros de línea, atraso de pedidos y por lo tanto perdida de dinero. En este caso el material faltante es: carretes de componentes, charolas de cpu, jeringas de soldadura en pasta, tubos de conectores, pcb virgen, rollos de limpieza (para máquina) y alcohol.

## Propuesta técnica
La función principal de nuestra API, será abastecer, en tiempo y forma, del material requerido por el operador de SMD para evitar tener paros de línea por falta de material, siendo de esta manera, muy útil nuestra aplicación.
* Entidades: carretes de componentes, charolas de cpu, jeringas de soldadura en pasta, tubos de conectores, pcb virgen, rollos de limpieza (para máquina) y alcohol.
* Atributos: número de parte, tipo de componente (resistencia, capacitor, transistor, etc), cantidad, proveedor, capacidad (ohms, faradios, etc), stock.
* Funcionalidades: registrar números de parte nuevos, modificarlos, actualizarlos y borrarlos; comparar material solicitado con el registrado en la base de datos y verificar si hay en existencia.
* Base de datos: La base de datos a utilizar será SQL server y la razón es que todo el equipo ya la conoce.

## Diagrama UML

![image](https://user-images.githubusercontent.com/89288570/130314037-bcdc883a-ccb6-41b7-b93c-6c11c0fd8a14.png)

## Diagrama Entidad-Relación

![Screen Shot 2021-08-21 at 1 08 26](https://user-images.githubusercontent.com/71417348/130313269-f360005c-7cee-4a63-a65e-4d944622f3e7.png)

