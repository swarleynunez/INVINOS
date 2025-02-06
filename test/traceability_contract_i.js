const fetch = require("node-fetch");

contract("Traceability", _ => {

    const API_URL = "http://localhost:8080/api/v1";
    let API_KEY;

    // Mocha hook
    beforeEach("Init", async () => {
        const r = await fetch(API_URL + "/create_instance", { method: "PUT" });
        const rbody = await r.text();
        API_KEY = rbody
    });

    it("E5.10-1: Añadir usuario (APIKey)", async () => {

        const code = 200;
        console.log("Código HTTP esperado:", code);

        const r = await fetch(API_URL + "/create_instance", { method: "PUT" });
        const rcode = r.status;
        const rbody = await r.text();

        API_KEY = rbody;
        console.log("Nueva APIKey:", API_KEY);

        assert.equal(rcode, code);
    });

    it("E5.10-2: Autenticación de usuario (EXISTE APIKey)", async () => {

        const id = "GRANEL1";
        const code = 200;
        console.log("APIKey:", API_KEY);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_PRODUCT_TYPE = {
            "id": id,
            "nombre": id,
            "info": {
                "color_vino": "tinto",
                "quimica_vino": {
                    "ga": 12.5,
                    "densidad": 0.99,
                    "otros": "5.5 g/l de acidez total"
                },
                "elaboracion": "joven",
                "composicion_vino": [
                    {
                        "porcentaje": 100,
                        "uva": {
                            "variedad": "Cabernet Sauvignon",
                            "añada": 2020,
                            "color_uva": "tinto",
                            "conduccion": "espaldera",
                            "vendimia": "manual",
                            "certificados_uva": []
                        }
                    }
                ],
                "certificados_vino": []
            }
        };

        const r = await fetch(API_URL + "/add/product_type", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-2: Autenticación de usuario (NO EXISTE APIKey)", async () => {

        const id = "GRANEL1";
        const code = 401;
        API_KEY = "0000000000000000000000000000000000000000000000000000000000000000"
        console.log("APIKey:", API_KEY);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_PRODUCT_TYPE = {
            "id": id,
            "nombre": id,
            "info": {
                "color_vino": "tinto",
                "quimica_vino": {
                    "ga": 12.5,
                    "densidad": 0.99,
                    "otros": "5.5 g/l de acidez total"
                },
                "elaboracion": "joven",
                "composicion_vino": [
                    {
                        "porcentaje": 100,
                        "uva": {
                            "variedad": "Cabernet Sauvignon",
                            "añada": 2020,
                            "color_uva": "tinto",
                            "conduccion": "espaldera",
                            "vendimia": "manual",
                            "certificados_uva": []
                        }
                    }
                ],
                "certificados_vino": []
            }
        };

        const r = await fetch(API_URL + "/add/product_type", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-3: Validar empresa (EXISTE)", async () => {

        const id = "EMPRESA1";
        const code = 409;
        console.log("ID empresa:", id);
        console.log("Código HTTP esperado:", code);

        await addCompany(id);
        var JSON_ADD_COMPANY = {
            "id": id,
            "nombre": id,
            "cif": "B12345678",
            "tipo_std": "Bodega",
            "info": {
                "direccion": "C/ Ejemplo, 123",
                "telefono": "987654321",
                "email": "contacto@empresa.com",
                "web": "www.empresa.com"
            }
        };

        const r = await fetch(API_URL + "/add/company", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_COMPANY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-3: Validar empresa (NO EXISTE)", async () => {

        const id = "ABCDEFG";
        const code = 409;
        const date = new Date();
        const productType = "GRANEL1";
        const quantity = 1000;
        const container = "CONTENEDOR1";
        console.log("ID empresa:", id);
        console.log("Código HTTP esperado:", code);

        await addProductType(productType);
        await addContainer(container);
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": id,
                "nombre": id,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };

        const r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-4: Añadir empresa (NO CREADA)", async () => {

        const id = "EMPRESA1";
        const code = 200;
        console.log("ID empresa:", id);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_COMPANY = {
            "id": id,
            "nombre": id,
            "cif": "B12345678",
            "tipo_std": "Bodega",
            "info": {
                "direccion": "C/ Ejemplo, 123",
                "telefono": "987654321",
                "email": "contacto@empresa.com",
                "web": "www.empresa.com"
            }
        };

        const r = await fetch(API_URL + "/add/company", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_COMPANY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-4: Añadir empresa (SIN DATOS)", async () => {

        const id = "";
        const code = 400;
        console.log("ID empresa:", id);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_COMPANY = {
            "id": id,
            "nombre": id,
            "cif": "B12345678",
            "tipo_std": "Bodega",
            "info": {
                "direccion": "C/ Ejemplo, 123",
                "telefono": "987654321",
                "email": "contacto@empresa.com",
                "web": "www.empresa.com"
            }
        };

        const r = await fetch(API_URL + "/add/company", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_COMPANY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-4: Añadir empresa (YA CREADA)", async () => {

        const id = "EMPRESA1";
        const code = 409;
        console.log("ID empresa:", id);
        console.log("Código HTTP esperado:", code);

        await addCompany(id);
        var JSON_ADD_COMPANY = {
            "id": id,
            "nombre": id,
            "cif": "B12345678",
            "tipo_std": "Bodega",
            "info": {
                "direccion": "C/ Ejemplo, 123",
                "telefono": "987654321",
                "email": "contacto@empresa.com",
                "web": "www.empresa.com"
            }
        };

        const r = await fetch(API_URL + "/add/company", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_COMPANY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-5: Validar contenedor (EXISTE)", async () => {

        const id = "CONTENEDOR1";
        const code = 409;
        console.log("ID contenedor:", id);
        console.log("Código HTTP esperado:", code);

        await addContainer(id);
        var JSON_ADD_CONTAINER = {
            "id": id,
            "info": {
                "tipo": "Descargadero",
                "dimensiones": {
                    "alto": 0,
                    "ancho": 0,
                    "largo": 0,
                    "capacidad": 0
                },
                "material": "Acero inoxidable"
            }
        };

        const r = await fetch(API_URL + "/add/container", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_CONTAINER)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-5: Validar contenedor (NO EXISTE)", async () => {

        const id = "ABCDEFG";
        const code = 409;
        const date = new Date();
        const productType = "GRANEL1";
        const quantity = 1000;
        const empresa = "EMPRESA1";
        console.log("ID contenedor:", id);
        console.log("Código HTTP esperado:", code);

        await addProductType(productType);
        await addCompany(empresa);
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": id,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };

        const r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-6: Añadir contenedor (NO CREADO)", async () => {

        const id = "CONTENEDOR1";
        const code = 200;
        console.log("ID contenedor:", id);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_CONTAINER = {
            "id": id,
            "info": {
                "tipo": "Descargadero",
                "dimensiones": {
                    "alto": 0,
                    "ancho": 0,
                    "largo": 0,
                    "capacidad": 0
                },
                "material": "Acero inoxidable"
            }
        };

        const r = await fetch(API_URL + "/add/container", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_CONTAINER)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-6: Añadir contenedor (SIN DATOS)", async () => {

        const id = "";
        const code = 400;
        console.log("ID contenedor:", id);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_CONTAINER = {
            "id": id,
            "info": {
                "tipo": "Descargadero",
                "dimensiones": {
                    "alto": 0,
                    "ancho": 0,
                    "largo": 0,
                    "capacidad": 0
                },
                "material": "Acero inoxidable"
            }
        };

        const r = await fetch(API_URL + "/add/container", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_CONTAINER)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-6: Añadir contenedor (YA CREADO)", async () => {

        const id = "CONTENEDOR1";
        const code = 409;
        console.log("ID contenedor:", id);
        console.log("Código HTTP esperado:", code);

        await addContainer(id);
        var JSON_ADD_CONTAINER = {
            "id": id,
            "info": {
                "tipo": "Descargadero",
                "dimensiones": {
                    "alto": 0,
                    "ancho": 0,
                    "largo": 0,
                    "capacidad": 0
                },
                "material": "Acero inoxidable"
            }
        };

        const r = await fetch(API_URL + "/add/container", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_CONTAINER)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-7: Validar tipo de producto (EXISTE)", async () => {

        const id = "GRANEL1";
        const code = 409;
        console.log("ID tipo de producto:", id);
        console.log("Código HTTP esperado:", code);

        await addProductType(id);
        var JSON_ADD_PRODUCT_TYPE = {
            "id": id,
            "nombre": id,
            "info": {
                "color_vino": "tinto",
                "quimica_vino": {
                    "ga": 12.5,
                    "densidad": 0.99,
                    "otros": "5.5 g/l de acidez total"
                },
                "elaboracion": "joven",
                "composicion_vino": [
                    {
                        "porcentaje": 100,
                        "uva": {
                            "variedad": "Cabernet Sauvignon",
                            "añada": 2020,
                            "color_uva": "tinto",
                            "conduccion": "espaldera",
                            "vendimia": "manual",
                            "certificados_uva": []
                        }
                    }
                ],
                "certificados_vino": []
            }
        };

        const r = await fetch(API_URL + "/add/product_type", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-7: Validar tipo de producto (NO EXISTE)", async () => {

        const id = "ABCDEFG";
        const code = 409;
        const date = new Date();
        const quantity = 1000;
        const empresa = "EMPRESA1";
        const container = "CONTENEDOR1";
        console.log("ID tipo de producto:", id);
        console.log("Código HTTP esperado:", code);

        await addCompany(empresa);
        await addContainer(container);
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": id,
                "tipo": id,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };

        const r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-8: Añadir tipo de producto (NO CREADO)", async () => {

        const id = "GRANEL1";
        const code = 200;
        console.log("ID tipo de producto:", id);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_PRODUCT_TYPE = {
            "id": id,
            "nombre": id,
            "info": {
                "color_vino": "tinto",
                "quimica_vino": {
                    "ga": 12.5,
                    "densidad": 0.99,
                    "otros": "5.5 g/l de acidez total"
                },
                "elaboracion": "joven",
                "composicion_vino": [
                    {
                        "porcentaje": 100,
                        "uva": {
                            "variedad": "Cabernet Sauvignon",
                            "añada": 2020,
                            "color_uva": "tinto",
                            "conduccion": "espaldera",
                            "vendimia": "manual",
                            "certificados_uva": []
                        }
                    }
                ],
                "certificados_vino": []
            }
        };

        const r = await fetch(API_URL + "/add/product_type", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-8: Añadir tipo de producto (SIN DATOS)", async () => {

        const id = "";
        const code = 400;
        console.log("ID tipo de producto:", id);
        console.log("Código HTTP esperado:", code);

        var JSON_ADD_PRODUCT_TYPE = {
            "id": id,
            "nombre": id,
            "info": {
                "color_vino": "tinto",
                "quimica_vino": {
                    "ga": 12.5,
                    "densidad": 0.99,
                    "otros": "5.5 g/l de acidez total"
                },
                "elaboracion": "joven",
                "composicion_vino": [
                    {
                        "porcentaje": 100,
                        "uva": {
                            "variedad": "Cabernet Sauvignon",
                            "añada": 2020,
                            "color_uva": "tinto",
                            "conduccion": "espaldera",
                            "vendimia": "manual",
                            "certificados_uva": []
                        }
                    }
                ],
                "certificados_vino": []
            }
        };

        const r = await fetch(API_URL + "/add/product_type", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-8: Añadir tipo de producto (YA CREADO)", async () => {

        const id = "GRANEL1";
        const code = 409;
        console.log("ID tipo de producto:", id);
        console.log("Código HTTP esperado:", code);

        await addProductType(id);
        var JSON_ADD_PRODUCT_TYPE = {
            "id": id,
            "nombre": id,
            "info": {
                "color_vino": "tinto",
                "quimica_vino": {
                    "ga": 12.5,
                    "densidad": 0.99,
                    "otros": "5.5 g/l de acidez total"
                },
                "elaboracion": "joven",
                "composicion_vino": [
                    {
                        "porcentaje": 100,
                        "uva": {
                            "variedad": "Cabernet Sauvignon",
                            "añada": 2020,
                            "color_uva": "tinto",
                            "conduccion": "espaldera",
                            "vendimia": "manual",
                            "certificados_uva": []
                        }
                    }
                ],
                "certificados_vino": []
            }
        };

        const r = await fetch(API_URL + "/add/product_type", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-9: Añadir transición (DATOS CORRECTOS)", async () => {

        const date = new Date();
        const quantity = 1000;
        const productType = "GRANEL1";
        const empresa = "EMPRESA1";
        const container = "CONTENEDOR1";
        const code = 200;
        console.log("Cantidad de producto:", quantity);
        console.log("ID tipo de producto:", productType);
        console.log("ID empresa:", empresa);
        console.log("ID contenedor:", container);
        console.log("Código HTTP esperado:", code);

        await addProductType(productType);
        await addCompany(empresa);
        await addContainer(container);
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };

        const r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        const rcode = r.status;

        assert.equal(rcode, code);
    });

    it("E5.10-9: Añadir transición (DATOS INCORRECTOS)", async () => {

        const date = new Date();
        var quantity = 0;
        var productType = "ABCDEFG";
        var empresa = "ABCDEFG";
        var container = "ABCDEFG";
        const code1 = 400;
        const code2 = 409;
        console.log("Cantidad de producto:", quantity);
        console.log("ID tipo de producto:", productType);
        console.log("ID empresa:", empresa);
        console.log("ID contenedor:", container);
        console.log("Códigos HTTP esperados:");
        console.log("   ", code1);
        console.log("   ", code2);
        console.log("   ", code2);

        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };
        var r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        var rcode = r.status;
        assert.equal(rcode, code1);

        quantity = 1000;
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };
        var r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        var rcode = r.status;
        assert.equal(rcode, code2);

        productType = "GRANEL1";
        await addProductType(productType);
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };
        var r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        var rcode = r.status;
        assert.equal(rcode, code2);

        empresa = "EMPRESA1";
        await addCompany(empresa);
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };
        var r = await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        var rcode = r.status;
        assert.equal(rcode, code2);
    });

    it("E5.10-10: Obtener trazabilidad (EXISTE LOTE)", async () => {

        const id = 1;
        const date = new Date();
        const quantity = 1000;
        const productType = "GRANEL1";
        const empresa = "EMPRESA1";
        const container = "CONTENEDOR1";
        const lot = "LT-1234";
        const code = 200;
        console.log("ID producto:", id);
        console.log("Número de lote:", lot);
        console.log("Código HTTP esperado:", code);
        console.log("Transiciones esperadas:");

        await addProductType(productType);
        await addCompany(empresa);
        await addContainer(container);
        var JSON_PRODUCT_ENTRY = {
            "id": "ID_INTERNO_ENTRADA",
            "fecha": date,
            "producto": {
                "id": 1,
                "nombre": productType,
                "tipo": productType,
                "cantidad": {
                    "valor": quantity,
                    "unidades": "Litros"
                },
                "info": {
                    "color_vino": "tinto",
                    "quimica_vino": {
                        "ga": 12.5,
                        "densidad": 0.99,
                        "otros": "5.5 g/l de acidez total"
                    },
                    "elaboracion": "joven",
                    "composicion_vino": [
                        {
                            "porcentaje": 100,
                            "uva": {
                                "variedad": "Cabernet Sauvignon",
                                "añada": 2020,
                                "color_uva": "tinto",
                                "conduccion": "espaldera",
                                "vendimia": "manual",
                                "certificados_uva": []
                            }
                        }
                    ],
                    "certificados_vino": []
                }
            },
            "vendedor": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };
        await fetch(API_URL + "/product/entry", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_ENTRY)
        });
        var JSON_PRODUCT_OUTPUT = {
            "id": "ID_INTERNO_SALIDA",
            "fecha": date,
            "salida": {
                "producto": {
                    "id": id,
                    "info": {
                        "color_vino": "tinto",
                        "quimica_vino": {
                            "ga": 12.5,
                            "densidad": 0.99,
                            "otros": "5.5 g/l de acidez total"
                        },
                        "elaboracion": "joven",
                        "composicion_vino": [
                            {
                                "porcentaje": 100,
                                "uva": {
                                    "variedad": "Cabernet Sauvignon",
                                    "añada": 2020,
                                    "color_uva": "tinto",
                                    "conduccion": "espaldera",
                                    "vendimia": "manual",
                                    "certificados_uva": []
                                }
                            }
                        ],
                        "certificados_vino": []
                    }
                },
                "lote": {
                    "num_lote": lot,
                    "caducidad": "2021-12-31",
                    "fecha_envasado": "2021-01-01",
                    "url": API_URL + "/traceability/" + lot
                }
            },
            "comprador": {
                "id": empresa,
                "nombre": empresa,
                "cif": "B12345678",
                "tipo_std": "Bodega",
                "info": {
                    "direccion": "C/ Ejemplo, 123",
                    "telefono": "987654321",
                    "email": "contacto@empresa.com",
                    "web": "www.empresa.com"
                }
            },
            "a_contenedor": {
                "id": container,
                "info": {
                    "tipo": "Descargadero",
                    "dimensiones": {
                        "alto": 0,
                        "ancho": 0,
                        "largo": 0,
                        "capacidad": 0
                    },
                    "material": "Acero inoxidable"
                }
            },
            "info": "Info adicional..."
        };
        await fetch(API_URL + "/product/output", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_PRODUCT_OUTPUT)
        });

        var r = await fetch(API_URL + "/traceability/" + lot, {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${API_KEY}`
            }
        });
        const rcode = r.status;
        const rbody = await r.json();
        assert.equal(rcode, code);

        printTraceability(rbody);

        const t1 = rbody[0];
        assert.equal(t1.CurrentProductID, id);
        assert.equal(t1.Transition.TypeID, 0);
        assert.equal(t1.Quantity, quantity);
        assert.equal(t1.Transition.LostQuantity, 0);
        assert.equal(t1.Transition.ProductTypeID, productType);
        assert.equal(t1.Transition.CompanyID, empresa);
        assert.equal(t1.Transition.ContainerID, container);
        const t2 = rbody[1];
        assert.equal(t2.CurrentProductID, id);
        assert.equal(t2.Transition.TypeID, 3);
        assert.equal(t2.Quantity, quantity);
        assert.equal(t2.Transition.LostQuantity, 0);
        assert.equal(t2.Transition.ProductTypeID, productType);
        assert.equal(t2.Transition.CompanyID, empresa);
        assert.equal(t2.Transition.ContainerID, container);

    });

    it("E5.10-10: Obtener trazabilidad (NO EXISTE LOTE)", async () => {

        const lot = "LT-999999";
        const code = 404;
        console.log("Número de lote:", lot);
        console.log("Código HTTP esperado:", code);

        var r = await fetch(API_URL + "/traceability/" + lot, {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${API_KEY}`
            }
        });
        const rcode = r.status;
        assert.equal(rcode, code);
    });

    async function addProductType(id) {
        var JSON_ADD_PRODUCT_TYPE = {
            "id": id,
            "nombre": id,
            "info": {
                "color_vino": "tinto",
                "quimica_vino": {
                    "ga": 12.5,
                    "densidad": 0.99,
                    "otros": "5.5 g/l de acidez total"
                },
                "elaboracion": "joven",
                "composicion_vino": [
                    {
                        "porcentaje": 100,
                        "uva": {
                            "variedad": "Cabernet Sauvignon",
                            "añada": 2020,
                            "color_uva": "tinto",
                            "conduccion": "espaldera",
                            "vendimia": "manual",
                            "certificados_uva": []
                        }
                    }
                ],
                "certificados_vino": []
            }
        };

        await fetch(API_URL + "/add/product_type", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
        });
    }

    async function addCompany(id) {
        var JSON_ADD_COMPANY = {
            "id": id,
            "nombre": id,
            "cif": "B12345678",
            "tipo_std": "Bodega",
            "info": {
                "direccion": "C/ Ejemplo, 123",
                "telefono": "987654321",
                "email": "contacto@empresa.com",
                "web": "www.empresa.com"
            }
        };

        await fetch(API_URL + "/add/company", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_COMPANY)
        });
    }

    async function addContainer(id) {
        var JSON_ADD_CONTAINER = {
            "id": id,
            "info": {
                "tipo": "Descargadero",
                "dimensiones": {
                    "alto": 0,
                    "ancho": 0,
                    "largo": 0,
                    "capacidad": 0
                },
                "material": "Acero inoxidable"
            }
        };

        await fetch(API_URL + "/add/container", {
            method: "POST",
            headers: {
                "Authorization": `Bearer ${API_KEY}`,
                "Content-Type": "application/json"
            },
            body: JSON.stringify(JSON_ADD_CONTAINER)
        });
    }

    function printTraceability(transitions) {
        transitions.reverse().forEach(t => {
            console.log(
                "    PRODUCT:" + t.CurrentProductID, "|",
                parseTransitionType(t.Transition.TypeID), "|",
                t.Quantity, "|",
                t.Transition.LostQuantity, "|",
                t.Transition.ProductTypeID, "|",
                t.Transition.CompanyID, "|",
                t.Transition.ContainerID, "|",
                parseTimestamp(t.Transition.CreatedAt), "|",
                "{JSON}");
        });
    }

    function parseTransitionType(type) {
        switch (type) {
            case 0:
                return "ENTRADA";
            case 1:
                return "PROCESADO";
            case 2:
                return "PARTICIÓN";
            case 3:
                return "SALIDA";
        }
    }

    function parseTimestamp(timestamp) {
        return new Date(timestamp * 1000);
    }
});


