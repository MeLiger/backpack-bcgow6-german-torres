Ejercicio
Una mueblería necesita la implementación de una base de datos para controlar las ventas que realiza por día, el stock de sus artículos (productos) y la lista de sus clientes que realizan las compras.

Se necesita plantear:
1 - ¿Cuáles serían las entidades de este sistema?
2 - ¿Qué atributos se determinarán para cada entidad? (Considerar los que se crean necesarios)
3 - ¿Cómo se conformarán las relaciones entre entidades? ¿Cuáles serían las cardinalidades?

Realizar un DER para modelar el escenario planteado.



1 - Cliente, Pedido, Producto.

2 - Cliente: Dni, Nombre, Apellido, Dirección, Fecha de Nacimiento, Teléfono.
    Venta: Numero, fecha
    Producto: Código, Precio, Nombre, Cantidad.

3 - Cliente -realiza - Pedido - tiene -Producto



Ejercicio
Realizar un diagrama de entidad - relación para el sistema de una concesionaria, que desea gestionar los servicios de los coches de sus clientes. 

Para el módulo del sistema, se necesita almacenar información de los clientes, los coches que estos poseen y los service/revisiones de cada uno de estos.
Utilizar el formato adecuado para representar las Primary y Foreign Key en el diagrama, además de los tipos de datos de cada atributo.
