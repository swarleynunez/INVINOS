# Proyecto INVINOS
            PRODUCT:1
            /       \
           /         \
          /           \
    PRODUCT:2       PRODUCT:3
                     /      \
                    /        \
                   /          \
                PRODUCT:4    PRODUCT:5
                               /      \
                              /        \
                             /          \
                        PRODUCT:6     PRODUCT:7


Lot number of PRODUCT:6 --> "LN-6"

![alt text](docs/images/ln-6.png)

Lot number of PRODUCT:7 --> "LN-7"

![alt text](docs/images/ln-7.png)

### Cuestiones a debatir:

- ID de producto enviado a nuestra API en las peticiones.
- Gestión de EOAs en la API o en el cliente?
- Gestión de mermas en la función productProcessing().
- Un mismo número de lote se utilizará para varios IDs de producto.
