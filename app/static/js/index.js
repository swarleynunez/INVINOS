const API_URL = "http://localhost:8080/api/v1";

// API requests
function createInstance() {
    // Requests
    fetch(API_URL + "/create_instance", { method: "PUT" })
        .then(response => {
            if (response.ok) {
                showToast("success");
                return response.text().then(data => {
                    document.getElementById("api-token").value = data;
                    document.getElementById("tt-body").innerHTML = '<td colspan="9">Sin datos</td>';
                });
            } else {
                showToast("danger");
            }
        });
}

function addProductType() {
    // Form validation
    const id = document.getElementById("add-product-type").value;
    const cert = document.getElementById("apt-cert").files[0];
    if (!id || !cert) { return; }

    // Data
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
            "certificados_vino": [
                {
                    "type_dop": "dop",
                    "name_dop": "D.O. Jumilla",
                    "certificado_dop": {
                        "name": "Certificado de producto",
                        "files": [
                            {
                                "name": cert.name
                            }
                        ],
                        "certificadora": "CCPAE",
                        "fecha_certificacion": "2021-01-01",
                        "fecha_expiracion": "2022-01-01",
                        "info": "Info adicional..."
                    }
                }
            ]
        }
    };

    // Requests
    const formData = new FormData();
    formData.append("file", cert);
    fetch(API_URL + "/certificate/upload", {
        method: "POST",
        headers: { "Authorization": `Bearer ${document.getElementById("api-token").value}` },
        body: formData
    }).then(response => {
        if (response.ok) {
            fetch(API_URL + "/add/product_type", {
                method: "POST",
                headers: {
                    "Authorization": `Bearer ${document.getElementById("api-token").value}`,
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(JSON_ADD_PRODUCT_TYPE)
            }).then(response => {
                if (response.ok) {
                    showToast("success");
                    document.getElementById("add-product-type").value = "";
                    document.getElementById("apt-cert").value = "";
                } else {
                    showToast("danger");
                }
            });
        } else {
            showToast("danger");
        }
    });
}

function addCompany() {
    // Form validation
    const id = document.getElementById("add-company").value;
    if (!id) { return; }

    // Data
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

    // Requests
    fetch(API_URL + "/add/company", {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${document.getElementById("api-token").value}`,
            "Content-Type": "application/json"
        },
        body: JSON.stringify(JSON_ADD_COMPANY)
    }).then(response => {
        if (response.ok) {
            showToast("success");
            document.getElementById("add-company").value = "";
        } else {
            showToast("danger");
        }
    });
}

function addContainer() {
    // Form validation
    const id = document.getElementById("add-container").value;
    if (!id) { return; }

    // Data
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

    // Requests
    fetch(API_URL + "/add/container", {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${document.getElementById("api-token").value}`,
            "Content-Type": "application/json"
        },
        body: JSON.stringify(JSON_ADD_CONTAINER)
    }).then(response => {
        if (response.ok) {
            showToast("success");
            document.getElementById("add-container").value = "";
        } else {
            showToast("danger");
        }

    });
}

function productEntry() {
    // Form validation
    const date = new Date();
    const productType = document.getElementById("pe-product-type").value;
    const quantity = document.getElementById("pe-quantity").value;
    const company = document.getElementById("pe-company").value;
    const container = document.getElementById("pe-container").value;
    const cert = document.getElementById("pe-cert").files[0];
    if (!productType || !quantity || !company || !container || !cert) { return; }

    // Data
    var JSON_PRODUCT_ENTRY = {
        "id": "ID_INTERNO_ENTRADA",
        "fecha": date,
        "producto": {
            "id": 1,
            "nombre": productType,
            "tipo": productType,
            "cantidad": {
                "valor": Number(quantity),
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
                "certificados_vino": [
                    {
                        "type_dop": "dop",
                        "name_dop": "D.O. Jumilla",
                        "certificado_dop": {
                            "name": "Certificado de producto",
                            "files": [
                                {
                                    "name": cert.name
                                }
                            ],
                            "certificadora": "CCPAE",
                            "fecha_certificacion": "2021-01-01",
                            "fecha_expiracion": "2022-01-01",
                            "info": "Info adicional..."
                        }
                    }
                ]
            }
        },
        "vendedor": {
            "id": company,
            "nombre": company,
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

    // Requests
    const formData = new FormData();
    formData.append("file", cert);
    fetch(API_URL + "/certificate/upload", {
        method: "POST",
        headers: { "Authorization": `Bearer ${document.getElementById("api-token").value}` },
        body: formData
    }).then(response => {
        if (response.ok) {
            fetch(API_URL + "/product/entry", {
                method: "POST",
                headers: {
                    "Authorization": `Bearer ${document.getElementById("api-token").value}`,
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(JSON_PRODUCT_ENTRY)
            }).then(response => {
                if (response.ok) {
                    showToast("success");
                    document.getElementById("pe-product-type").value = "";
                    document.getElementById("pe-quantity").value = "";
                    document.getElementById("pe-company").value = "";
                    document.getElementById("pe-container").value = "";
                    document.getElementById("pe-cert").value = "";
                } else {
                    showToast("danger");
                }
            });
        } else {
            showToast("danger");
        }
    });
}

function productProcessing() {
    // Form validation
    const date = new Date();
    const productId = document.getElementById("pp-product-id").value;
    const action = document.getElementById("pp-action").value;
    const productType = document.getElementById("pp-product-type").value;
    const losses = document.getElementById("pp-losses").value;
    const container = document.getElementById("pp-container").value;
    if (!productId || !action || !productType || !losses || !container) { return; }

    // Data
    var JSON_PRODUCT_PROCESSING = {
        "id": "ID_INTERNO_PROCESADO",
        "tipo": action,
        "fecha": date,
        "producto": {
            "id": Number(productId),
            "nombre": productType,
            "tipo": productType,
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
        "merma": {
            "valor": Number(losses),
            "unidades": "Litros"
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

    // Requests
    fetch(API_URL + "/product/processing", {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${document.getElementById("api-token").value}`,
            "Content-Type": "application/json"
        },
        body: JSON.stringify(JSON_PRODUCT_PROCESSING)
    }).then(response => {
        if (response.ok) {
            showToast("success");
            document.getElementById("pp-product-id").value = "";
            document.getElementById("pp-action").value = "";
            document.getElementById("pp-product-type").value = "";
            document.getElementById("pp-losses").value = "";
            document.getElementById("pp-container").value = "";

        } else {
            showToast("danger");
        }
    });
}

function productPartition() {
    // Form validation
    const date = new Date();
    const productId = document.getElementById("ppt-product-id").value;
    const quantity = document.getElementById("ppt-quantity").value;
    const container = document.getElementById("ppt-container").value;
    if (!productId || !quantity || !container) { return; }

    // Data
    var JSON_PRODUCT_PARTITION = {
        "id": "ID_INTERNO_PARTICION",
        "fecha": date,
        "producto": {
            "id": Number(productId),
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
        "cantidad": {
            "valor": Number(quantity),
            "unidades": "Litros"
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

    // Requests
    fetch(API_URL + "/product/partition", {
        method: "POST",
        headers: {
            "Authorization": `Bearer ${document.getElementById("api-token").value}`,
            "Content-Type": "application/json"
        },
        body: JSON.stringify(JSON_PRODUCT_PARTITION)
    }).then(response => {
        if (response.ok) {
            showToast("success");
            document.getElementById("ppt-product-id").value = "";
            document.getElementById("ppt-quantity").value = "";
            document.getElementById("ppt-container").value = "";
        } else {
            showToast("danger");
        }
    });
}

function productOutput() {
    // Form validation
    const date = new Date();
    const productId = document.getElementById("po-product-id").value;
    const lot = document.getElementById("po-lot").value;
    const company = document.getElementById("po-company").value;
    const container = document.getElementById("po-container").value;
    const cert = document.getElementById("po-cert").files[0];
    if (!productId || !lot || !company || !container || !cert) { return; }

    // Data
    var JSON_PRODUCT_OUTPUT = {
        "id": "ID_INTERNO_SALIDA",
        "fecha": date,
        "salida": {
            "producto": {
                "id": Number(productId),
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
                    "certificados_vino": [
                        {
                            "type_dop": "dop",
                            "name_dop": "D.O. Jumilla",
                            "certificado_dop": {
                                "name": "Certificado de producto",
                                "files": [
                                    {
                                        "name": cert.name
                                    }
                                ],
                                "certificadora": "CCPAE",
                                "fecha_certificacion": "2021-01-01",
                                "fecha_expiracion": "2022-01-01",
                                "info": "Info adicional..."
                            }
                        }
                    ]
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
            "id": company,
            "nombre": company,
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

    // Requests
    const formData = new FormData();
    formData.append("file", cert);
    fetch(API_URL + "/certificate/upload", {
        method: "POST",
        headers: { "Authorization": `Bearer ${document.getElementById("api-token").value}` },
        body: formData
    }).then(response => {
        if (response.ok) {
            fetch(API_URL + "/product/output", {
                method: "POST",
                headers: {
                    "Authorization": `Bearer ${document.getElementById("api-token").value}`,
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(JSON_PRODUCT_OUTPUT)
            }).then(response => {
                if (response.ok) {
                    showToast("success");
                    document.getElementById("po-product-id").value = "";
                    document.getElementById("po-lot").value = "";
                    document.getElementById("po-company").value = "";
                    document.getElementById("po-container").value = "";
                    document.getElementById("po-cert").value = "";
                } else {
                    showToast("danger");
                }
            });
        } else {
            showToast("danger");
        }
    });
}

function getTraceability() {
    // Form validation
    const lot = document.getElementById("gt-lot").value;
    if (!lot) { return; }

    // Requests
    fetch(API_URL + "/traceability/" + lot, {
        method: "GET",
        headers: { "Authorization": `Bearer ${document.getElementById("api-token").value}` }
    }).then(response => {
        if (response.ok) {
            showToast("success");
            return response.json().then(data => {
                printTraceability(data);
            });
        } else {
            showToast("danger");
        }
    });
}

// Utils
function showToast(mode) {
    const toast = document.getElementById("toast-" + mode);
    const t = new bootstrap.Toast(toast, { delay: 1000 });
    t.show();
}

function validateNumber(e) {
    if (e.key === "." || e.key === "," || e.key === "-") {
        e.preventDefault();
    }
}

function printTraceability(transitions) {
    const ttb = document.getElementById("tt-body");
    ttb.innerHTML = "";
    transitions.reverse().forEach(t => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${t.CurrentProductID}</td>
            <td>${parseTransitionType(t.Transition.TypeID)}</td>
            <td>${t.Quantity}</td>
            <td>${t.Transition.LostQuantity}</td>
            <td>${t.Transition.ProductTypeID}</td>
            <td>${t.Transition.CompanyID}</td>
            <td>${t.Transition.ContainerID}</td>
            <td>${JSON.parse(t.Transition.Info).fecha}</td>
        `;

        // Search for certificates
        const info = JSON.parse(t.Transition.Info);
        if ("producto" in info) {
            if ("certificados_vino" in info.producto.info) {
                const file = info.producto.info.certificados_vino[0].certificado_dop.files[0];
                row.innerHTML += `
                    <td>
                        <button type="button" class="btn btn-success"
                        onclick="downloadCertificate('${file.cid}', '${file.name}')">Descargar</button>
                    </td>
                `;

            }
        } else if ("salida" in info) {
            const file = info.salida.producto.info.certificados_vino[0].certificado_dop.files[0];
            row.innerHTML += `
                <td>
                    <button type="button" class="btn btn-success"
                    onclick="downloadCertificate('${file.cid}', '${file.name}')">Descargar</button>
                </td>
            `;
        }

        // Add row to the table
        ttb.appendChild(row);
    });
}

function downloadCertificate(cid, filename) {
    // Requests
    window.open(API_URL + "/certificate/download?cid=" + cid + "&filename=" + filename, "_blank");
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

function resetTraceabilityTable() {
    document.getElementById("tt-body").innerHTML = '<td colspan="9">Sin datos</td>';
}
