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

